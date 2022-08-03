package turkishstemmer

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadWordsFromSliceBytes Reads data from slide of bytes and split as newline ('\n')
func LoadWordsFromSliceBytes(data []byte) []string {
	var words []string
	for _, val := range bytes.Split(data, []byte{'\n'}) {
		words = append(words, string(val))
	}

	return words
}

// LoadWords Reads file line by line and returns string Slice
func LoadWords(path string) []string {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to open file %s", path))
	}

	var words []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	return words
}
