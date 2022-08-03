package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixNoun_Match(t *testing.T) {
	assert.True(t, NounSuffix1.Match("kitapçılar"))
	assert.True(t, NounSuffix1.Match("bebekler"))
	assert.True(t, NounSuffix2.Match("bulurum"))
	assert.False(t, NounSuffix1.Match("kitapçı"))
}

func TestSuffixNoun_GetOptionalLetter(t *testing.T) {
	assert.Equal(t, 'ı', *NounSuffix3.GetOptionalLetter("kitabımızı"))
	assert.Equal(t, 'y', *NounSuffix17.GetOptionalLetter("buluruy"))

	assert.Nil(t, NounSuffix17.GetOptionalLetter("anlamsız"))
	assert.Nil(t, NounSuffix1.GetOptionalLetter("bulurum"))
}

func TestSuffixNoun_RemoveSuffix(t *testing.T) {
	assert.Equal(t, "bebek", NounSuffix1.RemoveSuffix("bebekler"))
	assert.Equal(t, "buluru", NounSuffix2.RemoveSuffix("bulurum"))
}

func TestSuffixNoun_String(t *testing.T) {
	assert.Equal(t, "-lArI", NounSuffix7.String())
}

func TestSuffixNominalVerb_Match(t *testing.T) {
	assert.True(t, NominalVerbSuffix4.Match("satıyorsunuz"))
}

func TestSuffixNominalVerb_GetOptionalLetter(t *testing.T) {
	assert.Equal(t, 'y', *NominalVerbSuffix1.GetOptionalLetter("satıy"))

	assert.Nil(t, NominalVerbSuffix1.GetOptionalLetter("satıyor"))
	assert.Nil(t, NominalVerbSuffix6.GetOptionalLetter("satıyor"))
}

func TestSuffixNominalVerb_RemoveSuffix(t *testing.T) {
	assert.Equal(t, "satıyor", NominalVerbSuffix4.RemoveSuffix("satıyorsunuz"))
	assert.Equal(t, "satıyorsunuzz", NominalVerbSuffix4.RemoveSuffix("satıyorsunuzz"))
}

func TestSuffixNominalVerb_String(t *testing.T) {
	assert.Equal(t, "-n", NominalVerbSuffix7.String())
}

func TestSuffixDerivational_Match(t *testing.T) {
	assert.True(t, DerivationalSuffix1.Match("gozlu"))
}

func TestSuffixDerivational_GetOptionalLetter(t *testing.T) {
	assert.Nil(t, DerivationalSuffix1.GetOptionalLetter("gozlu"))
}

func TestSuffixDerivational_RemoveSuffix(t *testing.T) {
	assert.Equal(t, "goz", DerivationalSuffix1.RemoveSuffix("gozlu"))
}

func TestSuffixDerivational_String(t *testing.T) {
	assert.Equal(t, "-lU", DerivationalSuffix1.String())
}
