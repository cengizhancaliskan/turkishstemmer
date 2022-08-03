package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDerivationalState_NextState(t *testing.T) {
	assert.Equal(t, DerivationalStateB, DerivationalStateA.NextState(DerivationalSuffix1.String()))
	assert.Nil(t, DerivationalStateB.NextState(DerivationalSuffix1.String()))
	assert.Nil(t, DerivationalStateB.NextState("DerivationalSuffix"))
}

func TestDerivationalState_AddTransitions(t *testing.T) {
	var transitions Transitions

	DerivationalStateA.AddTransitions("gozlu", &transitions, DerivationalStateA)

	assert.Len(t, transitions, 1)

	assert.Equal(t, DerivationalStateA, transitions[0].StartState)
	assert.Equal(t, DerivationalStateB, transitions[0].NextState)
	assert.Equal(t, DerivationalSuffix1, transitions[0].Suffix)
}

func TestDerivationalState_InitialState(t *testing.T) {
	assert.True(t, DerivationalStateA.InitialState())
	assert.False(t, DerivationalStateB.InitialState())
}

func TestDerivationalState_FinalState(t *testing.T) {
	assert.False(t, DerivationalStateA.FinalState())
	assert.True(t, DerivationalStateB.FinalState())
}

func TestDerivationalState_Suffixes(t *testing.T) {
	assert.Contains(t, DerivationalStateA.Suffixes(), DerivationalSuffix1)
	assert.NotContains(t, DerivationalStateB.Suffixes(), DerivationalSuffix1)
	assert.Len(t, DerivationalStateB.Suffixes(), 0)
}
