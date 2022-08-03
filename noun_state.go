package turkishstemmer

type NounState struct {
	BaseState
}

var (
	NounStateA = NewNounState(true, true, NounSuffixValues)
	NounStateB = NewNounState(false, true, []Suffix{NounSuffix1, NounSuffix2, NounSuffix3, NounSuffix4, NounSuffix5})
	NounStateC = NewNounState(false, false, []Suffix{NounSuffix6, NounSuffix7})
	NounStateD = NewNounState(false, false, []Suffix{NounSuffix10, NounSuffix13, NounSuffix14})
	NounStateE = NewNounState(false, true, []Suffix{NounSuffix1, NounSuffix2, NounSuffix3, NounSuffix4, NounSuffix5, NounSuffix6, NounSuffix7, NounSuffix18}) //nolint:lll
	NounStateF = NewNounState(false, false, []Suffix{NounSuffix6, NounSuffix7, NounSuffix18})
	NounStateG = NewNounState(false, true, []Suffix{NounSuffix1, NounSuffix2, NounSuffix3, NounSuffix4, NounSuffix5, NounSuffix18}) //nolint:lll
	NounStateH = NewNounState(false, true, []Suffix{NounSuffix1})
	NounStateK = NewNounState(false, true, nil)
	NounStateL = NewNounState(false, true, []Suffix{NounSuffix18})
	NounStateM = NewNounState(false, true, []Suffix{NounSuffix1, NounSuffix2, NounSuffix3, NounSuffix4, NounSuffix5, NounSuffix6, NounSuffix6, NounSuffix7}) //nolint:lll

	// TTValues InitialState = true, FinalState = true
	TTValues = map[string]NounState{
		NounSuffix1.Name:  NounStateL,
		NounSuffix2.Name:  NounStateH,
		NounSuffix3.Name:  NounStateH,
		NounSuffix4.Name:  NounStateH,
		NounSuffix5.Name:  NounStateH,
		NounSuffix6.Name:  NounStateH,
		NounSuffix7.Name:  NounStateK,
		NounSuffix8.Name:  NounStateB,
		NounSuffix9.Name:  NounStateC,
		NounSuffix10.Name: NounStateE,
		NounSuffix11.Name: NounStateB,
		NounSuffix12.Name: NounStateF,
		NounSuffix13.Name: NounStateB,
		NounSuffix14.Name: NounStateF,
		NounSuffix15.Name: NounStateG,
		NounSuffix16.Name: NounStateC,
		NounSuffix17.Name: NounStateE,
		NounSuffix18.Name: NounStateD,
	}

	// InitialState = false, FinalState = true
	FTValues = map[string]NounState{
		NounSuffix1.Name:  NounStateL,
		NounSuffix2.Name:  NounStateH,
		NounSuffix3.Name:  NounStateH,
		NounSuffix4.Name:  NounStateH,
		NounSuffix5.Name:  NounStateH,
		NounSuffix6.Name:  NounStateH,
		NounSuffix7.Name:  NounStateK,
		NounSuffix18.Name: NounStateD,
	}

	// InitialState = false, FinalState = false
	FFValues = map[string]NounState{
		NounSuffix6.Name:  NounStateH,
		NounSuffix7.Name:  NounStateL,
		NounSuffix10.Name: NounStateE,
		NounSuffix13.Name: NounStateB,
		NounSuffix14.Name: NounStateF,
		NounSuffix18.Name: NounStateD,
	}
)

func NewNounState(initialState, finalState bool, suffixes []Suffix) NounState {
	return NounState{
		BaseState{
			initialState: initialState,
			finalState:   finalState,
			suffixes:     suffixes,
		},
	}
}

func (s NounState) GetValues() map[string]NounState {
	if s.InitialState() && s.FinalState() {
		return TTValues
	} else if s.FinalState() && !s.InitialState() {
		return FTValues
	}

	return FFValues
}

func (s NounState) NextState(suffix string) State {
	if ns, ok := s.GetValues()[suffix]; len(s.Suffixes()) > 0 && ok {
		return ns
	}

	return nil
}
