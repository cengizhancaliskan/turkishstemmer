package turkishstemmer

type NominalVerbState struct {
	BaseState
}

var (
	NominalVerbStateA = NewNominalVerbState(true, false, NominalVerbSuffixValues)
	NominalVerbStateB = NewNominalVerbState(false, true, []Suffix{NominalVerbSuffix14})
	NominalVerbStateC = NewNominalVerbState(false, true, []Suffix{NominalVerbSuffix10, NominalVerbSuffix12, NominalVerbSuffix13, NominalVerbSuffix14}) //nolint:lll
	NominalVerbStateD = NewNominalVerbState(false, false, []Suffix{NominalVerbSuffix12, NominalVerbSuffix13})
	NominalVerbStateE = NewNominalVerbState(false, true, []Suffix{NominalVerbSuffix1, NominalVerbSuffix2, NominalVerbSuffix3, NominalVerbSuffix4, NominalVerbSuffix5, NominalVerbSuffix14}) //nolint:lll
	NominalVerbStateF = NewNominalVerbState(false, true, nil)
	NominalVerbStateG = NewNominalVerbState(false, false, []Suffix{NominalVerbSuffix14})
	NominalVerbStateH = NewNominalVerbState(false, false, []Suffix{NominalVerbSuffix1, NominalVerbSuffix2, NominalVerbSuffix3, NominalVerbSuffix4, NominalVerbSuffix5, NominalVerbSuffix14}) //nolint:lll

	NominalVerbTFValues = map[string]NominalVerbState{
		NominalVerbSuffix1.Name:  NominalVerbStateB,
		NominalVerbSuffix2.Name:  NominalVerbStateB,
		NominalVerbSuffix3.Name:  NominalVerbStateB,
		NominalVerbSuffix4.Name:  NominalVerbStateB,
		NominalVerbSuffix5.Name:  NominalVerbStateC,
		NominalVerbSuffix6.Name:  NominalVerbStateD,
		NominalVerbSuffix7.Name:  NominalVerbStateD,
		NominalVerbSuffix8.Name:  NominalVerbStateD,
		NominalVerbSuffix9.Name:  NominalVerbStateD,
		NominalVerbSuffix10.Name: NominalVerbStateE,
		NominalVerbSuffix11.Name: NominalVerbStateH,
		NominalVerbSuffix12.Name: NominalVerbStateF,
		NominalVerbSuffix13.Name: NominalVerbStateF,
		NominalVerbSuffix14.Name: NominalVerbStateF,
		NominalVerbSuffix15.Name: NominalVerbStateF,
	}

	NominalVerbFTValues = map[string]NominalVerbState{
		NominalVerbSuffix1.Name:  NominalVerbStateG,
		NominalVerbSuffix2.Name:  NominalVerbStateG,
		NominalVerbSuffix3.Name:  NominalVerbStateG,
		NominalVerbSuffix4.Name:  NominalVerbStateG,
		NominalVerbSuffix5.Name:  NominalVerbStateG,
		NominalVerbSuffix10.Name: NominalVerbStateF,
		NominalVerbSuffix13.Name: NominalVerbStateF,
		NominalVerbSuffix14.Name: NominalVerbStateF,
	}

	NominalVerbFFValues = map[string]NominalVerbState{
		NominalVerbSuffix1.Name:  NominalVerbStateG,
		NominalVerbSuffix2.Name:  NominalVerbStateG,
		NominalVerbSuffix3.Name:  NominalVerbStateG,
		NominalVerbSuffix4.Name:  NominalVerbStateG,
		NominalVerbSuffix5.Name:  NominalVerbStateG,
		NominalVerbSuffix12.Name: NominalVerbStateF,
		NominalVerbSuffix14.Name: NominalVerbStateF,
	}
)

func GetInitialNominalVerbState() State {
	return NominalVerbStateA
}

func NewNominalVerbState(initialState, finalState bool, suffixes []Suffix) NominalVerbState {
	return NominalVerbState{
		BaseState{
			initialState: initialState,
			finalState:   finalState,
			suffixes:     suffixes,
		},
	}
}

func (s NominalVerbState) GetValues() map[string]NominalVerbState {
	if s.InitialState() && !s.FinalState() {
		return NominalVerbTFValues
	} else if s.FinalState() && !s.InitialState() {
		return NominalVerbFTValues
	}
	return NominalVerbFFValues
}

func (s NominalVerbState) NextState(suffix string) State {
	if ns, ok := s.GetValues()[suffix]; len(s.Suffixes()) > 0 && ok {
		return ns
	}

	return nil
}
