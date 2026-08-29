package turkishstemmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateID(t *testing.T) {
	assert.Equal(t, "NounStateA", NounStateA.ID())
	assert.Equal(t, "NominalVerbStateA", NominalVerbStateA.ID())
	assert.Equal(t, "DerivationalStateA", DerivationalStateA.ID())
	assert.NotEqual(t, NounStateA.ID(), NounStateB.ID())
}

func TestStatesEqual(t *testing.T) {
	assert.True(t, statesEqual(NounStateA, NounStateA))
	assert.False(t, statesEqual(NounStateA, NounStateB))
	assert.True(t, statesEqual(nil, nil))
	assert.False(t, statesEqual(DerivationalStateA, nil))
}
