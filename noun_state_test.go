package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNounState_NextState(t *testing.T) {
	assert.Equal(t, NounStateL, NounStateA.NextState(NounSuffix1.String()))
	assert.Nil(t, NounStateK.NextState(NounSuffix1.String()))
}

func TestNounState_Suffixes(t *testing.T) {
	assert.Contains(t, NounStateA.Suffixes(), NounSuffix1)
	assert.NotContains(t, NounStateB.Suffixes(), NounSuffix11)
	assert.Len(t, NounStateK.Suffixes(), 0)
}

func TestNounState_AddTransitions(t *testing.T) {
	var transitions Transitions

	NounStateA.AddTransitions("bebekler", &transitions, NounStateA)

	assert.Len(t, transitions, 1)

	assert.Equal(t, NounStateA, transitions[0].StartState)
	assert.Equal(t, NounStateL, transitions[0].NextState)
	assert.Equal(t, NounSuffix1, transitions[0].Suffix)
}
