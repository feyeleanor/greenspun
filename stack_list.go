package greenspun

import "fmt"

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
		r = "nil:||"
	}
	return
}

func (s *StackList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *StackList:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil, s == nil && o != nil, s.depth != o.depth:
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
	if s != nil {
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
	if s != nil {
		_, s.stackCell = s.stackCell.Pop()
		s.depth--
	}
}

func (s *StackList) Swap() {
	if s != nil {
		s.stackCell = s.stackCell.Swap()
	}
}

func (s *StackList) Replace(item interface{}) {
	if s == nil {
		*s = StackList{}
	}
	s.stackCell.data = item
}