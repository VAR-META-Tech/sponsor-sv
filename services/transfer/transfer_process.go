package transfer

import (
	"encoding/json"
	"errors"
	"log"
	"sponsor-sv/models"
	"sponsor-sv/services/account"

	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/std"
)

var errInvalidLenMsg error = errors.New("invalid sponsor message length")

// This TransferProcess reconstruct the std.TX{}, feed the ExecuteSponsorTransaction() with account number and sequence number
// Returns the encoded/marshalled ResultBroadcastTxCommit{} with hashTx inside :)
func TransferProcess(cli *gnoclient.Client, msgInput models.Tx) (maybeTxHash []byte, err error) {
	msg := std.Tx{}
	// need to mapping input msg and real msg
	// msg := msgInput
	// Check for length
	if !validSponsorLen(msg) {
		return []byte{}, errInvalidLenMsg
	}

	// Query for account number and account sequence
	baseInfo, err := cli.Signer.Info()
	log.Println("Error: ", err)
	sAddr := baseInfo.GetAddress()
	log.Printf("======= sAddr: %s\n", sAddr.String())
	sBaseAcc, err := account.GetAccountBaseWithAddr(cli, sAddr.String())
	log.Printf("======= sAccountNumb: %v\n", sBaseAcc.GetSequence())
	log.Printf("======= sSequence: %v\n", sBaseAcc.GetAccountNumber())
	stdSigs := msg.GetSignatures()
	lenSig := msg.GetSigners()
	log.Printf("======= len stdSig: %v\n", len(stdSigs))
	log.Printf("======= len signer`: %v\n", len(lenSig))
	if err != nil {
		log.Println("Error getaccount: ", err)
		return []byte{}, err
	}
	log.Printf("message before execute: %+v\n", msg)
	// Execute the commit
	resultExecute, err := cli.ExecuteSponsorTransaction(msg, sBaseAcc.AccountNumber, sBaseAcc.Sequence)
	if err != nil {
		log.Println("Error execute: ", err)
		return []byte{}, err
	}
	result := rebuildMessage(resultExecute)

	// Encode the result
	resultEncoded, err := json.Marshal(&result)
	if err != nil {
		log.Println("Error: ", err)
		return []byte{}, err
	}
	return resultEncoded, nil
}

func validSponsorLen(msg std.Tx) bool {
	return len(msg.Msgs) >= 2
}

func rebuildMessage(msg *ctypes.ResultBroadcastTxCommit) models.TransferResult {
	if msg.DeliverTx.IsErr() {
		return models.TransferResult{
			Success: false,
		}
	}
	return models.TransferResult{
		Success:     true,
		MessageHash: msg.Hash,
	}
}
