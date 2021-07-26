package seq

type SequenceType uint8

const (
	Unknown SequenceType = iota
	DNA     SequenceType = iota
	RNA     SequenceType = iota
	Protein SequenceType = iota
)

type Alphabet uint16

//type NucleicAcid uint16
//type AmnioAcid Alphabet
//type Alphabet

const (
	Gap Alphabet = 0
	A   Alphabet = 1 << iota
	C   Alphabet = 1 << iota
	G   Alphabet = 1 << iota
	T   Alphabet = 1 << iota
	U   Alphabet = 1 << iota

	K Alphabet = A | T
	S Alphabet = G | C
	Y Alphabet = T | C
	V Alphabet = A | C | G
	H Alphabet = A | C | T
	D Alphabet = A | G | T
	B Alphabet = C | G | T
	N Alphabet = A | C | G | T
)

const (
	Ala Alphabet = iota
	Asx Alphabet = iota
	Cys Alphabet = iota
	Asp Alphabet = iota
	Glu Alphabet = iota
	Phe Alphabet = iota
	Gly Alphabet = iota
	His Alphabet = iota
	Ile Alphabet = iota
	Lys Alphabet = iota
	Leu Alphabet = iota
	Met Alphabet = iota
	Asn Alphabet = iota
	Pro Alphabet = iota
	Gln Alphabet = iota
	Arg Alphabet = iota
	Ser Alphabet = iota
	Thr Alphabet = iota
	Val Alphabet = iota
	Trp Alphabet = iota
	Xxx Alphabet = iota
	Tyr Alphabet = iota
	Glx Alphabet = iota
)
