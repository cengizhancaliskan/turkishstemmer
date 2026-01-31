package turkishstemmer

import (
	"math"
	"reflect"
	"sort"
	"strings"
)

type Stems []string

type Stemmer struct {
	ProtectedWords            []string
	VowelHarmonyExceptions    []string
	LastConsonantExceptions   []string
	AverageStemSizeExceptions []string
}

// New constructs a new Stemmer.
func New() Stemmer {
	protectedWords := LoadWordsFromSliceBytes(DefaultProtectedWordsFile)
	vowelHarmonyExceptions := LoadWordsFromSliceBytes(DefaultVowelHarmonyExceptionsFile)
	lastConsonantExceptions := LoadWordsFromSliceBytes(DefaultLastConsonantExceptionsFile)
	averageStemSizeExceptions := LoadWordsFromSliceBytes(DefaultAverageStemSizeExceptionsFile)

	s := Stemmer{
		ProtectedWords:            protectedWords,
		VowelHarmonyExceptions:    vowelHarmonyExceptions,
		LastConsonantExceptions:   lastConsonantExceptions,
		AverageStemSizeExceptions: averageStemSizeExceptions,
	}

	return s
}

// Stem returns the stemmed word of a given un-stemmed word.
// In case it remained unstemmed it attempts to correct some mistypes such
// as 'u' instead of 'ü' and 'i' instead of 'ı'.
func (s Stemmer) Stem(word string, tryCount ...int) string {
	if !s.validateWord(word) {
		return word
	}

	var stems Stems

	// Process the word with the nominal verb suffix State machine.
	s.genericSuffixStripper(GetInitialNominalVerbState(), word, &stems)
	var wordsToStem = stems

	wordsToStem = append(wordsToStem, word)

	// Process each possible stem with the noun suffix State machine.
	for _, w := range wordsToStem {
		s.genericSuffixStripper(NounStateA, w, &stems)
	}

	wordsToStem = stems
	wordsToStem = append(wordsToStem, word)

	// If none of the stemming rules matches, then replace with vice versa and try stemming again
	if Contains(wordsToStem, word) && len(wordsToStem) < 2 && len(tryCount) < 1 {
		wordChars := []rune(word)
		lastLetterIndex := len(wordChars) - 1
		lastLetter := string(wordChars[lastLetterIndex])
		var wordChanged bool
		switch lastLetter {
		case "u":
			word = ReplaceStringAtIndex(word, 'ü', lastLetterIndex)
			wordChanged = true
		case "ü":
			word = ReplaceStringAtIndex(word, 'u', lastLetterIndex)
			wordChanged = true
		case "ı":
			word = ReplaceStringAtIndex(word, 'i', lastLetterIndex)
			wordChanged = true
		case "i":
			word = ReplaceStringAtIndex(word, 'ı', lastLetterIndex)
			wordChanged = true
		}

		if wordChanged {
			return s.Stem(word, 1)
		}
	}

	for _, w := range wordsToStem {
		s.genericSuffixStripper(DerivationalStateA, w, &stems)
	}

	return s.postProcess(stems, word)
}

func (s Stemmer) genericSuffixStripper(state State, word string, stems *Stems) {
	var transitions Transitions
	state.AddTransitions(word, &transitions, state)

	for len(transitions) > 0 {
		transition := transitions[0]
		// remove current index
		transitions = append(transitions[:0], transitions[0+1:]...)

		stem := s.stemWord(transition.Word, transition.Suffix)

		if stem != transition.Word {
			if transition.NextState.FinalState() {
				for i := 0; i < len(transitions); i++ {
					if transitions[i].Marked ||
						(reflect.DeepEqual(transitions[i].StartState, transition.StartState) &&
							reflect.DeepEqual(transitions[i].NextState, transition.NextState)) {
						copy(transitions[i:], transitions[i+1:])
						transitions[len(transitions)-1] = nil
						transitions = transitions[:len(transitions)-1]
						i--
					}
				}

				*stems = append(*stems, stem)

				transition.NextState.AddTransitions(stem, &transitions, transition.NextState)
			} else {
				// TODO: Possible bug
				for _, similarTransition := range *transition.SimilarTransitions(transitions) {
					similarTransition.Marked = true
				}
				transition.NextState.AddTransitions(stem, &transitions, transition.NextState)
			}
		}
	}
}

func (s Stemmer) stemWord(word string, suffix Suffix) string {
	stemmedWord := word
	if s.shouldBeMarked(word, suffix) && suffix.Match(word) {
		stemmedWord = suffix.RemoveSuffix(stemmedWord)
	}

	optionalLetter := suffix.GetOptionalLetter(stemmedWord)
	if optionalLetter != nil {
		if ValidateOptionalLetter(stemmedWord, optionalLetter) {
			runes := []rune(stemmedWord)
			stemmedWord = string(runes[:len(runes)-1])
		} else {
			stemmedWord = word
		}
	}
	return stemmedWord
}

// shouldBeMarked Returns whether the word should be stem or not.
func (s Stemmer) shouldBeMarked(word string, suffix Suffix) bool {
	return !Contains(s.ProtectedWords, word) &&
		(suffix.CheckHarmony && HasVowelHarmony(word) ||
			Contains(s.VowelHarmonyExceptions, word) ||
			!suffix.CheckHarmony)
}

// TODO: Refactor
func (s Stemmer) postProcess(stems []string, word string) string {
	var finalStems Stems
	for _, w := range stems {
		if w != word && CountSyllables(w) > 0 {
			finalStems = append(finalStems, s.lastConsonant(w))
		}
	}

	// Custom Sort
	sort.Slice(finalStems, func(i, j int) bool {
		if Contains(s.AverageStemSizeExceptions, finalStems[i]) {
			return true
		}
		if Contains(s.AverageStemSizeExceptions, finalStems[j]) {
			return false
		}
		s1Len, s2Len := len([]rune(finalStems[i])), len([]rune(finalStems[j]))

		averageDistance := math.Abs(float64(s1Len-AverageStemmerCount)) - math.Abs(float64(s2Len-AverageStemmerCount))
		if averageDistance == 0 {
			return (s1Len - s2Len) < 0
		} else {
			return averageDistance < 0
		}
	})

	if len(finalStems) > 0 {
		return finalStems[0]
	}
	return word
}

// validateWord Checks whether a word is acceptable for stemming or not.
func (s Stemmer) validateWord(word string) bool {
	word = strings.TrimSpace(word)

	if len(word) < 1 ||
		Contains(s.ProtectedWords, word) ||
		!IsTurkishWord(word) ||
		CountSyllables(word) < MinSyllableCount {
		return false
	}

	return true
}

// lastConsonant Checks the last consonant rule of a word
// returns a new word affected by the last consonant rule
func (s Stemmer) lastConsonant(word string) string {
	if Contains(s.LastConsonantExceptions, word) {
		return word
	}
	w := []rune(word)
	wordLen := len([]rune(word))
	lastChar := w[wordLen-1]
	if replaceChar, ok := LastConsonantRules[string(lastChar)]; ok {
		return string(w[:wordLen-1]) + replaceChar
	}

	return word
}
