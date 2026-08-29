package turkishstemmer

import "testing"

const benchmarkWordEriklimissincesine = "eriklimişsincesine"

var (
	benchmarkWords = []string{
		benchmarkWordEriklimissincesine,
		"satıyorsunuz",
		"taksicisiniz",
		"telefonları",
		"arasındaki",
		"gozluklerinde",
		"monitörü",
		// Exercise the typo-correction fallback on a representative misspelling.
		"çantasıı",
		"ağrılı",
		"kalelerimizdekilerden",
		"çocuğuymuşumcasına",
	}
	benchmarkStemResult    string
	benchmarkStemmerResult Stemmer
)

func BenchmarkStem(b *testing.B) {
	b.ReportAllocs()
	stemmer := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range benchmarkWords {
			benchmarkStemResult = stemmer.Stem(word)
		}
	}
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkStemmerResult = New()
	}
}
