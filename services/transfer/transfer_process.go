package transfer

import (
	"sponsor-sv/models"
	"encoding/json"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// This TransferProcess reconstruct the std.TX{}, feed the ExecuteSponsorTransaction() with account number and sequence number
// Returns the encoded/marshalled ResultBroadcastTxCommit{} with hashTx inside :)
func TransferProcess(cli *gnoclient.Client, msg models.Transaction) (maybeTxHash []byte, err error) {
	fakeTx := std.Tx{}
	fAccNumb := uint64(56)
	fSeqNumb := uint64(0)

	// Broadcast the commit
	result, err := cli.ExecuteSponsorTransaction(fakeTx, fAccNumb, fSeqNumb)
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
