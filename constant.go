package turkishstemmer

import _ "embed"

const (
	// The average size of turkish stems based on which the selection of the final stem is performed.
	// The idea behind the selection process is based on the paper
	// F.Can, S.Kocberber, E.Balcik, C.Kaynak, H.Cagdas, O.Calan, O.Vursavas
	// "Information Retrieval on Turkish Texts"

	AverageStemmerCount = 4
	MinSyllableCount    = 2
	// Alphabet Turkish alphabet. They are used for skipping not turkish words.
	Alphabet = "abcçdefgğhıijklmnoöprsştuüvyz"
	// Vowels Turkish vowels.
	Vowels = "üiıueöao"
	// Consonants Turkish consonants.
	Consonants = "bcçdfgğhjklmnprsştvyz"
	// RoundedVowels Rounded vowels which are used for checking roundness harmony.
	RoundedVowels = "oöuü"
	// FollowingRoundedVowels Vowels that follow rounded vowels.
	// They are combined with ROUNDED_VOWELS to check roundness harmony.
	FollowingRoundedVowels = "aeuü"
	// UnroundedVowels The unrounded vowels which are used for checking roundness harmony.
	UnroundedVowels = "iıea"
	// FrontVowels Front vowels which are used for checking frontness harmony.
	FrontVowels = "eiöü"
	// BackVowels Front vowels which are used for checking frontness harmony.
	BackVowels = "ıuao"
)

var (
	//go:embed data/protected_words.txt
	// DefaultProtectedWordsFile The path of the file that contains the default set of protected words.
	DefaultProtectedWordsFile []byte

	//go:embed data/vowel_harmony_exceptions.txt
	// DefaultVowelHarmonyExceptionsFile The path of the file that contains the default set of vowel harmony exceptions.
	DefaultVowelHarmonyExceptionsFile []byte

	//go:embed data/last_consonant_exceptions.txt
	// DefaultLastConsonantExceptionsFile The path of the file that contains the default set of last consonant exceptions.
	DefaultLastConsonantExceptionsFile []byte

	//go:embed data/average_stem_size_exceptions.txt
	// DefaultAverageStemSizeExceptionsFile The path of the file that contains the default set of average stem size exceptions.
	DefaultAverageStemSizeExceptionsFile []byte

	// LastConsonantRules Last consonant rules
	LastConsonantRules = map[string]string{"b": "p", "c": "ç", "d": "t", "ğ": "k"}
)
