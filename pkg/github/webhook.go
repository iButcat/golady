package github
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/iButcat/golady/pkg/db"
	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v60/github"
)
type WebhookHandler struct {
	discord *discordgo.Session
	storage *db.Storage
}
func NewWebhookHandler(discord *discordgo.Session, storage *db.Storage) *WebhookHandler {
	return &WebhookHandler{
		discord: discord,
		storage: storage,
	}
}
func (wh *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "" {
		http.Error(w, "Missing X-GitHub-Event header", http.StatusBadRequest)
		return
	}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	switch eventType {
	case "issues":
		wh.handleIssueEvent(payload)
	case "pull_request":
		wh.handlePullRequestEvent(payload)
	default:
		log.Printf("Received unhandled event type: %s", eventType)
	}
	w.WriteHeader(http.StatusOK)
}
func (wh *WebhookHandler) handleIssueEvent(payload []byte) {
	var event github.IssuesEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Error unmarshaling issue event: %v", err)
		return
	}
	repoName := event.GetRepo().GetName()
	repoOwner := event.GetRepo().GetOwner().GetLogin()
	fullName := repoOwner + "/" + repoName
	webhooks, err := wh.storage.Webhook.GetByRepository(repoName)
	if err != nil {
		log.Printf("Error getting webhooks for repository %s: %v", repoName, err)
		return
	}
	var title, description string
	action := event.GetAction()
	issue := event.GetIssue()
	switch action {
	case "opened":
		title = "🔔 New Issue Opened"
		description = "**Issue #" + strconv.Itoa(issue.GetNumber()) + ":** " + issue.GetTitle()
	case "closed":
		title = "✅ Issue Closed"
		description = "**Issue #" + strconv.Itoa(issue.GetNumber()) + ":** " + issue.GetTitle()
	default:
		return
	}
	for _, webhook := range webhooks {
		if strings.EqualFold(webhook.Owner, repoOwner) {
			embed := &discordgo.MessageEmbed{
				Title:       title,
				URL:         issue.GetHTMLURL(),
				Description: description + "\n\n" + issue.GetBody(),
				Color:       0x4078c0, // GitHub blue
				Author: &discordgo.MessageEmbedAuthor{
					Name:    issue.GetUser().GetLogin(),
					IconURL: issue.GetUser().GetAvatarURL(),
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: fullName,
				},
			}
			_, err := wh.discord.ChannelMessageSendEmbed(webhook.ChannelID, embed)
			if err != nil {
				log.Printf("Error sending message to channel %s: %v", webhook.ChannelID, err)
			}
		}
	}
}
func (wh *WebhookHandler) handlePullRequestEvent(payload []byte) {
	var event github.PullRequestEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Error unmarshaling pull request event: %v", err)
		return
	}
	repoName := event.GetRepo().GetName()
	repoOwner := event.GetRepo().GetOwner().GetLogin()
	fullName := repoOwner + "/" + repoName
	webhooks, err := wh.storage.Webhook.GetByRepository(repoName)
	if err != nil {
		log.Printf("Error getting webhooks for repository %s: %v", repoName, err)
		return
	}
	var title, description string
	action := event.GetAction()
	pr := event.GetPullRequest()
	switch action {
	case "opened":
		title = "🔄 New Pull Request"
		description = "**PR #" + strconv.Itoa(pr.GetNumber()) + ":** " + pr.GetTitle()
	case "closed":
		if pr.GetMerged() {
			title = "🎉 Pull Request Merged"
		} else {
			title = "❌ Pull Request Closed"
		}
		description = "**PR #" + strconv.Itoa(pr.GetNumber()) + ":** " + pr.GetTitle()
	default:
		return
	}
	for _, webhook := range webhooks {
		if strings.EqualFold(webhook.Owner, repoOwner) {
			embed := &discordgo.MessageEmbed{
				Title:       title,
				URL:         pr.GetHTMLURL(),
				Description: description + "\n\n" + pr.GetBody(),
				Color:       0x4078c0, // GitHub blue
				Author: &discordgo.MessageEmbedAuthor{
					Name:    pr.GetUser().GetLogin(),
					IconURL: pr.GetUser().GetAvatarURL(),
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: fullName,
				},
			}
			_, err := wh.discord.ChannelMessageSendEmbed(webhook.ChannelID, embed)
			if err != nil {
				log.Printf("Error sending message to channel %s: %v", webhook.ChannelID, err)
			}
		}
	}
}
