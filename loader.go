package turkishstemmer

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

// loadWordsFromSliceBytes Reads data from slide of bytes and split as newline ('\n')
func loadWordsFromSliceBytes(data []byte) []string {
	var words []string
	for _, val := range bytes.Split(data, []byte{'\n'}) {
		words = append(words, string(val))
	}

	return words
}

// loadWords Reads file line by line and returns string Slice
func loadWords(path string) []string {
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicf("failed to close file %s.", path)
		}
	}(file)

	if err != nil {
		log.Panicf("failed to open file %s.", path)
	}

	var words []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	return words
}
