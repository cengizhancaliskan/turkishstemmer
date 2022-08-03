package turkishstemmer

import (
	"regexp"
)

type Suffix struct {
	Name                  string
	Pattern               *regexp.Regexp
	OptionalLetterPattern *regexp.Regexp
	OptionalLetterCheck   bool
	CheckHarmony          bool
	OptionalLetter        string
}

var (
	DerivationalSuffix1 = NewSuffix("-lU", "lı|li|lu|lü", "", true)

	NominalVerbSuffix1  = NewSuffix("-(y)Um", "ım|im|um|üm", "y", true)
	NominalVerbSuffix2  = NewSuffix("-sUn", "sın|sin|sun|sün", "", true)
	NominalVerbSuffix3  = NewSuffix("-(y)Uz", "ız|iz|uz|üz", "y", true)
	NominalVerbSuffix4  = NewSuffix("-sUnUz", "sınız|siniz|sunuz|sünüz", "", true)
	NominalVerbSuffix5  = NewSuffix("-lAr", "lar|ler", "", true)
	NominalVerbSuffix6  = NewSuffix("-m", "m", "", true)
	NominalVerbSuffix7  = NewSuffix("-n", "n", "", true)
	NominalVerbSuffix8  = NewSuffix("-k", "k", "", true)
	NominalVerbSuffix9  = NewSuffix("-nUz", "nız|niz|nuz|nüz", "", true)
	NominalVerbSuffix10 = NewSuffix("-DUr", "tır|tir|tur|tür|dır|dir|dur|dür", "", true)
	NominalVerbSuffix11 = NewSuffix("-cAsInA", "casına|çasına|cesine|çesine", "", true)
	NominalVerbSuffix12 = NewSuffix("-(y)DU", "dı|di|du|dü|tı|ti|tu|tü", "y", true)
	NominalVerbSuffix13 = NewSuffix("-(y)sA", "sa|se", "y", true)
	NominalVerbSuffix14 = NewSuffix("-(y)mUş", "muş|miş|müş|mış", "y", true)
	NominalVerbSuffix15 = NewSuffix("-(y)ken", "ken", "y", true)

	NounSuffix1  = NewSuffix("-lAr", "lar|ler", "", true)
	NounSuffix2  = NewSuffix("-(U)m", "m", "ı|i|u|ü", true)
	NounSuffix3  = NewSuffix("-(U)mUz", "mız|miz|muz|müz", "ı|i|u|ü", true)
	NounSuffix4  = NewSuffix("-Un", "ın|in|un|ün", "", true)
	NounSuffix5  = NewSuffix("-(U)nUz", "nız|niz|nuz|nüz", "ı|i|u|ü", true)
	NounSuffix6  = NewSuffix("-(s)U", "ı|i|u|ü", "s", true)
	NounSuffix7  = NewSuffix("-lArI", "ları|leri", "", true)
	NounSuffix8  = NewSuffix("-(y)U", "ı|i|u|ü", "y", true)
	NounSuffix9  = NewSuffix("-nU", "nı|ni|nu|nü", "", true)
	NounSuffix10 = NewSuffix("-(n)Un", "ın|in|un|ün", "n", true)
	NounSuffix11 = NewSuffix("-(y)A", "a|e", "y", true)
	NounSuffix12 = NewSuffix("-nA", "na|ne", "", true)
	NounSuffix13 = NewSuffix("-DA", "da|de|ta|te", "", true)
	NounSuffix14 = NewSuffix("-nDA", "nta|nte|nda|nde", "", true)
	NounSuffix15 = NewSuffix("-DAn", "dan|tan|den|ten", "", true)
	NounSuffix16 = NewSuffix("-nDAn", "ndan|ntan|nden|nten", "", true)
	NounSuffix17 = NewSuffix("-(y)lA", "la|le", "y", true)
	NounSuffix18 = NewSuffix("-ki", "ki", "", false)
	NounSuffix19 = NewSuffix("-(n)cA", "ca|ce", "n", true)

	// The order of this slice definition determines the priority of the suffix.
	DerivationalSuffixValues = []Suffix{DerivationalSuffix1}
	NominalVerbSuffixValues  = []Suffix{NominalVerbSuffix11, NominalVerbSuffix4, NominalVerbSuffix14, NominalVerbSuffix15, NominalVerbSuffix2, NominalVerbSuffix5, NominalVerbSuffix9, NominalVerbSuffix10, NominalVerbSuffix3, NominalVerbSuffix1, NominalVerbSuffix12, NominalVerbSuffix13, NominalVerbSuffix6, NominalVerbSuffix7, NominalVerbSuffix8}
	NounSuffixValues         = []Suffix{NounSuffix16, NounSuffix7, NounSuffix3, NounSuffix5, NounSuffix1, NounSuffix14, NounSuffix15, NounSuffix17, NounSuffix10, NounSuffix19, NounSuffix4, NounSuffix9, NounSuffix12, NounSuffix13, NounSuffix18, NounSuffix2, NounSuffix6, NounSuffix8, NounSuffix11}
)

func NewSuffix(name, pattern, optionalLetter string, checkHarmony bool) Suffix {
	s := Suffix{
		Name:    name,
		Pattern: regexp.MustCompile("(" + pattern + ")$"),
	}

	if len(optionalLetter) > 0 {
		s.OptionalLetterCheck = true
		s.OptionalLetter = optionalLetter
		s.OptionalLetterPattern = regexp.MustCompile("(" + optionalLetter + ")$")
	}
	s.CheckHarmony = checkHarmony

	return s
}

func (s Suffix) Match(word string) bool {
	return s.Pattern.Match([]byte(word))
}

func (s Suffix) GetOptionalLetter(word string) *rune {
	if s.OptionalLetterCheck {
		res := s.OptionalLetterPattern.FindString(word)
		if res != "" {
			r := []rune(res)

			return &r[0]
		}
	}

	return nil
}

func (s Suffix) RemoveSuffix(word string) string {
	return s.Pattern.ReplaceAllString(word, "")
}

func (s Suffix) String() string {
	return s.Name
}
