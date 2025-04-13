package repository

import (
	"github.com/iButcat/golady/pkg/db"
	"gorm.io/gorm"
)

type webhook struct {
	client *gorm.DB
}

func NewWebhook(client *gorm.DB) db.Webhook {
	return &webhook{
		client: client,
	}
}
func (w *webhook) Subscribe(webhookData *db.WebhookData) error {
	var count int64
	w.client.Model(&db.WebhookData{}).
		Where("server_id = ? AND repository = ? AND owner = ?",
			webhookData.ServerID, webhookData.Repository, webhookData.Owner).
		Count(&count)
	if count > 0 {
		return w.client.Model(&db.WebhookData{}).
			Where("server_id = ? AND repository = ? AND owner = ?",
				webhookData.ServerID, webhookData.Repository, webhookData.Owner).
			Updates(webhookData).Error
	}
	return w.client.Create(webhookData).Error
}
func (w *webhook) Unsubscribe(serverID, repository string) error {
	return w.client.Where("server_id = ? AND repository = ?", serverID, repository).
		Delete(&db.WebhookData{}).Error
}
func (w *webhook) GetByServerID(serverID string) ([]*db.WebhookData, error) {
	var webhooks []*db.WebhookData
	if err := w.client.Where("server_id = ?", serverID).Find(&webhooks).Error; err != nil {
		return nil, err
	}
	return webhooks, nil
}
func (w *webhook) GetByRepository(repository string) ([]*db.WebhookData, error) {
	var webhooks []*db.WebhookData
	if err := w.client.Where("repository = ?", repository).Find(&webhooks).Error; err != nil {
		return nil, err
	}
	return webhooks, nil
}
