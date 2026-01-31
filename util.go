package turkishstemmer

import (
	"reflect"
	"strings"
)

// contains Returns whether e is within s.
func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ReplaceStringAtIndex Find exact index number in text then replaces with r and returns new text.
func ReplaceStringAtIndex(text string, r rune, i int) string {
	newText := []rune(text)
	newText[i] = r
	return string(newText)
}

// GetVowels returns the vowels of a word.
func GetVowels(word string) []rune {
	var vowels []rune
	for _, char := range word {
		if strings.Contains(Vowels, string(char)) {
			vowels = append(vowels, char)
		}
	}

	return vowels
}

// CountSyllables Returns the number/count of syllables of a word.
func CountSyllables(word string) int {
	return len(GetVowels(word))
}

// ValidateOptionalLetter Checks whether an optional letter is valid or not.
func ValidateOptionalLetter(word string, candidate *rune) bool {
	wordLength := len([]rune(word))

	if wordLength-2 < 0 {
		return false
	}

	previousChar := string([]rune(word)[wordLength-2])

	if strings.Contains(Vowels, string(*candidate)) {
		return strings.Contains(Consonants, previousChar)
	}

	return strings.Contains(Vowels, previousChar)
}

// VowelHarmony Checks the vowel harmony of two characters.
func VowelHarmony(vowel, candidate string) bool {
	return HasRoundness(vowel, candidate) && HasFrontness(vowel, candidate)
}

// IsTurkishWord Checks whether a word is written in Turkish alphabet or not.
func IsTurkishWord(word string) bool {
	for _, char := range word {
		if !strings.Contains(Alphabet, string(char)) {
			return false
		}
	}

	return true
}

// HasFrontness Checks the frontness harmony of two characters.
func HasFrontness(vowel, candidate string) bool {
	return (strings.Contains(FrontVowels, vowel) && strings.Contains(FrontVowels, candidate)) ||
		(strings.Contains(BackVowels, vowel) && strings.Contains(BackVowels, candidate))
}

// HasRoundness Checks the roundness harmony of two characters.
func HasRoundness(vowel, candidate string) bool {
	return (strings.Contains(UnroundedVowels, vowel) && strings.Contains(UnroundedVowels, candidate)) ||
		(strings.Contains(RoundedVowels, vowel) && strings.Contains(FollowingRoundedVowels, candidate))
}

// HasVowelHarmony Checks the vowel harmony of a word.
func HasVowelHarmony(word string) bool {
	vowels := GetVowels(word)
	wordLength := len(vowels)
	if wordLength-2 < 0 {
		return true
	}

	vowel := string(vowels[wordLength-2])
	candidate := string(vowels[wordLength-1])

	return VowelHarmony(vowel, candidate)
}

// GetType Returns type of interface.
func GetType(i interface{}) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
