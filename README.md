# GitHub Notify

A Discord bot that sends notifications for GitHub events to your Discord server.

> **⚠️ WORK IN PROGRESS**: This project is currently being refactored with improved design patterns and structure. It was initially started several years ago and is now being modernized. Features may change and documentation might be incomplete.

![GitHub Notify Screenshot](./docs/screenshot.png)

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

## Documentation

Visit our [GitHub Pages documentation](https://yourusername.github.io/github-notify/) for complete instructions on:
- Setting up the bot
- Available commands
- Self-hosting options
- Troubleshooting

## Installation

See our [installation guide](https://yourusername.github.io/github-notify/installation) for detailed instructions.

### Prerequisites

- Go 1.22 or higher
- A Discord Bot Token ([Create a Discord Bot](https://discord.com/developers/applications))
- A GitHub Personal Access Token ([Create a GitHub Token](https://github.com/settings/tokens))

### Quick Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/github-notify.git
cd github-notify

# Configure environment
cp .env.example .env
# Edit .env with your tokens

# Build and run
go build -o github-notify ./cmd/main.go
./github-notify
```

## Docker

```bash
docker run -d \
  --name github-notify \
  -p 8080:8080 \
  -v /path/to/data:/app/data \
  -e DISCORD_TOKEN=your_discord_bot_token \
  -e GITHUB_TOKEN=your_github_personal_access_token \
  yourusername/github-notify:latest
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [DiscordGo](https://github.com/bwmarrin/discordgo) - Go library for Discord
- [Go-GitHub](https://github.com/google/go-github) - Go library for GitHub API