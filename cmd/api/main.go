package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	avltree "github.com/jpfcabral/search-suggestion-api/internal/trees"
)

func main() {
	r := gin.Default()
	tree := &avltree.AVLTree{}
	tree.PopulateFromCorpus("assets/br-utf8.txt")

	r.GET("/search", func(c *gin.Context) {
		prefix := c.Query("prefix")
		if prefix == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Prefix parameter is required"})
			return
		}

		words := tree.WordsWithPrefix(prefix)
		c.JSON(http.StatusOK, gin.H{"words": words})
	})

	// Run the server
	r.Run(":8080")
}
