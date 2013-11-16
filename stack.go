package greenspun

import "fmt"

type stackCell struct {
	data	interface{}
	*stackCell	
}

func stack(items... interface{}) (r *stackCell) {
	for i := len(items); i > 0; {
		i--
		r = r.Push(items[i])
	}
	return
}

func (s *stackCell) String() (r string) {
	if s != nil {
		r = fmt.Sprintf("%v", s.data)
		for tos := s.stackCell; tos != nil; tos = tos.stackCell {
			r = fmt.Sprintf("%v %v", r, tos.data)
		}
	}
	return fmt.Sprintf("|%v|", r)
}

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
	}
	return
}

func (s *stackCell) Push(item interface{}) (r *stackCell) {
	return &stackCell{ data: item, stackCell: s }
}

func (s *stackCell) Top() (r interface{}) {
	if s != nil {
		r = s.data
	}
	return
}

func (s *stackCell) Pop() (v interface{}, r *stackCell) {
	if s != nil {
		v = s.data
		r = s.stackCell
	}
	return
}

func (s *stackCell) Len() (r int) {
	if s != nil {
		if s.stackCell != nil {
			r = s.stackCell.Len()
		}
		r++
	}
	return
}

func (s *stackCell) Dup() (r *stackCell) {
	if s != nil {
		r = s.Push(s.data)
	}
	return
}

func (s *stackCell) Swap() (r *stackCell) {
	if s != nil {
		if s.stackCell != nil {
			r = &stackCell{ data: s.stackCell.data, stackCell: &stackCell{ s.data, s.stackCell.stackCell }}
		} else {
			r = stack(s.data)
		}
	}
	return
}

func (s *stackCell) Copy(n int) (r *stackCell) {
	r = new(stackCell)
	for x := r; n > 0 && s != nil; n-- {
		x.stackCell = stack(s.data)
		x = x.stackCell
		s = s.stackCell
	}
	return r.stackCell
}

func (s *stackCell) end() (r *stackCell) {
	for r = s; r != nil && r.stackCell != nil; r = r.stackCell {}
	return
}

func (s *stackCell) pickCell(n int) (r *stackCell) {
	if s != nil {
		for r = s; n > 0 && r.stackCell != nil; r = r.stackCell { n-- }
		if n > 0 {
			r = nil
		}
	}
	return
}

func (s *stackCell) Pick(n int) (r *stackCell) {
	if x := s.pickCell(n); x != nil {
		r = s.Push(x.data)
	} else {
		r = s
	}
	return
}

func (s *stackCell) Roll(n int) (r *stackCell) {
	switch {
	case n == 0:
		r = s
	case n > 0:
		if x := s.pickCell(n - 1); x == nil || x.stackCell == nil {
			r = s
		} else {
			r = &stackCell{ data: x.stackCell.data, stackCell: s }
			x.stackCell = x.stackCell.stackCell
		}
	}
	return
}