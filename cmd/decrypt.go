package cmd

import (
	"book-cryptor/inter/decrypt"
	"book-cryptor/inter/file"
	"book-cryptor/inter/oper"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Command for decrypting input file",
	Long: `Command for decrypting input file by using key in book/text format and specifed cryptograpy mode.
For example:

./book-cryptor decrypt --in=file1.txt --key=book.txt --mode=beale --out=file2.txt
./book-cryptor decrypt --in=file1.txt --key=book.pdf --mode=ottendorf --out=file2.txt
./book-cryptor decrypt --in=file1.txt --key=book.epub --mode=beale --out=file2.txt

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
		var plain string
		switch mode {
		case "beale":
			plain, err = decrypt.DecryptBeale(inputFile, keyFile, separator)
		case "ottendorf":
			// cipher, err = decrypt.DecryptOttendorf(inputFile, keyFile)
			plain, err = decrypt.DecryptBeale(inputFile, keyFile, separator)
		default:
			fmt.Fprintf(os.Stderr, "Incorrect decryption mode: %s", mode)
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		defer inputFile.Close()
		defer keyFile.Close()
		if outputFilePath != "" {
			err = oper.SaveOutput(outputFilePath, plain)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot save plain text to output file.")
				os.Exit(1)
			}
		} else {
			fmt.Println(plain)
		}
	},
}

func init() {
	decryptCmd.Flags().StringVarP(&inputFilePath, "in", "i", "", "input file for decryption")
	decryptCmd.Flags().StringVarP(&keyFilePath, "key", "k", "", "key file for decryption (book/text)")
	decryptCmd.Flags().StringVarP(&mode, "mode", "m", "", "decryption mode implementation")
	decryptCmd.Flags().StringVarP(&separator, "separator", "s", ", ", "separator in file for decryption")
	decryptCmd.Flags().StringVarP(&outputFilePath, "out", "o", "", "decryption output file")
	rootCmd.AddCommand(decryptCmd)
}
