package seq

type Seq interface {
	at(index int) Alphabet
	len() int
	append(seq Seq) bool
	lappend(seq Seq) bool
	insert(loc int, seq Seq) bool
}


type ProteinSeq struct {
	seq []AmnioAcid
}

type GenomeSeq struct {
	seq []NucleicAcid
}

func (p *GenomeSeq) lappend(seq Seq) bool {
	target, ok := seq.(*GenomeSeq)
	if ok {
		result := make([]NucleicAcid, 0, seq.len() + p.len())
		result = append(result, p.seq...)
		result = append(result, target.seq...)
		p.seq = result
		return true
	}
	return false
}

func (p *GenomeSeq) insert(loc int, seq Seq) bool {
	target := seq.(*GenomeSeq)
	if ok {
		result := make([]NucleicAcid, 0, seq.len() + p.len())
		result = append(result, p.seq[:loc]...)
		result = append(result, target.seq...)
		result = append(result, p.seq[loc:]...)
		p.seq = result
		return true
	}
	return false
}

func (p ProteinSeq) len() int  {
	return len(p.seq)
}

func (p* ProteinSeq) lappend(seq Seq) bool  {
	target, ok := seq.(*ProteinSeq)
	if ok {
		result := make([]AmnioAcid, 0, seq.len() + p.len())
		result = append(result, p.seq...)
		result = append(result, target.seq...)
		p.seq = result
		return true
	}
	return false
}

func (p* ProteinSeq) insert(loc int, seq Seq) bool {
	target, ok := seq.(*ProteinSeq)
	if ok {
		result := make([]AmnioAcid, 0, seq.len() + p.len())
		result = append(result, p.seq[:loc]...)
		result = append(result, target.seq...)
		result = append(result, p.seq[loc:]...)
		p.seq = result
		return true
	}
	return false
}


func (p* ProteinSeq) append(seq Seq) bool{
	target, ok := seq.(*ProteinSeq)
	if ok {
		p.seq = append(p.seq,target.seq...)
		return true
	}
	return false
}

func (p ProteinSeq) at(index int) Alphabet {
	return Alphabet(p.seq[index])
}

func (p *GenomeSeq) at(index int) Alphabet {
	return Alphabet(p.seq[index])
}

func (p* GenomeSeq) append(seq Seq) bool{
	target, ok := seq.(*GenomeSeq)
	if ok {
		p.seq = append(p.seq,target.seq...)
		return true
	}
	return false
}


func (p *GenomeSeq) len() int {
	return len(p.seq)
}

func NewGeneSeq(seq ...NucleicAcid)  *GenomeSeq {
	return &GenomeSeq{
		seq:  seq,
	}
}


//
//type Seq struct {
//	isReference bool
//	seqType    SequenceType;
//	characters []Alphabet;
//}
//
//func (s Seq) seq() []Alphabet {
//	return s.characters
//}
//
//func (s Seq) copy() Sequence {
//	return Seq {
//		seqType: s.seqType,
//		characters: s.characters,
//	}
//}
//
//func (s Seq) ref() bool {
//	return s.isReference;
//}
//
//func (s Seq) len() int {
//	return len(s.characters)
//}
//
//func (s Seq) at(loc int) Alphabet {
//	return s.characters[loc]
//}
//
//func newSeq(c[] Alphabet, seqType SequenceType) Sequence {
//	return Seq {
//		seqType: seqType,
//		characters: c,
//	}
//}
//
//func newReserveSeq(seqType SequenceType, reserve int) Sequence {
//	return Seq {
//		seqType: seqType,
//		characters: make([]Alphabet, 0, reserve),
//	}
//}
//
//func newSeqByString(value string, seqType SequenceType) Sequence {
//	result := make([]Alphabet, 0, len(value))
//	mm := map[rune] Alphabet {
//		'A' : A,
//		'C' : C,
//		'G' : G,
//		'T' : T,
//		'U' : U,
//		'K' : K,
//		'S' : S,
//		'Y' : Y,
//		'V' : V,
//		'H' : H,
//		'D' : D,
//		'B' : B,
//		'N' : N,
//	}
//	for _, char := range value {
//		if val, ok := mm[char]; ok {
//			result = append(result, val)
//		} else {
//			result = append(result, Gap)
//		}
//	}
//	return Seq {
//		seqType: seqType,
//		characters: result,
//	}
//}
//
