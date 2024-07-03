package transfer

import (
	"log"
	"sponsor-sv/services/gclient"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/std"
)

func SponsorMsg() std.Tx {
	sponsorTxCfg := gnoclient.SponsorTxCfg{
        BaseTxCfg: gnoclient.BaseTxCfg{
            GasFee:         "1000000ugnot",  // gas price
            GasWanted:      1000000,         // gas limit
            Memo:           "",              // transaction memo
        },
        SponsorAddress: crypto.MustAddressFromString("g162jgpk4740r6a7g53cgz9ahxqtyuekgqchw6w9"),
    }
	callerCli := gclient.GetCallerClient()
    tx, err := callerCli.NewSponsorTransaction(sponsorTxCfg, gnoclient.MsgCall{
        PkgPath:  "gno.land/r/demo/deep/very/deep",
        FuncName: "Render",
        Args:     []string{""},
    })
    if err != nil {
        log.Fatalf("Failed to Create SponsorTransaction : %v\n", err)
    }

    sponsorTx, err := callerCli.SignTransaction(*tx,0, 0)
    if err != nil {
        log.Fatalf("Failed to Create SponsorTransaction : %v\n", err)
    }
	return *sponsorTx
}