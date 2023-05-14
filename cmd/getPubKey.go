package cmd

import (
	"fmt"

	"github.com/asa93/ecdh-cli-go/ecdh"
	"github.com/spf13/cobra"
)

// exportPubKeyCmd represents the exportPubKey command
var getPubKey = &cobra.Command{
	Use:   "getPubKey",
	Short: "Generate your uncompressed public key",
	Long:  `This key has to be shared with the other party so they can encrypt/decrypt file using it.`,
	Run: func(cmd *cobra.Command, args []string) {

		privKey, _ := ecdh.GetKeys()

		fmt.Println("🔑 your uncompressed pubkey is: \n ", privKey.PubKey())

	},
}

func init() {
	rootCmd.AddCommand(getPubKey)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportPubKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportPubKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
