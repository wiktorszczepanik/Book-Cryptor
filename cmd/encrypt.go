/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"book-cryptor/internal"
	"crypto/cipher"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFilePath string
var keyFilePath string
var mode string
var outputFilePath string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Command for enrypting input file",
	Long: `Command for enrypting input file by using key in book/text format and specified cryptography mode.
For example:

./book-cryptor encrypt --in=file1.txt --key=book.txt --mode=beale --out=file2.txt
./book-cryptor encrypt --in=file1.txt --key=book.pdf --mode=ottendorf --out=file2.txt
./book-cryptor encrypt --in=file1.txt --key=book.epub --mode=beale --out=file2.txt

Supported key formats are .txt .pdf .epub
Supported mode techniques are "beale" "Ottendorf"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := internal.GetFile(inputFilePath)
		keyFile := internal.GetFile(keyFilePath)
		var cipher string
		var err error
		switch mode {
		case "beale":
			cipher, err := internal.EncryptBeale(inputFile, keyFile)
		case "ottendorf":
			//...
		default:
			fmt.Fprint(os.Stderr, "Incorrect encryption mode: %s", mode)
			os.Exit(1)
		}
		defer inputFile.Close()
		defer keyFile.Close()
		fmt.Println("encrypt called")
		fmt.Println(cipher)
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&inputFilePath, "in", "i", "", "Input file for encryption")
	encryptCmd.Flags().StringVarP(&keyFilePath, "key", "k", "", "Key file for encryption (book/text)")
	encryptCmd.Flags().StringVarP(&mode, "mode", "m", "", "Encryption mode implementation")
	encryptCmd.Flags().StringVarP(&outputFilePath, "out", "o", "", "Encryption output file")
	rootCmd.AddCommand(encryptCmd)
}
