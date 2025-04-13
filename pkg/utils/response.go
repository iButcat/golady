package utils
import "github.com/bwmarrin/discordgo"
func DiscordResponse(s *discordgo.Session, m *discordgo.MessageCreate, response string) {
	_ = s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	_, _ = s.ChannelMessageSend(m.ChannelID, response)
}
