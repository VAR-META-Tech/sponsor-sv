package account

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountHandler(c *gin.Context) {
	targetAddr := c.Request.URL.Query().Get("addr")
	if targetAddr == "" {
		prob := models.ProblemDetail{
			Error: "bad query",
			Details: "can not get address query",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	cli := gclient.GetClient()
	baseAccount, err := GetAccountBaseWithAddr(cli, targetAddr)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not query account",
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, prob)
		return
	}
	accToReponse := models.AccountInfo{
		Addr:           baseAccount.Address.String(),
		Balance:        baseAccount.Coins.String(),
		AccountNumber:  baseAccount.AccountNumber,
		SequenceNumber: baseAccount.Sequence,
	}
	c.JSON(http.StatusOK, accToReponse)
}
