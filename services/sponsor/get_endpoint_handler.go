package sponsor

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEndpoint(c *gin.Context) {
	// Now we just take sponsor address from query
	targetSponsor := c.Query("addr")
	client := gclient.GetClient()
	sponsorEndpoint, err := GetEndpointProcess(client, targetSponsor, "gno.land/r/thinhnx/sponsor_realm", "GetSponsorEndpoint")
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "Can not query to chain",
		}
		c.JSON(http.StatusInternalServerError, prob)
	}
	c.JSON(http.StatusOK, sponsorEndpoint)
}
