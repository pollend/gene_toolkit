package seq

//type IGenomeSeq interface {
//
//}
//
//func (p *GenomeSeq) At(index int) NucleicAcid{
//	return p.seq[index]
//}
//
//func (p *GenomeSeq) Append(seq* GenomeSeq)  {
//	p.Append(seq)
//}
//
//func (p *GenomeSeq) Len() int  {
//	return p.Len()
//}
//
//func (p *GenomeSeq) LeftAppend(seq* GenomeSeq) {
//	p.LAppend(seq)
//}


//
//type Operations interface {
//	equal(seq Sequence) bool
//	count(seq Sequence) int
//	add(seq Sequence) Sequence
//	multi(cnt int) Sequence
//	slice(start int, end int) (Sequence, error);
//}
//
//func (s ISeq) equal(target Sequence) bool {
//	if s.Len() != target.Len() {
//		return false;
//	}
//	for i, val := range s.seq() {
//		if val != target.At(i) {
//			return false
//		}
//	}
//	return true
//}
//
//func (s ISeq) multi(cnt int) Sequence {
//	result := make([]Alphabet, 0, cnt * s.Len())
//	copy(result, s.characters)
//	for i := 0; i < cnt; i++ {
//		result = Append(result, s.characters...)
//	}
//	s.characters = result
//	return s
//}
//func (s ISeq) add(seq Sequence) Sequence {
//	s.characters = Append(s.characters, seq.seq()...)
//	return s
//}
//
//func (s ISeq) slice(start int, end int) (Sequence, error) {
//	if(start > end) {
//		return nil, errors.New("start can't greater then end")
//	}
//	if start > 0 && start < s.Len() && end > 0 && end < s.Len() {
//		return ISeq {
//			isReference: true,
//			seqType: s.seqType,
//			characters: s.characters[start:end],
//		}, nil
//	}
//	return nil, errors.New("")
//}
//
//
//func (s ISeq) count(target Sequence) int {
//	i := 0
//	count := 0
//	for i + target.Len() < s.Len(){
//		res, _ := s.slice(i,i + target.Len())
//		if target.equal(res) {
//			i += target.Len()
//			count++;
//		}
//		i++
//	}
//	return count
//}
