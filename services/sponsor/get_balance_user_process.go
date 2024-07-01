package sponsor

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	core_types "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
)

func GetBalanceOfUser(cli *gnoclient.Client, targetUser string) (string, error) {
	// Evaluate expression
	qevalRes, _, err := cli.QEval("gno.land/r/varmeta/vmt721", "BalanceOf(\"g162jgpk4740r6a7g53cgz9ahxqtyuekgqchw6w9\")")
	if err != nil {
		return "", err
	}
	return qevalRes, nil
}

// Parse the result to get the endpoint
func ParseResultToEndpoint(res core_types.ResultBroadcastTxCommit) string {
	return "localhost:8765"
}
