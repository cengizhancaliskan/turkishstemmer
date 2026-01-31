package turkishstemmer

type State interface {
	AddTransitions(word string, transitions *Transitions, startState State)
	NextState(suffix string) State
	InitialState() bool
	FinalState() bool
	Suffixes() []Suffix
}

type BaseState struct {
	initialState bool
	finalState   bool
	suffixes     []Suffix
}

func (s BaseState) AddTransitions(word string, transitions *Transitions, startState State) {
	for _, suffix := range s.Suffixes() {
		if suffix.Match(word) {
			*transitions = append(*transitions, NewTransition(
				startState,
				startState.NextState(suffix.String()),
				word,
				suffix,
			))
		}
	}
}

func (s BaseState) InitialState() bool {
	return s.initialState
}

func (s BaseState) FinalState() bool {
	return s.finalState
}

func (s BaseState) Suffixes() []Suffix {
	return s.suffixes
}
