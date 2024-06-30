package sponsor

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	core_types "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
)

func GetEndpointProcess(cli *gnoclient.Client, targetSponsor string, rPath string, fnName string) (string, error) {
	//config - still hardcode

	txConfig := gnoclient.BaseTxCfg{
		GasFee:         "100000ugnot",
		GasWanted:      100000,
		AccountNumber:  56, //?
		SequenceNumber: 0,
	}
	callArgs := []string{targetSponsor}
	msgGetEndpoint := gnoclient.MsgCall{
		PkgPath:  rPath,
		FuncName: fnName,
		Args:     callArgs,
		Send:     "", //?
	}

	// Call to the realm
	res, err := cli.Call(txConfig, msgGetEndpoint)
	if err != nil {
		return "", err
	}
	endpointString := ParseResultToEndpoint(*res)
	return endpointString, nil
}

// Parse the result to get the endpoint
func ParseResultToEndpoint(res core_types.ResultBroadcastTxCommit) string {
	return "localhost:8765"
}
