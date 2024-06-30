package gclient

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

var Client gnoclient.Client

func init() {
	// Initialize keybase from a directory
	keybase, _ := keys.NewKeyBaseFromDir(".")

	// Create signer
	signer := gnoclient.SignerFromKeybase{
		Keybase:  keybase,
		Account:  "<keypair_name>",     // Name of your keypair in keybase
		Password: "<keypair_password>", // Password to decrypt your keypair
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

}

func GetClient() *gnoclient.Client {
	return &Client
}
