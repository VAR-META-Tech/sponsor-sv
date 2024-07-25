package gclient

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

var Client gnoclient.Client
var CallerClient gnoclient.Client

func init() {
	// Initialize keybase from a directory
	// keybase, _ := keys.NewKeyBaseFromDir("/home/vm/thinhnx/sponsor-sv")
	 keybase, _ := keys.NewKeyBaseFromDir("/Users/thinhnx/Library/Application Support/gno")
	// Create signer
	signer := gnoclient.SignerFromKeybase{
		Keybase:  keybase,
		Account:  "myAdenaKey",     // Name of your keypair in keybase
		Password: "1", // Password to decrypt your keypair
		ChainID:  "dev",                // id of Gno.land chain
	}

	// Create signer
	callerSigner := gnoclient.SignerFromKeybase{
		Keybase:  keybase,
		Account:  "testKey",     // Name of your keypair in keybase
		Password: "1", // Password to decrypt your keypair
		ChainID:  "dev",                // id of Gno.land chain
	}

	// Initialize the RPC client
	rpc, err := rpcclient.NewHTTPClient("0.0.0.0:26657")
	if err != nil {
		panic(err)
	}

	// Initialize the gnoclient
	Client = gnoclient.Client{
		Signer:    signer,
		RPCClient: rpc,
	}
	CallerClient = gnoclient.Client{
		Signer: callerSigner,
		RPCClient: rpc,
	}

}

func GetClient() *gnoclient.Client {
	return &Client
}
func GetCallerClient() *gnoclient.Client {
	return &CallerClient
}
