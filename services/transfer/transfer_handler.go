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
	_ "github.com/gnolang/gno/tm2/pkg/std"
	"github.com/golang/protobuf/proto"
)

/*
This handler will decode message into encodedTransaction and then decode recieved data into std.Tx{} with protobuf decoder.
expectation: recieving a struct like this:
type MsgNoop struct {
 	EncodedTransaction string `json:"encoded-transaction"`
}
*/

type MsgFromFE struct {
	EncodedTransaction string `json:"encoded-transaction"`
}

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

	msg := MsgFromFE{}
	errMarshal := json.Unmarshal(bodyBytes, &msg)
	if errMarshal != nil {
		prob := models.ProblemDetail{
			Error:   "can not decode transfer message",
			Details: errMarshal.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	log.Printf("\nmsg from FE: %+v\n", msg)
	cli := gclient.GetClient()

	// sponsorTx := SponsorMsg()
	// x, _ := amino.MarshalJSON(sponsorTx)
	// log.Printf("predefined sponsorTx: %s\n", x)

	// decode the encodedTransaction into std.Tx{} with proto
	// need this definition from proto files
	msgProto := models.Tx{}
	//decode base64
	decodeBytes, errBase64 := base64.RawStdEncoding.DecodeString(msg.EncodedTransaction)

	if errBase64 != nil {
		prob := models.ProblemDetail{
			Error:   "can not decode base64",
			Details: errBase64.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	log.Printf("\n\nMsg in Bytes: %v\n", decodeBytes)

	errDecodeProto := proto.Unmarshal(decodeBytes, &msgProto)
	if errDecodeProto != nil {
		log.Println("erro decode from proto: ", errDecodeProto)

		prob := models.ProblemDetail{
			Error:   "can not decode proto from encoded transaction",
			Details: errDecodeProto.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, prob)
		return
	}
	log.Printf("\nrecieved msg decoded: %+v\n", msgProto)
	result, err := TransferProcess(cli, msgProto)
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
