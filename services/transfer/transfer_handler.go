package transfer

import (
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransferHandler(c *gin.Context) {
	// Get the body

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not get transfer message",
		}
		c.JSON(http.StatusBadRequest, prob)
	}

	// Decode the message ?
	msg := models.Transaction{}
	errMarshal := json.Unmarshal(bodyBytes, &msg)
	if errMarshal != nil {
		log.Println("Error get sign msg: ", err)
		prob := models.ProblemDetail{
			Error:   errMarshal.Error(),
			Details: "can not decode transfer message",
		}
		c.JSON(http.StatusBadRequest, prob)
		return
	}
	cli := gclient.GetClient()
	result, err := TransferProcess(cli, msg)
	if err != nil {
		prob := models.ProblemDetail{
			Error:   err.Error(),
			Details: "can not process transfering",
		}
		c.JSON(http.StatusBadRequest, prob)
		return
	}
	c.JSON(http.StatusOK, result)
}
