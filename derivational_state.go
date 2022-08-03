package turkishstemmer

type DerivationalState struct {
	BaseState
}

var (
	DerivationalStateA = NewDerivationalState(true, false, DerivationalSuffixValues)
	DerivationalStateB = NewDerivationalState(false, true, nil)
)

func NewDerivationalState(initialState, finalState bool, suffixes []Suffix) DerivationalState {
	return DerivationalState{
		BaseState{
			initialState: initialState,
			finalState:   finalState,
			suffixes:     suffixes,
		},
	}
}

func (s DerivationalState) NextState(suffix string) State {
	if len(s.suffixes) > 0 && s.initialState && DerivationalSuffix1.String() == suffix {
		return DerivationalStateB
	}

	return nil
}
