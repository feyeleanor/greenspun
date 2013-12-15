package greenspun

//import "fmt"

/*
	This is a wrapper for a spaghetti stack data structure as implemented by the stackCell type. It's
	inspired by Go's SliceHeader and StringHeader types.

	Whilst bare spaghetti stack structures are immutable containers, the Lifo implements Lisp's
	Rplaca and Rplacd functions and allows stack elements to be modified in situ.

	TODO: consider ditching mutability and making methods return new lifo headers
*/

type Lifo struct {
	*stackCell
	depth		int
}

func Stack(items... interface{}) *Lifo {
	return &Lifo{ depth: len(items), stackCell: stack(items...) }
}

func (s *Lifo) String() (r string) {
	if s != nil {
		r = s.stackCell.String()
	} else {
		r = "()"
	}
	return
}

func (s *Lifo) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Lifo:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil:
			r = s.stackCell == nil
		case s == nil && o != nil:
			r = o.stackCell == nil
		case s.depth != o.depth:
			r = false
		default:
			x, y := s.stackCell, o.stackCell
			for r = true ; x != nil && r; x, y = x.stackCell, y.stackCell {
				if v, ok := x.data.(Equatable); ok {
					r = v.Equal(y.data)
				} else if v, ok := y.data.(Equatable); ok {
					r = v.Equal(x.data)
				} else {
					r = x.data == y.data
				}
			}
		}
	case *stackCell:
		if s != nil {
			r = s.stackCell.Equal(o)
		} else {
			r = o == nil
		}
	case nil:
		r = s == nil || s.stackCell == nil
	}
	return
}

func (s *Lifo) Push(item interface{}) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Push(item)
	s.depth++
}

func (s *Lifo) Peek() interface{} {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	return s.stackCell.Peek()
}

func (s *Lifo) Pop() (r interface{}) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	r, s.stackCell = s.stackCell.Pop()
	s.depth--
	return
}

func (s *Lifo) Len() (r int) {
	if s != nil {
	 	r = s.depth
	}
	return
}

func (s *Lifo) IsNil() (r bool) {
	return s == nil
}

func (s *Lifo) Drop() {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	if s.stackCell != nil {
		s.stackCell = s.stackCell.stackCell
		s.depth--
	}
}

func (s *Lifo) Dup() {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Dup()
	s.depth++
}

func (s *Lifo) Swap() {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Swap()
}

/*
	Make a new stack containing n cells where each cell contains the same value as is stored at the same depth
	in the existing stack.

	If the stack is shorter than n we copy the entire stack.
*/
func (s *Lifo) Copy(n int) (r *Lifo) {
	if s != nil {
		if n > s.depth {
			n = s.depth
		}
	 	r = &Lifo{ stackCell: s.stackCell.Copy(n), depth: n }
	}
 	return
}

/*
	Move to the Nth cell from the top of the stack, or return an error if there are fewer than N cells.
*/
func (s *Lifo) Move(n int) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Move(n)
	s.depth -= n
}

/*
	Move to the Nth cell from the top of the stack and create a new cell with the same value and pointing to the
	top of the current stack.
*/
func (s *Lifo) Pick(n int) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Push(s.stackCell.Move(n).data)
	s.depth++
}

/*
	Create a new stack common with the current stack from the Nth+1 element. The Nth item of the current stack becames
	the first item of the new stack and then successive elements are filled with corresponding values starting with
	that at the top of the current stack.
*/
func (s *Lifo) Roll(n int) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Roll(n)
}

/*
	Replace the data item stored in the cell at the top of the stack.
*/
func (s *Lifo) Rplaca(item interface{}) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	s.stackCell.data = item
}

/*
	Change the stack pointed to by the top of the stack.
*/
func (s *Lifo) Rplacd(tail interface{}) {
	switch {
	case s == nil:
		panic(LIST_UNINITIALIZED)
	case s.stackCell == nil:
		panic(LIST_EMPTY)
	}
	switch tail := tail.(type) {
	case *stackCell:
		s.stackCell.stackCell = tail
		s.depth = tail.Len() + 1
	case *Lifo:
		s.stackCell.stackCell = tail.stackCell
		s.depth = tail.depth + 1
	case nil:
		s.stackCell.stackCell = nil
		s.depth = 1
	default:
		panic(LIST_REQUIRED)
	}
}