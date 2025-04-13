package api
import (
	"net/http"
	"github.com/iButcat/golady/pkg/db"
	"github.com/iButcat/golady/pkg/github"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)
type API struct {
	storage        *db.Storage
	discord        *discordgo.Session
	webhookHandler *github.WebhookHandler
}
func NewAPI(storage *db.Storage, discord *discordgo.Session) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api := &API{
		storage:        storage,
		discord:        discord,
		webhookHandler: github.NewWebhookHandler(discord, storage),
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":        "GitHub Notify",
			"description": "Discord bot for GitHub notifications",
			"status":      "running",
			"version":     "1.0.0",
		})
	})
	v1 := router.Group("/api/v1")
	{
		v1.POST("/webhook", api.handleWebhook)
		v1.GET("/status", api.getStatus)
	}
	return router
}
func (api *API) handleWebhook(c *gin.Context) {
	api.webhookHandler.HandleWebhook(c.Writer, c.Request)
}
func (api *API) getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "1.0.0",
	})
}
