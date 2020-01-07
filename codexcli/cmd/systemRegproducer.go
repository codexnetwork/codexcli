package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/codexnetwork/codexio-go/ecc"
	"github.com/codexnetwork/codexio-go/system"
)

var systemRegisterProducerCmd = &cobra.Command{
	Use:   "regproducer [account_name] [public_key] [website_url]",
	Short: "Register an account as a block producer candidate.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		accountName := toAccount(args[0], "account name")
		publicKey, err := ecc.NewPublicKey(args[1])
		errorCheck(fmt.Sprintf("%q invalid public key", args[1]), err)
		websiteURL := args[2]

		pushEOSCActions(api,
			system.NewRegProducer(accountName, publicKey, websiteURL, 1000),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemRegisterProducerCmd)

	systemRegisterProducerCmd.Flags().IntP("location", "", 0, "Location number (reserved)")
}
