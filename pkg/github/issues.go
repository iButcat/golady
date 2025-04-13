package github

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/iButcat/golady/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v60/github"
)

type Issues interface {
	GetFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate)
	GetAllFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate)
	CreateFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate)
	GenerateForm(s *discordgo.Session, m *discordgo.MessageCreate)
	AddAssignee(s *discordgo.Session, m *discordgo.MessageCreate)
	GetActiveOrClosed(s *discordgo.Session, m *discordgo.MessageCreate)
}

type issues struct {
	ghClient *github.Client
}

func NewIssues(ghClient *github.Client) Issues {
	return &issues{
		ghClient: ghClient,
	}
}

func (i *issues) GetFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 3
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "Usage: !issue owner repo issue_number")
		return
	}

	id, _ := strconv.Atoi(args[2])

	res, _, err := i.ghClient.Issues.Get(context.Background(), args[0], args[1], id)
	if err != nil {
		utils.DiscordError(s, m, "Error retrieving issue: "+err.Error())
		return
	}

	_ = s.MessageReactionAdd(m.ChannelID, m.ID, "👍")

	var description = make([]string, 0)
	description = append(description,
		"**Body:** "+res.GetBody(),
		"**State:** "+res.GetState(),
		"**Created at:** "+utils.FormatGitHubTime(res.CreatedAt),
		"**Created by:** "+res.GetUser().GetLogin(),
	)

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		URL:         res.GetHTMLURL(),
		Title:       res.GetTitle(),
		Description: strings.Join(description, "\n"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: res.GetUser().GetAvatarURL(),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "GitHub",
		},
		Color: 0x00ff00,
	})
}

func (i *issues) GetAllFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 2
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "Usage: !issues owner repo")
		return
	}

	issues, _, err := i.ghClient.Issues.ListByRepo(context.Background(), args[0], args[1], nil)
	if err != nil {
		utils.DiscordError(s, m, "Error retrieving issues: "+err.Error())
		return
	}

	if len(issues) == 0 {
		utils.DiscordResponse(s, m, "No issues found for "+args[0]+"/"+args[1])
		return
	}

	var issueList []string
	for _, issue := range issues {
		issueList = append(issueList, "- #"+strconv.Itoa(issue.GetNumber())+" "+issue.GetTitle())
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Issues for " + args[0] + "/" + args[1],
		Description: strings.Join(issueList, "\n"),
		URL:         "https://github.com/" + args[0] + "/" + args[1] + "/issues",
		Color:       0x00ff00,
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

type formStruct struct {
	title    string
	body     string
	label    []string
	assignee string
}

func (i *issues) GenerateForm(s *discordgo.Session, m *discordgo.MessageCreate) {
	utils.DiscordResponse(s, m, "title: '' \n body: '' \n assignee: '' \n label: ''")
}

func (i *issues) CreateFromRepoName(s *discordgo.Session, m *discordgo.MessageCreate) {
	args, nb := utils.ParseArgs(m.Content)
	if nb != 2 {
		utils.DiscordError(s, m, "Usage: !issue-create owner repo 'title' 'body'")
		return
	}

	reg := regexp.MustCompile("'(.*?)'")
	res := reg.FindAllString(m.Content, -1)

	expected := 2
	if len(res) != expected {
		utils.DiscordError(s, m, "Please provide title and body in single quotes. Example: !issue-create owner repo 'My title' 'Description here'")
		return
	}

	form := formStruct{
		title: utils.RemoveSingleQuote(res[0]),
		body:  utils.RemoveSingleQuote(res[1]),
	}

	issue, _, err := i.ghClient.Issues.Create(context.Background(), args[0], args[1], &github.IssueRequest{
		Title: &form.title,
		Body:  &form.body,
	})
	if err != nil {
		utils.DiscordError(s, m, "Error creating issue: "+err.Error())
		return
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		URL:         issue.GetHTMLURL(),
		Title:       "Issue Created",
		Description: "**Title:** " + *issue.Title + "\n\n" + *issue.Body,
		Color:       0x00ff00,
	})
}

func (i *issues) AddAssignee(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Implementation left as a future task
	utils.DiscordResponse(s, m, "This command is not yet implemented")
}

func (i *issues) GetActiveOrClosed(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 3
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "Usage: !issues-state owner repo state")
		return
	}

	issues, _, err := i.ghClient.Issues.ListByRepo(context.Background(), args[0], args[1], &github.IssueListByRepoOptions{
		State: args[2],
	})
	if err != nil {
		utils.DiscordError(s, m, "Error retrieving issues: "+err.Error())
		return
	}

	if len(issues) == 0 {
		utils.DiscordResponse(s, m, "No "+args[2]+" issues found for "+args[0]+"/"+args[1])
		return
	}

	for _, issue := range issues {
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:         issue.GetHTMLURL(),
			Title:       issue.GetTitle(),
			Description: issue.GetBody(),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: issue.GetUser().GetAvatarURL(),
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "GitHub",
			},
			Color: 0x00ff00,
		})
	}
}
