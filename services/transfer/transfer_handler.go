package transfer

import (
	"io"
	"net/http"
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"

	"github.com/gin-gonic/gin"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/std"
)

func TransferHandler(c *gin.Context) {
	// Get the body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not get transfer message",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}

	// Decode the message with amino json type
	msg := std.Tx{}
	errMarshal := amino.UnmarshalJSON(bodyBytes, &msg)
	if errMarshal != nil {
		prob := models.ProblemDetail{
			Error:   "can not decode transfer message",
			Details: errMarshal.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	cli := gclient.GetClient()
	// sponsorTx := SponsorMsg()
	// x, _ := amino.MarshalJSON(sponsorTx)
	// log.Printf("predefined sponsorTx: %s\n", x)

	result, err := TransferProcess(cli, msg)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   "can not process transfering",
			Details: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	c.JSON(http.StatusOK, result)
}
