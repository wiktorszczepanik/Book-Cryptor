package cmd

import (
	"book-cryptor/inter/encrypt"
	"book-cryptor/inter/file"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
		inputFile, keyFile, err = file.GetEssensialFiles(inputFilePath, keyFilePath)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		var cipher string
		switch mode {
		case "beale":
			cipher, err = encrypt.Beale(inputFile, keyFile, separator, exact)
		case "ottendorf":
			// cipher, err = encrypt.EncryptOttendorf(inputFile, keyFile)
			cipher, err = encrypt.Beale(inputFile, keyFile, separator, exact)
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
		if outputFilePath != "" {
			err = file.SaveOutput(outputFilePath, cipher)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot save cipher to output file.")
				os.Exit(1)
			}
		} else {
			fmt.Println(cipher)
		}
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&inputFilePath, "in", "i", "", "input file for encryption")
	encryptCmd.Flags().StringVarP(&keyFilePath, "key", "k", "", "key file for encryption (book/text)")
	encryptCmd.Flags().StringVarP(&mode, "mode", "m", "", "encryption mode implementation")
	encryptCmd.Flags().StringVarP(&separator, "separator", "s", ", ", "separator in file for encryption")
	encryptCmd.Flags().StringVarP(&outputFilePath, "out", "o", "", "encryption output file")
	encryptCmd.Flags().BoolVarP(&exact, "exact", "e", false, "include exact set of characters in input file")
	rootCmd.AddCommand(encryptCmd)
}
