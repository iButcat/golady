# Installation Guide

This guide helps you set up your own instance of GitHub Notify.

## Prerequisites

- Go 1.22 or higher
- A Discord Bot Token ([Create a Discord Bot](https://discord.com/developers/applications))
- A GitHub Personal Access Token ([Create a GitHub Token](https://github.com/settings/tokens))

## Method 1: Docker (Recommended)

The easiest way to run GitHub Notify is using Docker:

```bash
# Pull the image
docker pull yourusername/github-notify:latest

# Run the container
docker run -d \
  --name github-notify \
  -p 8080:8080 \
  -v /path/to/data:/app/data \
  -e DISCORD_TOKEN=your_discord_bot_token \
  -e GITHUB_TOKEN=your_github_personal_access_token \
  yourusername/github-notify:latest
```

## Method 2: Manual Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/github-notify.git
cd github-notify
```

### 2. Create a `.env` file:

```
DISCORD_TOKEN=your_discord_bot_token
GITHUB_TOKEN=your_github_personal_access_token
DB_PATH=golady.db
SERVER_PORT=8080
```

### 3. Build and run:

```bash
go build -o github-notify ./cmd/main.go
./github-notify
```

## Setting Up the Discord Bot

1. Go to [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a new application and set up a bot
3. Enable the "Message Content Intent" under Bot Settings
4. Copy your bot token and add it to the `.env` file
5. Generate an invite link with the "bot" scope and permissions: 
   - Read Messages/View Channels
   - Send Messages
   - Embed Links
   - Read Message History
6. Invite the bot to your server using the link

## Configuration Options

| Environment Variable | Description | Default |
|----------------------|-------------|---------|
| `DISCORD_TOKEN` | Your Discord Bot token | (required) |
| `GITHUB_TOKEN` | Your GitHub Personal Access Token | (required) |
| `DB_PATH` | Path to SQLite database file | `golady.db` |
| `SERVER_PORT` | Port for the webhook server | `8080` |

## Troubleshooting

### Bot doesn't respond to commands
- Ensure the bot has proper permissions in the channel
- Verify the "Message Content Intent" is enabled in the Discord Developer Portal
- Check logs for any errors

### Not receiving webhook events
- Make sure your webhook URL is publicly accessible
- Verify the webhook is properly configured in GitHub
- Check logs for webhook processing errors 