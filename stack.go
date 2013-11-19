package greenspun

import "fmt"

/*
	This is an implementation of a classic Stack data structure, using a singly-linked list of cells.
	Each cell contains a generic item of data which must be unboxed before anything useful can be done
	with it, and a link to the previous cell in the stack.

			cf:			http://en.wikipedia.org/wiki/Spaghetti_stack

	A similar data structure could be implemented with the Pair data-structure used to represent Lisp
	Cons values. However in instances where we know we're dealing with a Stack (such as state management
	in the SECD virtual machine) we save the additional cost of unboxing the link value when stepping
	through the stack.

	A small number of additional primitive operations are implemented which will be more familiar to
	Forth programmers, allowing values to be copied from a particular place in the Stack to the top.
	The decision has been made to implement these in an immutable manner whereby a new list of cells
	is generated whenever there would otherwise be a change to existing cells.
*/

type stackCell struct {
	data	interface{}
	*stackCell	
}

/*
	A constructor function used internally to Greenspun and for testing.
*/
func stack(items... interface{}) (r *stackCell) {
	for i := len(items); i > 0; {
		i--
		r = r.Push(items[i])
	}
	return
}

/*
	Produce a string representation for the current list of cells.
*/
func (s *stackCell) String() (r string) {
	if s != nil {
		r = fmt.Sprintf("%v", s.data)
		for tos := s.stackCell; tos != nil; tos = tos.stackCell {
			r = fmt.Sprintf("%v %v", r, tos.data)
		}
	}
	return fmt.Sprintf("<%v]", r)
}

/*
	Check the equality of two lists of cells based upon their contents.
	Because lists are immutable, when two lists are confirmed to be identical then code using them can
	discard one and perform all of its operations in terms of the other.
*/
func (s *stackCell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *stackCell:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil, s == nil && o != nil:
			r = false
		default:
			x, y := s.stackCell, o.stackCell
			for r = true ; r && x != nil && y != nil; x, y = x.stackCell, y.stackCell {
				if v, ok := x.data.(Equatable); ok {
					r = v.Equal(y.data)
				} else if v, ok := y.data.(Equatable); ok {
					r = v.Equal(x.data)
				} else {
					r = x.data == y.data
				}
			}
			if r {
				r = x == nil && y == nil
			}
		}
	case nil:
		r = s == nil
	}
	return
}

/*
	Create a new cell containing the specified item and then append the stack to it.
	If the current cell is nil then the returned cell will be the terminal link of a new stack.
*/
func (s *stackCell) Push(item interface{}) (r *stackCell) {
	return &stackCell{ data: item, stackCell: s }
}

/*
	Return the data item stored in the top cell of the stack, or a nil if the stack is empty.
	If the current cell is nil then return nil.
*/
func (s *stackCell) Top() (r interface{}) {
	if s != nil {
		r = s.data
	}
	return
}

/*
	Return the data item stored in the top cell of the stack, or a nil if the stack is empty, along with a
	reference to the succeeding cell in the stack.
	If the current cell is nil then return nils.
	Beccause the cells are immutable this allows other stacks which reference the current stack cell to
	continue doing so.
*/
func (s *stackCell) Pop() (v interface{}, r *stackCell) {
	if s != nil {
		v = s.data
		r = s.stackCell
	}
	return
}

/*
	Iterate though the list of cells and calculate the depth of the stack.
	This routine is named Len() for interoperability with third-party packages.
*/
func (s *stackCell) Len() (r int) {
	if s != nil {
		if s.stackCell != nil {
			r = s.stackCell.Len()
		}
		r++
	}
	return
}

/*
	Make a copy of the data item stored in the current cell and then store this in a new cell which points
	to this cell.
	If the current cell is nil then return nil.
*/
func (s *stackCell) Dup() (r *stackCell) {
	if s != nil {
		r = s.Push(s.data)
	}
	return
}

/*
	Use the top two items on the stack to create a new list of cells in which their position is exchanged.
*/
func (s *stackCell) Swap() *stackCell {
	return s.Roll(1)
}

/*
	Make a new stack containing n cells where each cell contains the same value as is stored at the same depth
	in the existing stack.
*/
func (s *stackCell) Copy(n int) (r *stackCell) {
	r = new(stackCell)
	for x := r; n > 0 && s != nil; n-- {
		x.stackCell = stack(s.data)
		x = x.stackCell
		s = s.stackCell
	}
	return r.stackCell
}

/*
	Return the Nth cell from the top of the stack, or nil if there are fewer than N cells in the current stack.
*/
func (s *stackCell) Select(n int) (r *stackCell) {
	if s != nil {
		for r = s; n > 0 && r.stackCell != nil; r = r.stackCell { n-- }
		if n > 0 {
			r = nil
		}
	}
	return
}

/*
	Select the Nth cell from the top of the stack and if it is not nil, create a new cell with the same value and
	pointing to the top of the current stack.
*/
func (s *stackCell) Pick(n int) (r *stackCell) {
	if x := s.Select(n); x != nil {
		r = s.Push(x.data)
	} else {
		r = s
	}
	return
}

/*
	Create a new stack common with the current stack from the Nth+1 element. The Nth item of the current stack becames
	the first item of the new stack and then successive elements are filled with corresponding values starting with
	that at the top of the current stack.
*/
func (s *stackCell) Roll(n int) (r *stackCell) {
	switch {
	case n == 0:
		r = s
	case n > 0:
		if x := s.Select(n - 1); x == nil || x.stackCell == nil {
			r = s
		} else {
			r = &stackCell{ data: x.stackCell.data, stackCell: s }
			x.stackCell = x.stackCell.stackCell
		}
	}
	return
}