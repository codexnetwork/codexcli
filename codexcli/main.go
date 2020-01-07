package main

import (
	// Load all contracts here, so we can always read and decode
	// transactions with those contracts.
	_ "github.com/codexnetwork/codexio-go/msig"
	_ "github.com/codexnetwork/codexio-go/system"
	_ "github.com/codexnetwork/codexio-go/token"

	"github.com/codexnetwork/codexcli/codexcli/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
