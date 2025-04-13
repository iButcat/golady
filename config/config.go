package config
import (
	"os"
)
type Config struct {
	DiscordToken string
	GithubToken  string
	DBPath string
	ServerPort string
}
func NewConfig() (config Config) {
	config = Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		GithubToken:  os.Getenv("GITHUB_TOKEN"),
		DBPath:       os.Getenv("DB_PATH"),
		ServerPort:   os.Getenv("SERVER_PORT"),
	}
	if config.DBPath == "" {
		config.DBPath = "golady.db"
	}
	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}
	return config
}
