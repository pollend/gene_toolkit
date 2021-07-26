package seq

import "errors"

type Operations interface {
	equal(seq Sequence) bool
	count(seq Sequence) int
	add(seq Sequence) Sequence
	multi(cnt int) Sequence
	slice(start int, end int) (Sequence, error);
}

func (s Seq) equal(target Sequence) bool {
	if s.len() != target.len() {
		return false;
	}
	for i, val := range s.seq() {
		if val != target.at(i) {
			return false
		}
	}
	return true
}

func (s Seq) multi(cnt int) Sequence {
	result := make([]Alphabet, 0, cnt * s.len())
	copy(result, s.characters)
	for i := 0; i < cnt; i++ {
		result = append(result, s.characters...)
	}
	s.characters = result
	return s
}
func (s Seq) add(seq Sequence) Sequence {
	s.characters = append(s.characters, seq.seq()...)
	return s
}

func (s Seq) slice(start int, end int) (Sequence, error) {
	if(start > end) {
		return nil, errors.New("start can't greater then end")
	}
	if start > 0 && start < s.len() && end > 0 && end < s.len() {
		return Seq {
			isReference: true,
			seqType: s.seqType,
			characters: s.characters[start:end],
		}, nil
	}
	return nil, errors.New("")
}


func (s Seq) count(target Sequence) int {
	i := 0
	count := 0
	for i + target.len() < s.len(){
		res, _ := s.slice(i,i + target.len())
		if target.equal(res) {
			i += target.len()
			count++;
		}
		i++
	}
	return count
}
