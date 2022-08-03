package turkishstemmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadWordsFromSliceBytes(t *testing.T) {
	words := LoadWordsFromSliceBytes([]byte("hello\nfrom\ntest"))

	assert.Len(t, words, 3)
}

func TestLoadWords(t *testing.T) {
	words := LoadWords("data/last_consonant_exceptions.txt")
	assert.Len(t, words, 4)
}
