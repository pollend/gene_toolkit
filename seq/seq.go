package seq

// interface for internal use within the library
type ISeq interface {
}

type IProteinSeq interface {
	ISeq
	At(index int) AmnioAcid
	Len() int
	Append(seq IProteinSeq)
	LAppend(seq IProteinSeq)
	Insert(loc int, seq IProteinSeq)
}

type IGenomeSeq interface {
	ISeq
	At(index int) NucleicAcid
	Len() int
	Append(seq IGenomeSeq)
	LAppend(seq IGenomeSeq)
	Insert(loc int, seq IGenomeSeq)
}


type ProteinSeq struct {
	seq []AmnioAcid
}

type GenomeSeq struct {
	seq []NucleicAcid
}

func (p *GenomeSeq) LAppend(seq IGenomeSeq) {
	target := seq.(*GenomeSeq)
	result := make([]NucleicAcid, 0, seq.Len() + p.Len())
	result = append(result, p.seq...)
	result = append(result, target.seq...)
	p.seq = result
}

func (p* ProteinSeq) LAppend(seq IProteinSeq)  {
	target := seq.(*ProteinSeq)
	result := make([]AmnioAcid, 0, seq.Len() + p.Len())
	result = append(result, p.seq...)
	result = append(result, target.seq...)
	p.seq = result
}

func (p *GenomeSeq) Insert(loc int, seq IGenomeSeq) {
	target := seq.(*GenomeSeq)
	result := make([]NucleicAcid, 0, seq.Len() + p.Len())
	result = append(result, p.seq[:loc]...)
	result = append(result, target.seq...)
	result = append(result, p.seq[loc:]...)
	p.seq = result
}

func (p* ProteinSeq) Insert(loc int, seq IProteinSeq) {
	target := seq.(*ProteinSeq)
	result := make([]AmnioAcid, 0, seq.Len() + p.Len())
	result = append(result, p.seq[:loc]...)
	result = append(result, target.seq...)
	result = append(result, p.seq[loc:]...)
	p.seq = result
}

func (p *GenomeSeq) At(index int) NucleicAcid {
	return p.seq[index]
}

func (p* GenomeSeq) Append(seq IGenomeSeq){
	target := seq.(*GenomeSeq)
	p.seq = append(p.seq,target.seq...)
}


func (p *GenomeSeq) Len() int {
	return len(p.seq)
}

func NewGeneSeq(seq ...NucleicAcid)  *GenomeSeq {
	return &GenomeSeq{
		seq:  seq,
	}
}

func (p ProteinSeq) Len() int  {
	return len(p.seq)
}



func (p* ProteinSeq) Append(seq IProteinSeq) {
	target := seq.(*ProteinSeq)
	p.seq = append(p.seq,target.seq...)
}

func (p ProteinSeq) At(index int) AmnioAcid {
	return p.seq[index]
}
