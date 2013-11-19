package greenspun

import "fmt"

/*
	This is a wrapper for a cactus stack data structure as implemented by the stackCell type. It's inspired
	by Go's SliceHeader and StringHeader types.
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
		r = fmt.Sprintf("%v:%v", s.depth, s.stackCell)
	} else {
		r = "0:<]"
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

func (s *StackList) Push(item interface{}) (r *StackList) {
	if s == nil {
		r = Stack(item)
	} else {
		s.stackCell = s.stackCell.Push(item)
		s.depth++
		r = s
	}
	return
}

func (s *StackList) Top() (r interface{}) {
	if s != nil {
		r = s.stackCell.Top()
	}
	return
}

func (s *StackList) Pop() (r interface{}) {
	if s != nil && s.stackCell != nil {
		r, s.stackCell = s.stackCell.Pop()
		s.depth--
	}
	return
}

func (s *StackList) Len() (r int) {
	if s != nil {
	 	r = s.depth
	}
	return
}

func (s *StackList) Drop() {
	if s != nil && s.stackCell != nil {
		_, s.stackCell = s.stackCell.Pop()
		s.depth--
	}
}

func (s *StackList) Swap() {
	if s != nil {
		s.stackCell = s.stackCell.Swap()
	}
}

func (s *StackList) Roll(n int) {
	if s != nil {
		s.stackCell = s.stackCell.Roll(n)
	}
}

func (s *StackList) Replace(n int, item interface{}) {
	if s != nil && s.stackCell != nil {
		if cell := s.Select(n); cell != nil {
			cell.data = item
		}
	}
}