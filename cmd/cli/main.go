package main

import (
	"fmt"
	"log"
	"os"

	avltree "github.com/jpfcabral/search-suggestion-api/internal/trees"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "autocom",
		Short: "An autocomplete system",
		Run: func(cmd *cobra.Command, args []string) {
			tree := &avltree.AVLTree{}
			tree.PopulateFromCorpus("assets/br-utf8.txt")

			if len(args) < 1 {
				fmt.Println("Provide a prefix")
				return
			} else if len(args) > 1 {
				fmt.Println("Provide just one prefix")
				return
			}

			prefix := args[0]

			fmt.Println(tree.WordsWithPrefix(prefix))

		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
