package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// var verbose bool
var inputFilePath string
var keyFilePath string
var mode string
var separator string
var outputFilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "book-cryptor",
	Short: "A tool for encrypting and decrypting book ciphers",
	Long: `Book-Cryptor is a versatile CLI tool for encrypting and decrypting book-based ciphers. 

It allows users to encrypt and decrypt text using classic book cipher techniques, 
where a specific book or text is used as the key. The tool supports various methods 
of encoding, such as location-based ciphers (e.g., page, line, word), and offers 
an easy-to-use interface for both beginners and cryptography enthusiasts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.book-cryptor.yaml)")
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose operations output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
