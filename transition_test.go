package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransition_SimilarTransitions_NominalVerb(t *testing.T) {
	similarTransition := NewTransition(NominalVerbStateA, NominalVerbStateB, "someWord", NominalVerbSuffix2)
	differentTransition := NewTransition(NominalVerbStateC, NominalVerbStateD, "someWord", NominalVerbSuffix2)

	transitions := Transitions{similarTransition, differentTransition}
	transition := NewTransition(NominalVerbStateA, NominalVerbStateB, "aword", NominalVerbSuffix1)

	assert.Equal(t, transition.SimilarTransitions(transitions), &Transitions{similarTransition})
}

func TestTransition_SimilarTransitions_Noun(t *testing.T) {
	similarTransition := NewTransition(NounStateA, NounStateB, "someWord", NounSuffix2)
	differentTransition := NewTransition(NounStateC, NounStateD, "someWord", NounSuffix2)

	transitions := Transitions{similarTransition, differentTransition}

	transition := NewTransition(NounStateA, NounStateB, "aword", NounSuffix1)
	assert.Equal(t, transition.SimilarTransitions(transitions), &Transitions{similarTransition})
}

func TestTransition_SimilarTransitions_Derivational(t *testing.T) {
	similarTransition := NewTransition(DerivationalStateA, DerivationalStateB, "someWord", DerivationalSuffix1)
	differentTransition := NewTransition(DerivationalStateA, nil, "someWord", DerivationalSuffix1)

	transitions := Transitions{similarTransition, differentTransition}

	transition := NewTransition(DerivationalStateA, DerivationalStateB, "aword", DerivationalSuffix1)
	assert.Equal(t, transition.SimilarTransitions(transitions), &Transitions{similarTransition})
}

func TestTransition_String(t *testing.T) {
	transition := NewTransition(DerivationalStateA, DerivationalStateB, "someWord", DerivationalSuffix1)

	assert.Equal(t, "DerivationalState(-lU) -> DerivationalState", transition.String())
}
