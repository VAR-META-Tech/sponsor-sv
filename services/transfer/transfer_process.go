package transfer

import (
	"encoding/json"
	"fmt"
	"log"
	"sponsor-sv/models"
	"sponsor-sv/services/account"

	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// This TransferProcess reconstruct the std.TX{}, feed the ExecuteSponsorTransaction() with account number and sequence number
// Returns the encoded/marshalled ResultBroadcastTxCommit{} with hashTx inside :)
func TransferProcess(cli *gnoclient.Client, msg std.Tx) (maybeTxHash []byte, err error) {
	// Check for length
	if !validSponsorLen(msg) {
		return []byte{}, models.ErrInvalidLenMsg
	}

	// Query for account number and account sequence
	baseInfo, err := cli.Signer.Info()
	if err != nil {
		log.Println("Error: ", err)
	}
	sAddr := baseInfo.GetAddress()
	log.Printf("======= sAddr: %s\n", sAddr.String())
	sBaseAcc, err := account.GetAccountBaseWithAddr(cli, sAddr.String())
	log.Printf("======= sAccountNumb: %v\n", sBaseAcc.GetSequence())
	log.Printf("======= sSequence: %v\n", sBaseAcc.GetAccountNumber())

	stdSigs := msg.GetSignatures()
	log.Printf("======= len stdSig: %v\n", len(stdSigs))
	if err != nil {
		log.Println("Error getaccount: ", err)
		return []byte{}, err
	}
	// add the message into Tx
	nullSignature := std.Signature{}
	newMsgSigs := append([]std.Signature{nullSignature}, msg.Signatures...)
	msg.Signatures = newMsgSigs

	// valid check, if not valid then no execute
	if checkValidCall(cli, msg) {
		log.Printf("======= message before execute: %+v\n", msg)

		// Execute the tx
		resultExecute, err := cli.ExecuteSponsorTransaction(msg, sBaseAcc.GetAccountNumber(), sBaseAcc.GetSequence())
		if err != nil {
			log.Println("Error execute: ", err)
			return []byte{}, err
		}

		// this will check for msg.Deliver.IsErr or not
		result := rebuildMessage(resultExecute)
		if !result.Success {
			log.Println("======= failed on execute message")
			return []byte{}, models.ErrOnExecuteMsg
		}
		// Encode the result
		resultEncoded, err := json.Marshal(&result)
		if err != nil {
			log.Println("Error: ", err)
			return []byte{}, err
		}

		return resultEncoded, nil
	}
	return []byte{}, models.ErrInvalidCallerMsg
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

// currently just support checking vm_msgCall here. Add your checks here to gurantee that your sponsor will not process bad call to chain
func checkValidCall(cli *gnoclient.Client, txs std.Tx) bool {
	for i, tx := range txs.GetMsgs() {
		log.Println("====== checking for message ", i, " with msgType: ", tx.Type())
		switch tx.Type() {
		case "no_op":
			log.Println("====== skip checking msg noop")
		case "exec":
			realmToCheck := tx.(vm.MsgCall).PkgPath
			argsCalled := tx.(vm.MsgCall).Args
			funcCalled := tx.(vm.MsgCall).Func
			qevalExpression := fmt.Sprintf("%s(\"%s\")", funcCalled, argsCalled[0])
			log.Printf("====== checking calling to /r %s with args %s\n", realmToCheck, qevalExpression)
			_, _, err := cli.QEval(realmToCheck, qevalExpression)
			if err != nil {
				log.Println("====== err checking ", err.Error())
				return false
			}
			log.Println("====== checking done, msg ok")
			return true
		default:
			// continue
		}
	}
	return true
}
