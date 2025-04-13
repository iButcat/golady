package repository

import (
	"errors"

	"github.com/iButcat/golady/pkg/db"
	"gorm.io/gorm"
)

type configServer struct {
	client *gorm.DB
}

func NewConfigServer(client *gorm.DB) db.ServerConfig {
	return &configServer{
		client: client,
	}
}
func (c *configServer) configExists(serverID string) bool {
	var count int64
	c.client.Model(&db.ConfigServer{}).Where("server_id = ?", serverID).Count(&count)
	return count > 0
}
func (c *configServer) Create(configServer *db.ConfigServer) (*db.ConfigServer, error) {
	if c.configExists(configServer.ServerID) {
		return nil, errors.New("config already exists for this server")
	}
	if err := c.client.Create(configServer).Error; err != nil {
		return nil, err
	}
	return configServer, nil
}
func (c *configServer) Update(configServer *db.ConfigServer) (*db.ConfigServer, error) {
	if err := c.client.Model(&db.ConfigServer{}).Where("server_id = ?", configServer.ServerID).Updates(configServer).Error; err != nil {
		return nil, err
	}
	return configServer, nil
}
func (c *configServer) GetByServerID(serverID string) (*db.ConfigServer, error) {
	var configServer db.ConfigServer
	if err := c.client.Where("server_id = ?", serverID).First(&configServer).Error; err != nil {
		return nil, err
	}
	return &configServer, nil
}
func (c *configServer) Delete(serverID string) error {
	if err := c.client.Where("server_id = ?", serverID).Delete(&db.ConfigServer{}).Error; err != nil {
		return err
	}
	return nil
}
