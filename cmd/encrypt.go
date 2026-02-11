package cmd

import (
	"book-cryptor/inter/encrypt"
	"book-cryptor/inter/file"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFilePath string
var keyFilePath string
var mode string
var outputFilePath string

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
		var inputFile, keyFile *os.File
		var err error
		if inputFile, err = file.GetFile(inputFilePath); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if keyFile, err = file.GetFile(keyFilePath); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		var cipher string
		switch mode {
		case "beale":
			cipher, err = encrypt.EncryptBeale(inputFile, keyFile)
		case "ottendorf":
			// cipher, err = encrypt.EncryptOttendorf(inputFile, keyFile)
			cipher, err = encrypt.EncryptBeale(inputFile, keyFile)
		default:
			fmt.Fprintf(os.Stderr, "Incorrect encryption mode: %s", mode)
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		defer inputFile.Close()
		defer keyFile.Close()
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
