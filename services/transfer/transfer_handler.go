package transfer

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sponsor-sv/models"
	"sponsor-sv/services/gclient"

	"github.com/gin-gonic/gin"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/std"
)

/*
This handler will decode message into encodedTransaction and then decode recieved data into std.Tx{} with protobuf decoder.
expectation: recieving a struct like this:
type MsgNoop struct {
 	EncodedTransaction string `json:"encoded-transaction"`
}
*/

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
	// msg2 := std.Tx{}
	// errMarshal := amino.UnmarshalJSON(bodyBytes, &msg)

	msg := models.MsgFromFE{}
	errMarshal := json.Unmarshal(bodyBytes, &msg)
	if errMarshal != nil {
		prob := models.ProblemDetail{
			Error:   "can not decode transfer message",
			Details: errMarshal.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	cli := gclient.GetClient()

	sponsorTx := std.Tx{}
	//decode base64
	decodeBytes, errBase64 := base64.StdEncoding.DecodeString(msg.EncodedTransaction)
	if errBase64 != nil {
		log.Println("error from decode base64: ", errBase64.Error())
		prob := models.ProblemDetail{
			Error:   "can not decode base64",
			Details: errBase64.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}

	// amino supports for Unmarshal proto messages
	errDecodeProto := amino.Unmarshal(decodeBytes, &sponsorTx)
	if errDecodeProto != nil {
		log.Println("error decode from proto: ", errDecodeProto)
		prob := models.ProblemDetail{
			Error:   "can not decode proto from encoded transaction",
			Details: errDecodeProto.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	result, err := TransferProcess(cli, sponsorTx)
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
