package cmd
import (
	"strings"
	"github.com/iButcat/golady/pkg/utils"
	"github.com/bwmarrin/discordgo"
)
type Command struct {
	Name        string
	Description string
	Example     string
	Def         func(s *discordgo.Session, m *discordgo.MessageCreate)
	IsAdminCmd  bool
	IsPremium   bool
}
type Commands struct {
	List []Command
}
func NewCommands() *Commands {
	return &Commands{
		List: make([]Command, 0),
	}
}
func (c *Commands) NewCommand(name string,
	dsc string,
	def func(s *discordgo.Session, m *discordgo.MessageCreate),
	isAdminCmd bool) {
	c.appendCommand(Command{
		Name:        name,
		Description: dsc,
		Def:         def,
		IsAdminCmd:  isAdminCmd,
	})
}
func (c *Commands) appendCommand(cmd Command) {
	c.List = append(c.List, cmd)
}
func (c *Commands) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	args, nb := utils.ParseArgs(m.Content)
	cmdsHelp := make([]string, 0)
	if nb == 1 {
		for _, cmd := range c.List {
			if args[0] == cmd.Name {
				utils.DiscordResponse(s, m,
					cmd.Name+" "+cmd.Description+" "+cmd.Example)
			}
		}
	} else {
		utils.DiscordResponse(s, m,
			"help <command name> for a specific command description")
		for _, cmd := range c.List {
			cmdsHelp = append(cmdsHelp,
				cmd.Name+" "+cmd.Description+" "+cmd.Example)
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "List of commands",
			Description: strings.Join(cmdsHelp, "\n"),
		})
	}
}
