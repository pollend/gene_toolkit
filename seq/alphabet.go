package seq

type SequenceType uint8

const (
	Unknown SequenceType = iota
	DNA     SequenceType = iota
	RNA     SequenceType = iota
	Protein SequenceType = iota
)

type Alphabet uint16

type NucleicAcid Alphabet
type AmnioAcid Alphabet

const (
	Gap NucleicAcid = 0
	A   NucleicAcid = 1 << iota
	C   NucleicAcid = 1 << iota
	G   NucleicAcid = 1 << iota
	T   NucleicAcid = 1 << iota
	U   NucleicAcid = 1 << iota

	K NucleicAcid = A | T
	S NucleicAcid = G | C
	Y NucleicAcid = T | C
	V NucleicAcid = A | C | G
	H NucleicAcid = A | C | T
	D NucleicAcid = A | G | T
	B NucleicAcid = C | G | T
	N NucleicAcid = A | C | G | T
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
