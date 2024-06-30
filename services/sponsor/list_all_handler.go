package sponsor

import (
	"sponsor-sv/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllHandler(c *gin.Context) {
	listSponsor := []models.Sponsor{
		{
			Name:    "VarMeta",
			Address: "g1aksdjnsajdnskjdn",
		},
	}
	c.JSON(http.StatusOK, listSponsor)
}
