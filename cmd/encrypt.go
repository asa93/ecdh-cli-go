package cmd

import (
	"fmt"
	"os"

	"github.com/asa93/ecdh-cli/ecdh"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file using ECDH",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := cmd.Flags().GetString("path")

		if path == "" {
			fmt.Println("No path provided")
			os.Exit(1)
		}

		privKey, pubKey2 := ecdh.GetKeys()

		fmt.Println("🔑 your pubkey       --", privKey.PubKey())
		fmt.Println("🔑 recipient pubkey  --", pubKey2)

		ecdh.Encrypt(privKey, pubKey2, path)

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.PersistentFlags().String("path", "", "path to file to encrypt")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
