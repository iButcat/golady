package db

import (
	"time"
)

type Storage struct {
	ServerConfig ServerConfig
	Webhook      Webhook
}
type ServerConfig interface {
	Create(configServer *ConfigServer) (*ConfigServer, error)
	Update(configServer *ConfigServer) (*ConfigServer, error)
	GetByServerID(serverID string) (*ConfigServer, error)
	Delete(serverID string) error
}
type Webhook interface {
	Subscribe(webhookData *WebhookData) error
	Unsubscribe(serverID, repository string) error
	GetByServerID(serverID string) ([]*WebhookData, error)
	GetByRepository(repository string) ([]*WebhookData, error)
}
type ConfigServer struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ServerID    string `gorm:"uniqueIndex"`
	ChannelID   string
	GithubToken string
}
type WebhookData struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ServerID   string
	ChannelID  string
	Repository string
	Owner      string
	Events     string
}
