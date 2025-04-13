package github
import (
	"strings"
	"github.com/iButcat/golady/pkg/db"
	"github.com/iButcat/golady/pkg/utils"
	"github.com/bwmarrin/discordgo"
)
type SubscriptionHandler struct {
	storage *db.Storage
}
func NewSubscriptionHandler(storage *db.Storage) *SubscriptionHandler {
	return &SubscriptionHandler{
		storage: storage,
	}
}
func (s *SubscriptionHandler) Subscribe(session *discordgo.Session, m *discordgo.MessageCreate) {
	args, count := utils.ParseArgs(m.Content)
	if count != 1 {
		utils.DiscordError(session, m, "Usage: !subscribe owner/repo")
		return
	}
	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		utils.DiscordError(session, m, "Invalid repository format. Use: owner/repo")
		return
	}
	owner := parts[0]
	repo := parts[1]
	webhookData := &db.WebhookData{
		ServerID:   m.GuildID,
		ChannelID:  m.ChannelID,
		Repository: repo,
		Owner:      owner,
		Events:     "issues,pull_request", // Default events
	}
	if err := s.storage.Webhook.Subscribe(webhookData); err != nil {
		utils.DiscordError(session, m, "Failed to subscribe: "+err.Error())
		return
	}
	utils.DiscordResponse(session, m, "Successfully subscribed to "+owner+"/"+repo+
		"\nYou will receive notifications for new issues and pull requests in this channel.")
}
func (s *SubscriptionHandler) Unsubscribe(session *discordgo.Session, m *discordgo.MessageCreate) {
	args, count := utils.ParseArgs(m.Content)
	if count != 1 {
		utils.DiscordError(session, m, "Usage: !unsubscribe owner/repo")
		return
	}
	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		utils.DiscordError(session, m, "Invalid repository format. Use: owner/repo")
		return
	}
	repo := parts[1]
	if err := s.storage.Webhook.Unsubscribe(m.GuildID, repo); err != nil {
		utils.DiscordError(session, m, "Failed to unsubscribe: "+err.Error())
		return
	}
	utils.DiscordResponse(session, m, "Successfully unsubscribed from "+args[0])
}
func (s *SubscriptionHandler) ListSubscriptions(session *discordgo.Session, m *discordgo.MessageCreate) {
	webhooks, err := s.storage.Webhook.GetByServerID(m.GuildID)
	if err != nil {
		utils.DiscordError(session, m, "Failed to get subscriptions: "+err.Error())
		return
	}
	if len(webhooks) == 0 {
		utils.DiscordResponse(session, m, "No active subscriptions found for this server.")
		return
	}
	var subscriptions []string
	for _, webhook := range webhooks {
		subscriptions = append(subscriptions, webhook.Owner+"/"+webhook.Repository)
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Repository Subscriptions",
		Description: strings.Join(subscriptions, "\n"),
		Color:       0x4078c0,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Use !unsubscribe owner/repo to remove a subscription",
		},
	}
	_, err = session.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		utils.DiscordError(session, m, "Failed to send message: "+err.Error())
	}
}
