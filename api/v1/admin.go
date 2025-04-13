package api
import "github.com/gin-gonic/gin"
func (api *API) AdminDashboard(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": "yes"})
}
