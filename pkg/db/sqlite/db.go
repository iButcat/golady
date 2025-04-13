package sqlite

import (
	"log"

	"github.com/iButcat/golady/pkg/db"
	"github.com/iButcat/golady/pkg/db/sqlite/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dbPath string) *db.Storage {
	client, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err := client.AutoMigrate(&db.ConfigServer{}, &db.WebhookData{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	return &db.Storage{
		ServerConfig: repository.NewConfigServer(client),
		Webhook:      repository.NewWebhook(client),
	}
}
