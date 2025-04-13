# GitHub Notify

A Discord bot that sends notifications for GitHub events to your Discord server.

## Features

- 🔔 **Real-time notifications**: Get instant notifications for GitHub events in your Discord channels
- 📊 **Issue tracking**: Receive notifications for new issues, comments, and status changes
- 🔄 **Pull request monitoring**: Be notified of new PRs, reviews, and merges
- 🔗 **Simple subscription system**: Use Discord commands to subscribe to repositories
- 🛠️ **Easy setup**: Set up with just a few commands

## Quick Start

1. [Invite the bot](https://discord.com/api/oauth2/authorize?client_id=YOUR_CLIENT_ID&permissions=8&scope=bot) to your server
2. Use `!subscribe owner/repo` to start receiving notifications
3. That's it! You'll now receive notifications for new issues and PRs

## Commands

| Command | Description |
|---------|-------------|
| `!help` | Show available commands |
| `!config [token]` | Configure your server with a GitHub token |
| `!subscribe owner/repo` | Subscribe to a repository |
| `!unsubscribe owner/repo` | Unsubscribe from a repository |
| `!subscriptions` | List all active subscriptions |
| `!issue owner repo number` | Get details about a specific issue |
| `!repo owner repo` | Get information about a repository |
| `!webhook-setup` | Get instructions for setting up webhooks |

## Setting Up Webhooks

For real-time notifications, you can set up GitHub webhooks:

1. Go to your GitHub repository settings
2. Navigate to "Webhooks" > "Add webhook"
3. Enter your webhook URL: `http://your-server/api/v1/webhook`
4. Set content type to `application/json`
5. Select the events you want to receive notifications for
6. Save the webhook

## Self-Hosting

See our [installation guide](installation.md) for detailed instructions on hosting your own instance.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/yourusername/github-notify/blob/main/LICENSE) file for details. 