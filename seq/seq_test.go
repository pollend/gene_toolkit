package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	var result = newSeq([]Alphabet{A, C, G, T}, DNA)
	var target, err = result.slice(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, target.at(0), C)
}

func TestEquality(t *testing.T)  {
	assert.True(t, newSeq([] Alphabet{A,C}, DNA).equal(newSeq([] Alphabet{A,C}, DNA)))
	assert.False(t, newSeq([] Alphabet{A,C}, DNA).equal(newSeq([] Alphabet{A,C,G}, DNA)))
}
