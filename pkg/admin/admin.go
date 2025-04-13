package admin
import (
	"strconv"
	"github.com/iButcat/golady/pkg/utils"
	"github.com/bwmarrin/discordgo"
)
type (
	Admin interface {
		Delete(s *discordgo.Session, m *discordgo.MessageCreate)
	}
	admin struct {
	}
)
func NewAdmin() Admin {
	return &admin{}
}
func (a *admin) Delete(s *discordgo.Session, m *discordgo.MessageCreate) {
	args, nb := utils.ParseArgs(m.Content)
	if nb != 1 {
		utils.DiscordError(s, m, "wrong number of arguments")
		return
	}
	_, err := strconv.Atoi(args[0])
	if err != nil {
		utils.DiscordError(s, m, "while converting argument to int")
	}
}
