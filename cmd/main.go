package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
	"github.com/iButcat/golady/api/v1"
	"github.com/iButcat/golady/config"
	"github.com/iButcat/golady/pkg/bot"
	"github.com/iButcat/golady/pkg/db/sqlite"
)
func init() {
	_ = gotenv.Load()
}
func main() {
	cfg := config.NewConfig()
	storage := sqlite.New(cfg.DBPath)
	bot := bot.NewBot(cfg.DiscordToken, storage, cfg.GithubToken)
	err := bot.DiscordSession.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer bot.DiscordSession.Close()
	botWrapper := &botWrapper{
		bot: bot,
	}
	bot.DiscordSession.AddHandler(botWrapper.messageCreate)
	botWrapper.registerCmds()
	router := api.NewAPI(storage, bot.DiscordSession)
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("received signal: %s", <-c)
	}()
	go func() {
		port := ":" + cfg.ServerPort
		log.Printf("Starting server on port %s", port)
		errs <- http.ListenAndServe(port, router)
	}()
	log.Println("Bot is now running. Press CTRL-C to exit.")
	log.Fatalf("Exiting: %v", <-errs)
}
type botWrapper struct {
	bot *bot.Bot
}
func (b *botWrapper) registerCmds() {
	b.bot.CmdList.NewCommand("help", "Display available commands", b.bot.CmdList.Help, false)
	b.bot.CmdList.NewCommand("config", "Configure the bot for this server", b.bot.InitConfig, false)
	b.bot.CmdList.NewCommand("subscribe", "Subscribe to a GitHub repository", b.bot.GithubSubscriptions.Subscribe, false)
	b.bot.CmdList.NewCommand("unsubscribe", "Unsubscribe from a GitHub repository", b.bot.GithubSubscriptions.Unsubscribe, false)
	b.bot.CmdList.NewCommand("subscriptions", "List all subscriptions", b.bot.GithubSubscriptions.ListSubscriptions, false)
	b.bot.CmdList.NewCommand("issue", "Get information about an issue", b.bot.GithubIssues.GetFromRepoName, false)
	b.bot.CmdList.NewCommand("repo", "Get information about a repository", b.bot.GithubRepos.Get, false)
	b.bot.CmdList.NewCommand("webhook-setup", "Get instructions for setting up webhooks", b.bot.GithubSetup.SendWebhookInstructions, false)
}
func (b *botWrapper) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!") {
		cmdName := strings.Split(m.Content, " ")[0][1:]
		for _, cmd := range b.bot.CmdList.List {
			if cmd.Name == cmdName {
				cmd.Def(s, m)
				break
			}
		}
	}
}
