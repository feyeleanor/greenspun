package greenspun

//import "fmt"

/*
	This is a wrapper for a spaghetti stack data structure as implemented by the stackCell type. It's
	inspired by Go's SliceHeader and StringHeader types.

	Whilst bare spaghetti stack structures are immutable containers, the StackList implements Lisp's
	Rplaca and Rplacd functions and allows stack elements to be modified in situ.
*/

type StackList struct {
	*stackCell
	depth		int
}

func Stack(items... interface{}) *StackList {
	return &StackList{ depth: len(items), stackCell: stack(items...) }
}

func (s *StackList) String() (r string) {
	if s != nil {
		r = s.stackCell.String()
	} else {
		r = "()"
	}
	return
}

func (s *StackList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *StackList:
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

func (s *StackList) Push(item interface{}) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Push(item)
	s.depth++
}

func (s *StackList) Peek() interface{} {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	return s.stackCell.Peek()
}

func (s *StackList) Pop() (r interface{}) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	r, s.stackCell = s.stackCell.Pop()
	s.depth--
	return
}

func (s *StackList) Len() (r int) {
	if s != nil {
	 	r = s.depth
	}
	return
}

func (s *StackList) IsNil() (r bool) {
	return s == nil
}

func (s *StackList) Drop() {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	if s.stackCell != nil {
		s.stackCell = s.stackCell.stackCell
		s.depth--
	}
}

func (s *StackList) Dup() {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Dup()
	s.depth++
}

func (s *StackList) Swap() {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Swap()
}

/*
	Make a new stack containing n cells where each cell contains the same value as is stored at the same depth
	in the existing stack.

	If the stack is shorter than n we copy the entire stack.
*/
func (s *StackList) Copy(n int) (r *StackList) {
	if s != nil {
		if n > s.depth {
			n = s.depth
		}
	 	r = &StackList{ stackCell: s.stackCell.Copy(n), depth: n }
	}
 	return
}

/*
	Move to the Nth cell from the top of the stack, or return an errot if there are fewer than N cells.
*/
func (s *StackList) Move(n int) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Move(n)
	s.depth -= n
}

/*
	Move the Nth cell from the top of the stack and create a new cell with the same value and pointing to the
	top of the current stack.
*/
func (s *StackList) Pick(n int) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Push(s.stackCell.Move(n))
}

/*
	Create a new stack common with the current stack from the Nth+1 element. The Nth item of the current stack becames
	the first item of the new stack and then successive elements are filled with corresponding values starting with
	that at the top of the current stack.
*/
func (s *StackList) Roll(n int) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell = s.stackCell.Roll(n)
}

/*
	Replace the data item stored in the cell at the top of the stack.
*/
func (s *StackList) Rplaca(item interface{}) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.stackCell.data = item
}

/*
	Change the stack pointed to by the top of the stack.
*/
func (s *StackList) Rplacd(tail interface{}) {
	switch {
	case s == nil:
		panic(STACK_UNINITIALIZED)
	case s.stackCell == nil:
		panic(STACK_EMPTY)
	}
	switch tail := tail.(type) {
	case *stackCell:
		s.stackCell.stackCell = tail
		s.depth = tail.Len() + 1
	case *StackList:
		s.stackCell.stackCell = tail.stackCell
		s.depth = tail.depth + 1
	case nil:
		s.stackCell.stackCell = nil
		s.depth = 1
	default:
		panic(STACK_REQUIRED)
	}
}