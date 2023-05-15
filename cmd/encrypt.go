package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/asa93/ecdh-cli/ecdh"
	"github.com/spf13/cobra"
)

type PublicKeyJson struct {
	keys []string `json:"keys"`
}

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

		batch, _ := cmd.Flags().GetString("batch")

		privKey, pubKey := ecdh.GetKeys()

		fmt.Println("🔑 your pubkey       --", privKey.PubKey())

		if batch == "" {

			fmt.Println("🔑 recipient pubkey  --", pubKey)

			ecdh.Encrypt(privKey, pubKey, path, "src/encrypted-"+pubKey.GetX().String()[0:12])
		} else {
			file, err := ioutil.ReadFile(batch)
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			publicKeys := strings.Split(string(file), "\n")

			for i := 0; i < len(publicKeys); i++ {
				pubKey := ecdh.ParsePublicKey(publicKeys[0])

				ecdh.Encrypt(privKey, pubKey, path, "src/encrypted-"+pubKey.GetX().String()[0:12])
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.PersistentFlags().String("path", "", "path to file to encrypt")

	encryptCmd.PersistentFlags().String("batch", "", "path to file with batch of public keys")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
