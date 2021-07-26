package seq

type Matcher func(a1 Sequence, a2 Sequence) bool
func strictMatch() Matcher {
	return func(a1 Sequence, a2 Sequence) bool {
		if a1.len() != a2.len() {
			return false;
		}
		for i, val := range a1.seq() {
			if val != a2.at(i) {
				return false
			}
		}
		return true
	}
}
