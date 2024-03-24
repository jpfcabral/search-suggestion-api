package main

import (
	avltree "github.com/jpfcabral/search-suggestion-api/internal/trees"
)

func main() {
	tree := new(avltree.AVLTree)
	tree.PopulateFromCorpus("assets/br-utf8.txt")
	tree.DisplayInOrder()
}
