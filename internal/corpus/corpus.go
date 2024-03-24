package corpus

import (
	"bufio"
	"os"
	"strings"
)

type Corpus struct {
	Words []string
}

// NewCorpusFromFile creates a new Corpus instance by reading words from a text file.
func NewCorpusFromFile(filename string) (*Corpus, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := make([]string, 0)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text()) // Convert to lowercase
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Corpus{Words: words}, nil
}
