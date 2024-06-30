package account

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountHandler(c *gin.Context) {
	targetAddr := c.Request.URL.Query().Get("addr")
	cli := gclient.GetClient()
	baseAccount, err := GetAccountBaseWithAddr(cli, targetAddr)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not query account",
		}
		c.JSON(http.StatusInternalServerError, prob)
	}
	accToReponse := models.AccountInfo{
		Addr:           baseAccount.Address.String(),
		Balance:        baseAccount.Coins.String(),
		AccountNumber:  baseAccount.AccountNumber,
		SequenceNumber: baseAccount.Sequence,
	}
	c.JSON(http.StatusOK, accToReponse)
}
