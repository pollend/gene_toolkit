package seq

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
	Ala AmnioAcid = iota
	Asx AmnioAcid = iota
	Cys AmnioAcid = iota
	Asp AmnioAcid = iota
	Glu AmnioAcid = iota
	Phe AmnioAcid = iota
	Gly AmnioAcid = iota
	His AmnioAcid = iota
	Ile AmnioAcid = iota
	Lys AmnioAcid = iota
	Leu AmnioAcid = iota
	Met AmnioAcid = iota
	Asn AmnioAcid = iota
	Pro AmnioAcid = iota
	Gln AmnioAcid = iota
	Arg AmnioAcid = iota
	Ser AmnioAcid = iota
	Thr AmnioAcid = iota
	Val AmnioAcid = iota
	Trp AmnioAcid = iota
	Xxx AmnioAcid = iota
	Tyr AmnioAcid = iota
	Glx AmnioAcid = iota
)
