package turkishstemmer

import (
	"fmt"
	"reflect"
)

type Transitions []*Transition
type Transition struct {
	StartState State
	NextState  State
	Word       string
	Suffix     Suffix
	Marked     bool
}

func NewTransition(startState, nextState State, word string, suffix Suffix) *Transition {
	return &Transition{
		StartState: startState,
		NextState:  nextState,
		Word:       word,
		Suffix:     suffix,
		Marked:     false,
	}
}

func (t Transition) SimilarTransitions(transitions Transitions) *Transitions {
	var similars Transitions

	for _, transition := range transitions {
		if reflect.DeepEqual(t.StartState, transition.StartState) &&
			reflect.DeepEqual(t.NextState, transition.NextState) {
			similars = append(similars, transition)
		}
	}

	return &similars
}

func (t Transition) String() string {
	return fmt.Sprintf("%s(%s) -> %s", GetType(t.StartState), t.Suffix, GetType(t.NextState))
}
