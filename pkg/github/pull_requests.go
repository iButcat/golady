package github

import (
	"context"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v60/github"
	"github.com/iButcat/golady/pkg/utils"
)

type (
	PullRequests interface {
		Create(s *discordgo.Session, m *discordgo.MessageCreate)
		GetAllOpen(s *discordgo.Session, m *discordgo.MessageCreate)
		GenerateForm(s *discordgo.Session, m *discordgo.MessageCreate)
	}
	pullRequests struct {
		ghClient *github.Client
	}
)

func NewPullRequests(ghClient *github.Client) PullRequests {
	return &pullRequests{
		ghClient: ghClient,
	}
}

type form struct {
	title string
	body  string
	head  string
	base  string
}

func (p *pullRequests) GenerateForm(s *discordgo.Session, m *discordgo.MessageCreate) {
	utils.DiscordResponse(s, m, "title: '' \n body: '' \n head: '' \n base: ''")
}
func (p *pullRequests) Create(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 2
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "wrong number of arguments")
		return
	}
	reg := regexp.MustCompile("'(.*?)'")
	res := reg.FindAllString(m.Content, -1)
	var form = form{
		title: utils.RemoveSingleQuote(res[0]),
		body:  utils.RemoveSingleQuote(res[1]),
		head:  utils.RemoveSingleQuote(res[2]),
		base:  utils.RemoveSingleQuote(res[3]),
	}
	pullRequest, _, err := p.ghClient.PullRequests.Create(context.Background(), args[0], args[1], &github.NewPullRequest{
		Title: &form.title,
		Body:  &form.body,
		Head:  &form.head,
		Base:  &form.base,
	})
	if err != nil {
		utils.DiscordError(s, m, "Error: while creating pull request")
		return
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		URL:         *pullRequest.HTMLURL,
		Title:       *pullRequest.Title,
		Description: *pullRequest.Body,
	})
}
func (p *pullRequests) GetAllOpen(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 3
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "wrong number of arguments")
		return
	}
	prs, _, err := p.ghClient.PullRequests.List(context.Background(), args[0], args[1], &github.PullRequestListOptions{
		State: args[2],
	})
	if err != nil {
		utils.DiscordError(s, m, "Error: while getting pull requests")
		return
	}
	for _, pr := range prs {
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: *pr.Title,
		})
	}
}
func (p *pullRequests) Get(s *discordgo.Session, m *discordgo.MessageCreate) {
}
