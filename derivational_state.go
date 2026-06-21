package turkishstemmer

type DerivationalState struct {
	BaseState
}

var (
	DerivationalStateA = NewDerivationalState("DerivationalStateA", true, false, DerivationalSuffixValues)
	DerivationalStateB = NewDerivationalState("DerivationalStateB", false, true, nil)
)

func NewDerivationalState(id string, initialState, finalState bool, suffixes []Suffix) DerivationalState {
	return DerivationalState{
		BaseState{
			id:           id,
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
