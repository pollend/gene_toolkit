package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	res := NewGeneSeq(A,C,G)
	res.Append(NewGeneSeq(A,C))

	assert.Equal(t, res.seq, NewGeneSeq(A,C,G,A,C).seq)
}

func TestEquality(t *testing.T)  {
	//assert.True(t, newSeq([] Alphabet{A,C}, DNA).equal(newSeq([] Alphabet{A,C}, DNA)))
	//assert.False(t, newSeq([] Alphabet{A,C}, DNA).equal(newSeq([] Alphabet{A,C,G}, DNA)))
}
