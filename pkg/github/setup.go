package github
import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v60/github"
	"golang.org/x/oauth2"
)
type Setup interface {
	SendWebhookInstructions(s *discordgo.Session, m *discordgo.MessageCreate)
}
type setup struct {
	ghClient *github.Client
}
func NewSetup(ghToken string) Setup {
	return &setup{
		ghClient: NewGitHubClient(ghToken),
	}
}
func (st *setup) SendWebhookInstructions(s *discordgo.Session, m *discordgo.MessageCreate) {
	instructions := `
To set up GitHub notifications for this channel, follow these steps:
1. Go to your GitHub repository
2. Click on "Settings" > "Webhooks" > "Add webhook"
3. For Payload URL, enter: http://your-server-url/api/v1/webhook
4. Content type: application/json
5. Select events you want to receive (at least Issues and Pull requests)
6. Click "Add webhook"
Alternatively, use the following command to subscribe to a repository:
!subscribe owner/repo
This will automatically receive notifications for new issues and pull requests.
`
	_, _ = s.ChannelMessageSend(m.ChannelID, instructions)
}
func NewGitHubClient(ghToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
