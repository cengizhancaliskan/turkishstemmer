package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadWordsFromSliceBytes(t *testing.T) {
	words := loadWordsFromSliceBytes([]byte("hello\nfrom\ntest"))

	assert.Len(t, words, 3)
}

func TestLoadWords(t *testing.T) {
	words := loadWords("data/last_consonant_exceptions.txt")
	assert.Len(t, words, 4)
}
