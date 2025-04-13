package utils
import "github.com/bwmarrin/discordgo"
func DiscordError(s *discordgo.Session, m *discordgo.MessageCreate, errorMessage string) {
	_ = s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
	_, _ = s.ChannelMessageSend(m.ChannelID, "Error: "+errorMessage)
}
