package greenspun

//import "fmt"

/*
	This is a wrapper for a spaghetti stack data structure as implemented by the stackCell type. It's
	inspired by Go's SliceHeader and StringHeader types.

	Whilst bare spaghetti stack structures are immutable containers, the Lifo implements Lisp's
	Rplaca and Rplacd functions and allows stack elements to be modified in situ.
*/

type Fifo struct {
	head		*stackCell
	tail		*stackCell
	length	int
}

func Queue(items... interface{}) (r *Fifo) {
	r = &Fifo{ length: len(items) }
	if r.length > 0 {
		r.head = &stackCell{ data: items[0] }
		r.tail = r.head
		for _, v := range items[1:] {
			r.tail.stackCell = &stackCell{ data: v }
			r.tail = r.tail.stackCell
		}
	}
	return
}

func (s *Fifo) String() (r string) {
	if s != nil {
		r = s.head.String()
	} else {
		r = "()"
	}
	return
}

func (s *Fifo) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Fifo:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil:
			r = s.head == nil
		case s == nil && o != nil:
			r = o.head == nil
		case s.length != o.length:
			r = false
		default:
			x, y := s.head, o.head
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
			r = s.head.Equal(o)
		} else {
			r = o == nil
		}
	case nil:
		r = s == nil || s.head == nil
	}
	return
}

func (s *Fifo) Put(item interface{}) {
	if s == nil {
		panic(STACK_UNINITIALIZED)
	}
	s.tail.stackCell = &stackCell{ data: item }
	s.tail= s.tail.stackCell
	s.length++
}