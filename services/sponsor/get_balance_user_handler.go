package sponsor

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	// Now we just take sponsor address from query
	targetUser := c.Query("addr")
	client := gclient.GetClient()
	result, err := GetBalanceOfUser(client, targetUser)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "Can not query to chain",
		}
		c.JSON(http.StatusInternalServerError, prob)
		return
	}
	c.JSON(http.StatusOK, result)
}
