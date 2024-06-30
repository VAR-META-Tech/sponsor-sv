package account

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/std"
)

func GetAccountBaseWithAddr(client *gnoclient.Client, addrQuery string) (*std.BaseAccount, error) {
	// Convert Gno address string to `crypto.Address`
	//g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5
	addr, err := crypto.AddressFromBech32(addrQuery) // your Gno address
	if err != nil {
		return nil, err
	}

	accountRes, _, err := client.QueryAccount(addr)
	if err != nil {
		return nil, err
	}
	return accountRes, nil

}
