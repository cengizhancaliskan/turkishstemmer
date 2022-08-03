package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	words := []string{"lungo", "espresso", "long black"}
	assert.True(t, Contains(words, "lungo"))
	assert.False(t, Contains(words, "ristretto"))
}

func TestReplaceStringAtIndex(t *testing.T) {
	assert.Equal(t, "kaşaği", ReplaceStringAtIndex("kaşağı", 'i', 5))
	assert.Equal(t, "Hallo", ReplaceStringAtIndex("Hello", 'a', 1))
}

func TestGetVowels(t *testing.T) {
	assert.Equal(t, []rune("aeıioöuü"), GetVowels("aeıioöuü"))
	assert.Equal(t, []rune("üu"), GetVowels("ükulş"))
	assert.Nil(t, GetVowels("bcç"))
	assert.Nil(t, GetVowels(""))
}

func TestCountSyllables(t *testing.T) {
	assert.Equal(t, 2, CountSyllables("okul"))
	assert.Equal(t, 1, CountSyllables("test"))
	assert.Equal(t, 0, CountSyllables("bçdk"))
	assert.Equal(t, 0, CountSyllables(""))
}

func TestValidateOptionalLetter(t *testing.T) {
	assert.True(t, ValidateOptionalLetter("türkiy", &([]rune("y"))[0]))
	assert.True(t, ValidateOptionalLetter("kebap", &([]rune("p"))[0]))
	assert.True(t, ValidateOptionalLetter("keba", &([]rune("a"))[0]))
	assert.False(t, ValidateOptionalLetter("kebp", &([]rune("p"))[0]))
	assert.False(t, ValidateOptionalLetter("kea", &([]rune("a"))[0]))
	assert.False(t, ValidateOptionalLetter("at", &([]rune("a"))[0]))
	assert.False(t, ValidateOptionalLetter("c", nil))
}

func TestVowelHarmony(t *testing.T) {
	assert.True(t, VowelHarmony("e", "i"))
	assert.False(t, VowelHarmony("a", "i"))
}

func TestIsTurkishWord(t *testing.T) {
	assert.True(t, IsTurkishWord("kebapçi"))
	assert.False(t, IsTurkishWord("Straße"))
}

func TestHasFrontness(t *testing.T) {
	assert.True(t, HasFrontness("e", "i"))
	assert.False(t, HasFrontness("a", "i"))
}

func TestHasRoundness(t *testing.T) {
	assert.True(t, HasRoundness("o", "u"))
	assert.False(t, HasRoundness("o", "i"))
}

func TestHasVowelHarmony(t *testing.T) {
	assert.True(t, HasVowelHarmony("okul"))
	assert.True(t, HasVowelHarmony("k"))
	assert.False(t, HasVowelHarmony("okuller"))
}

func TestGetType(t *testing.T) {
	assert.Equal(t, "NounState", GetType(NounStateA))
	assert.Equal(t, "Suffix", GetType(NounSuffix17))
	assert.Equal(t, "Suffix", GetType(NounSuffix1))
}
