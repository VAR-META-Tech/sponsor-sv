package transfer

import (
	"io"
	"log"
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

	// Decode the message ?
	msg := std.Tx{}
	errMarshal := amino.UnmarshalJSON(bodyBytes, &msg)
	if errMarshal != nil {
		prob := models.ProblemDetail{
			Error:   errMarshal.Error(),
			Details: "can not decode transfer message",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	log.Printf("Debug: %v\n", msg)
	cli := gclient.GetClient()
	result, err := TransferProcess(cli, msg)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not process transfering",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	c.JSON(http.StatusOK, result)
}
