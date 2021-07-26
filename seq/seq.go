package seq

type Sequence interface {
	Operations
	copy() Sequence
	len() int
	at(loc int) Alphabet
	ref() bool
	seq() []Alphabet
}

type Seq struct {
	isReference bool
	seqType    SequenceType;
	characters []Alphabet;
}

func (s Seq) seq() []Alphabet {
	return s.characters
}

func (s Seq) copy() Sequence {
	return Seq {
		seqType: s.seqType,
		characters: s.characters,
	}
}

func (s Seq) ref() bool {
	return s.isReference;
}

func (s Seq) len() int {
	return len(s.characters)
}

func (s Seq) at(loc int) Alphabet {
	return s.characters[loc]
}

func newSeq(c[] Alphabet, seqType SequenceType) Sequence {
	return Seq {
		seqType: seqType,
		characters: c,
	}
}

func newReserveSeq(seqType SequenceType, reserve int) Sequence {
	return Seq {
		seqType: seqType,
		characters: make([]Alphabet, 0, reserve),
	}
}

func newSeqByString(value string, seqType SequenceType) Sequence {
	result := make([]Alphabet, 0, len(value))
	mm := map[rune] Alphabet {
		'A' : A,
		'C' : C,
		'G' : G,
		'T' : T,
		'U' : U,
		'K' : K,
		'S' : S,
		'Y' : Y,
		'V' : V,
		'H' : H,
		'D' : D,
		'B' : B,
		'N' : N,
	}
	for _, char := range value {
		if val, ok := mm[char]; ok {
			result = append(result, val)
		} else {
			result = append(result, Gap)
		}
	}
	return Seq {
		seqType: seqType,
		characters: result,
	}
}

