package cmd

import (
	"fmt"
	"os"

	"github.com/netscale-technologies/netcomp-cli/cipher"

	"github.com/spf13/cobra"
)

// Encryption command
var encryptCmd = &cobra.Command{
	Use:   "encrypt [string to encrypt]",
	Short: "Encrypt value using AES 256 CBC algorithm",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get key value from flags
		key, err := cmd.Flags().GetString("key")
		cobra.CheckErr(err)

		// Get text from args
		text := args[0]

		// Key must be provided
		if key == "" {
			fmt.Fprintln(os.Stderr, "Error: invalid key or iv")
			os.Exit(-1)
		}

		result, err := cipher.EncryptAES(key, text)
		cobra.CheckErr(err)

		fmt.Println(result)
	},
}

// Decryption command
var decryptCmd = &cobra.Command{
	Use:   "decrypt [string to decrypt]",
	Short: "Decrypt value using AES 256 CBC algorithm",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get key value from flags
		key, err := cmd.Flags().GetString("key")
		cobra.CheckErr(err)

		// Get text from args
		text := args[0]

		// Key must be provided
		if key == "" {
			fmt.Fprintln(os.Stderr, "Error: invalid key or iv")
			os.Exit(-1)
		}

		result, err := cipher.DecryptAES(key, text)
		cobra.CheckErr(err)

		fmt.Println(result)
	},
}

func init() {
	encryptCmd.Flags().StringP("key", "k", "", "256 bits secret key")
	decryptCmd.Flags().StringP("key", "k", "", "256 bits secret key")
}
