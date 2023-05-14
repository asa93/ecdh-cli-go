package cmd

import (
	"fmt"
	"os"

	"github.com/asa93/ecdh-cli/ecdh"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file using ECDH",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := cmd.Flags().GetString("path")

		if path == "" {
			fmt.Println("No path provided")
			os.Exit(1)
		}

		privKey, pubKey2 := ecdh.GetKeys()

		fmt.Println("pubkey 1 --", privKey.PubKey())
		fmt.Println("pubkey 2 --", pubKey2)

		ecdh.Decrypt(privKey, pubKey2, path)

	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.PersistentFlags().String("path", "", "path to file to decrypt")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
