package sponsor

import (
	"fmt"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	core_types "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
)

func GetUserBalance(cli *gnoclient.Client, targetUser string) (string, error) {
	// Evaluate expression
	qevalExpression := fmt.Sprintf("BalanceOf(\"%s\")", targetUser)
	qevalRes, _, err := cli.QEval("gno.land/r/varmeta/vmt721", qevalExpression)
	if err != nil {
		return "", err
	}
	return qevalRes, nil
}

// Parse the result to get the endpoint
func ParseResultToEndpoint(res core_types.ResultBroadcastTxCommit) string {
	return "localhost:8765"
}
