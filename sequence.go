package greenspun

type Sequence interface {
	Peek() interface{}
	Next() Sequence
}

//	Check the equality of two cells based upon their contents
//
func MatchValue(lhs, rhs Sequence) (r bool) {
	switch {
	case lhs == nil && rhs == nil:
		r = true
	case lhs != nil && rhs != nil:
		if v, ok := lhs.Peek().(Equatable); ok {
			r = v.Equal(rhs.Peek())
		} else if v, ok := rhs.Peek().(Equatable); ok {
			r = v.Equal(lhs.Peek())
		} else {
			r = lhs.Peek() == rhs.Peek()
		}
	}
	return
}
