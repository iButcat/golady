package bot
import (
	"fmt"
	"github.com/iButcat/golady/pkg/cmd"
	"github.com/iButcat/golady/pkg/db"
	"github.com/iButcat/golady/pkg/github"
	"github.com/iButcat/golady/pkg/utils"
	"github.com/bwmarrin/discordgo"
)
type Bot struct {
	DiscordSession      *discordgo.Session
	GithubIssues        github.Issues
	GithubRepos         github.Repositories
	GithubSubscriptions *github.SubscriptionHandler
	GithubSetup         github.Setup
	Storage             *db.Storage
	CmdList             *cmd.Commands
}
func NewBot(
	discordToken string,
	storage *db.Storage,
	ghToken string) *Bot {
	discordSession, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		panic(fmt.Errorf("error creating Discord session: %v", err))
	}
	client := github.NewGitHubClient(ghToken)
	githubIssues := github.NewIssues(client)
	githubRepo := github.NewRepositories(client)
	githubSubscriptions := github.NewSubscriptionHandler(storage)
	githubSetup := github.NewSetup(ghToken)
	return &Bot{
		DiscordSession:      discordSession,
		GithubIssues:        githubIssues,
		GithubRepos:         githubRepo,
		GithubSubscriptions: githubSubscriptions,
		GithubSetup:         githubSetup,
		Storage:             storage,
		CmdList:             cmd.NewCommands(),
	}
}
func (b *Bot) InitConfig(s *discordgo.Session, m *discordgo.MessageCreate) {
	args, count := utils.ParseArgs(m.Content)
	if count != 1 {
		utils.DiscordError(s, m, "Usage: !config YOUR_GITHUB_TOKEN")
		return
	}
	githubToken := args[0]
	configServer := &db.ConfigServer{
		ServerID:    m.GuildID,
		ChannelID:   m.ChannelID,
		GithubToken: githubToken,
	}
	_, err := b.Storage.ServerConfig.Create(configServer)
	if err != nil {
		utils.DiscordError(s, m, "Error creating config: "+err.Error())
		return
	}
	utils.DiscordResponse(s, m, "Configuration created successfully! You can now subscribe to repositories using !subscribe owner/repo")
}
func (b *Bot) DeleteConfig(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := b.Storage.ServerConfig.Delete(m.GuildID)
	if err != nil {
		utils.DiscordError(s, m, "Error deleting config: "+err.Error())
		return
	}
	utils.DiscordResponse(s, m, "Configuration deleted successfully")
}
func (b *Bot) GetConfig(s *discordgo.Session, m *discordgo.MessageCreate) {
	srvCfg, err := b.Storage.ServerConfig.GetByServerID(m.GuildID)
	if err != nil {
		utils.DiscordError(s, m, err.Error())
	}
	fmt.Print(srvCfg)
}
