package transfer

import (
	"encoding/json"
	"log"
	"sponsor-sv/services/account"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// This TransferProcess reconstruct the std.TX{}, feed the ExecuteSponsorTransaction() with account number and sequence number
// Returns the encoded/marshalled ResultBroadcastTxCommit{} with hashTx inside :)
func TransferProcess(cli *gnoclient.Client, msg std.Tx) (maybeTxHash []byte, err error) {
	fakeTx := std.Tx{}
	baseInfo, _ := cli.Signer.Info()
	sAddr := baseInfo.GetAddress()
	log.Printf("sAddr: %s\n", sAddr.String())
	sBaseAcc, err := account.GetAccountBaseWithAddr(cli, sAddr.String())
	if err != nil {
		return []byte{}, err
	}

	// Broadcast the commit
	result, err := cli.ExecuteSponsorTransaction(fakeTx, sBaseAcc.Sequence, sBaseAcc.Sequence)
	if err != nil {
		return []byte{}, err
	}
	// Encode the result
	resultEncoded, err := json.Marshal(&result)
	if err != nil {
		return []byte{}, err
	}
	return resultEncoded, nil
}
