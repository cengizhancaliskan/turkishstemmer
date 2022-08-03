package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	stemmer := New()

	assert.NotNil(t, stemmer)
}

func TestStemmer_Stem(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{actual: "eriklimişsincesine", expected: "erik"},
		{actual: "ayfon", expected: "ayfon"},
		{actual: "adrese", expected: "adre"},
		{actual: "abiyeler", expected: "abiye"},
		{actual: "eriklimişsincesine", expected: "erik"},
		{actual: "guzelim", expected: "guzel"},
		{actual: "satıyorsunuz", expected: "satıyor"},
		{actual: "taksicisiniz", expected: "taksiç"},
		{actual: "türkiyedir", expected: "türki"},
		{actual: "telefonları", expected: "telefon"},
		{actual: "acana", expected: "acan"},
		{actual: "tekken", expected: "tekken"},
		{actual: "telefonu", expected: "telefon"},
		{actual: "telefon", expected: "telefon"},
		{actual: "alarm", expected: "alarm"},
		{actual: "alarmı", expected: "alarm"},
		{actual: "adını", expected: "adın"},
		{actual: "adın", expected: "adın"},
		{actual: "altın", expected: "altın"},
		{actual: "aparatı", expected: "aparat"},
		{actual: "arada", expected: "ara"},
		{actual: "arasındaki", expected: "ara"},
		{actual: "arasındakı", expected: "arasındak"},
		{actual: "gozluklerinde", expected: "gozluk"},
		{actual: "monitoru", expected: "monitor"},
		{actual: "monitörü", expected: "monitör"},
		{actual: "monitorü", expected: "monitor"},
		{actual: "monitöru", expected: "monitör"},
		{actual: "çantası", expected: "çanta"},
		{actual: "çantasıı", expected: "çantası"},
		{actual: "çantasi", expected: "çanta"},
		{actual: "ağrılı", expected: "ağrı"},
	}
	stemmer := New()

	for _, tt := range tests {
		t.Run("TestStem", func(t *testing.T) {
			assert.Equal(t, tt.expected, stemmer.Stem(tt.actual))
		})
	}
}

func TestStemmer_stemWord(t *testing.T) {
	stemmer := New()

	assert.Equal(t, "arasında", stemmer.stemWord("arasındaki", NounSuffix18))
	assert.Equal(t, "arasındaki", stemmer.stemWord("arasındaki", NounSuffix6))
	assert.Equal(t, "arasındaki", stemmer.stemWord("arasındaki", NounSuffix8))
	assert.Equal(t, "arası", stemmer.stemWord("arasında", NounSuffix14))
	assert.Equal(t, "arasın", stemmer.stemWord("arasında", NounSuffix13))
	assert.Equal(t, "ara", stemmer.stemWord("arası", NounSuffix6))
	assert.Equal(t, "aras", stemmer.stemWord("arasın", NounSuffix4))
}

func TestStemmer_validateWord(t *testing.T) {
	stemmer := New()

	assert.True(t, stemmer.validateWord("saatler"))
	assert.True(t, stemmer.validateWord("latte"))
	assert.False(t, stemmer.validateWord("qué pasa"))
	assert.False(t, stemmer.validateWord("tekken"))
	assert.False(t, stemmer.validateWord("ağda"))
	assert.False(t, stemmer.validateWord("a"))
	assert.False(t, stemmer.validateWord(""))
}
