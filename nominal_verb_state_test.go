package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNominalVerbState_GetInitialNominalVerbState(t *testing.T) {
	assert.Equal(t, NominalVerbStateA, GetInitialNominalVerbState())
}

func TestNominalVerbState_NextState(t *testing.T) {
	assert.Equal(t, NominalVerbStateB, NominalVerbStateA.NextState(NominalVerbSuffix1.String()))
	assert.Nil(t, NominalVerbStateF.NextState(NominalVerbSuffix1.String()))
	assert.Nil(t, NominalVerbStateF.NextState("RandomSuffix"))
}

func TestNominalVerbState_Suffixes(t *testing.T) {
	assert.Contains(t, NominalVerbStateA.Suffixes(), NominalVerbSuffix1)
	assert.NotContains(t, NominalVerbStateB.Suffixes(), NominalVerbSuffix11)
	assert.Len(t, NominalVerbStateF.Suffixes(), 0)
}

func TestNominalVerbState_AddTransitions(t *testing.T) {
	var transitions Transitions

	NominalVerbStateA.AddTransitions("satıyorsunuz", &transitions, NominalVerbStateA)

	assert.Len(t, transitions, 3)

	assert.Equal(t, NominalVerbStateA, transitions[0].StartState)
	assert.Equal(t, NominalVerbStateB, transitions[0].NextState)
	assert.Equal(t, NominalVerbSuffix4, transitions[0].Suffix)

	assert.Equal(t, NominalVerbStateA, transitions[1].StartState)
	assert.Equal(t, NominalVerbStateD, transitions[1].NextState)
	assert.Equal(t, NominalVerbSuffix9, transitions[1].Suffix)

	assert.Equal(t, NominalVerbStateA, transitions[2].StartState)
	assert.Equal(t, NominalVerbStateB, transitions[2].NextState)
	assert.Equal(t, NominalVerbSuffix3, transitions[2].Suffix)
}
