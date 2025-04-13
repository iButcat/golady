package github

import (
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v60/github"
	"github.com/iButcat/golady/pkg/utils"
)

type Repositories interface {
	Get(s *discordgo.Session, m *discordgo.MessageCreate)
	GetAll(s *discordgo.Session, m *discordgo.MessageCreate)
}
type repositories struct {
	ghClient *github.Client
}

func NewRepositories(ghClient *github.Client) Repositories {
	return &repositories{
		ghClient: ghClient,
	}
}
func (r *repositories) Get(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 2
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "Usage: !repo owner name")
		return
	}
	repo, _, err := r.ghClient.Repositories.Get(context.Background(), args[0], args[1])
	if err != nil {
		utils.DiscordError(s, m, "Error retrieving repository: "+err.Error())
		return
	}
	var description []string
	description = append(description,
		"**Description:** "+repo.GetDescription(),
		"**Language:** "+repo.GetLanguage(),
		"**Stars:** "+utils.IntToString(repo.GetStargazersCount()),
		"**Forks:** "+utils.IntToString(repo.GetForksCount()),
		"**Open Issues:** "+utils.IntToString(repo.GetOpenIssuesCount()),
		"**Created At:** "+utils.FormatDate(repo.GetCreatedAt().Time),
	)
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		URL:         repo.GetHTMLURL(),
		Title:       repo.GetFullName(),
		Description: strings.Join(description, "\n"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: repo.GetOwner().GetAvatarURL(),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "GitHub",
		},
		Color: 0x00ff00,
	})
}
func (r *repositories) GetAll(s *discordgo.Session, m *discordgo.MessageCreate) {
	requiredArgs := 1
	args, nb := utils.ParseArgs(m.Content)
	if nb != requiredArgs {
		utils.DiscordError(s, m, "Usage: !repos owner")
		return
	}
	repos, _, err := r.ghClient.Repositories.List(context.Background(), args[0], nil)
	if err != nil {
		utils.DiscordError(s, m, "Error retrieving repositories: "+err.Error())
		return
	}
	if len(repos) == 0 {
		utils.DiscordResponse(s, m, "No repositories found for "+args[0])
		return
	}
	var repoList []string
	maxRepos := 25
	if len(repos) > maxRepos {
		repos = repos[:maxRepos]
		repoList = append(repoList, "Showing first "+utils.IntToString(maxRepos)+" repositories:")
	}
	for _, repo := range repos {
		repoList = append(repoList, "- ["+repo.GetName()+"]("+repo.GetHTMLURL()+") - ⭐ "+utils.IntToString(repo.GetStargazersCount()))
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Repositories for " + args[0],
		Description: strings.Join(repoList, "\n"),
		URL:         "https://github.com/" + args[0],
		Color:       0x00ff00,
	})
}
