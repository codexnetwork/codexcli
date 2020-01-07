package main

import (
	// Load all contracts here, so we can always read and decode

	// transactions with those contracts.

	"github.com/codexnetwork/codexcli/codexcli/cmd"
	"github.com/codexnetwork/codexio-go/ecc"
	_ "github.com/codexnetwork/codexio-go/msig"
	_ "github.com/codexnetwork/codexio-go/system"
	_ "github.com/codexnetwork/codexio-go/token"
)

var version = "dev"

func init() {
	cmd.Version = version
	ecc.PublicKeyPrefixCompat = "FOSC"
}

func main() {
	cmd.Execute()
}
