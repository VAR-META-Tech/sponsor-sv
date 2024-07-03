package sponsor

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	// Now we just take sponsor address from query
	targetUser := c.Param("id")
	client := gclient.GetClient()
	result, err := GetUserBalance(client, targetUser)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "error query to chain",
		}
		c.JSON(http.StatusInternalServerError, prob)
		return
	}
	c.JSON(http.StatusOK, result)
}
