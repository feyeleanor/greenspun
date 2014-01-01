package greenspun

import (
	"fmt"
	"strings"
)

/*
	This is a wrapper for a functional queue data structure implemented as a pair of stackCells. It's
	inspired by Chris Okasaki's non-lazy functional queue and Go's SliceHeader and StringHeader types.

	When an item is appended to the queue, it's pushed onto the tail stack.
	When an item is popped from the queue, it's popped from the head stack.
	When the head stack is empty, the tail stack is reversed and assigned as the head stack.
	We also maintain a cached length for the queue which is incremented on append & decremented on pop.
*/


type Fifo struct {
	head			*stackCell
	tail			*stackCell
	length		int
	Mutex
}

func Queue(items... interface{}) (r *Fifo) {
	r = &Fifo{ length: len(items) }
	if r.length > 0 {
		r.head = &stackCell{ data: items[0] }
		for _, v := range items[1:] {
			r.tail = &stackCell{ data: v, stackCell: r.tail }
		}
	}
	return
}

func (s *Fifo) copyHeader() (r *Fifo) {
	if s != nil {
		r = &Fifo{ head: s.head, tail: s.tail, length: s.length }
	}
	return
}

/*
	A functional queue contains two stacks representing the front and back of the queue.
	reverseTail() is an optimisation used when the front stack is empty to reverse the back
	stack elements and create a new front stack.

	This is a destructive operation and uses the header mutex. This is an optimisation to
	minimise the number of reversal operations.

	Normally when a reversal occurs it's associated with an operation which will create a new
	header, therefore we return a copy of the modified queue header as a convenience.
*/
func (s *Fifo) reverseTail() *Fifo {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}

	if s.head == nil {
		n := new(Fifo)
		s.tail.Each(func(v interface{}) {
			n.head = n.head.Push(v)
		})
		s.CriticalSection(func() {
			s.head = n.head
			s.tail = nil
		})
	}
	return s.copyHeader()
}

func (s *Fifo) String() (r string) {
	if s != nil {
		l := s.length
		b := make([]string, l, l)
		s.head.Each(func(i int, v interface{}) {
			b[i] = fmt.Sprintf("%v", v)
		})
		s.tail.Each(func(v interface{}) {
			l--
			b[l] = fmt.Sprintf("%v", v)
		})
		r = strings.Join(b, " ")
	}
	return "(" + r + ")"
}

func (s *Fifo) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Fifo:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil:
			r = s.head == nil && s.tail == nil
		case s == nil && o != nil:
			r = o.head == nil && o.tail == nil
		case s.length != o.length:
			r = false
		default:
			shead, ohead, stail, otail := s.head, o.head, s.tail, o.tail
			for r = true; r && shead != nil && ohead != nil; shead , ohead = shead.stackCell, ohead.stackCell {
				r = MatchValue(shead, ohead)
			}
			for ; r && stail != nil && otail != nil; stail , otail = stail.stackCell, otail.stackCell {
				r = MatchValue(stail, otail)
			}
			if r {
				switch {
				case shead != nil && otail != nil, ohead != nil && stail != nil:
					r = shead.Equal(otail.Reverse())
				}
			}
		}
	case *stackCell:
		switch {
		case s == nil && o == nil:
			r = true
		case s == nil && o != nil:
			r = false
		case s != nil && o == nil:
			r = s.head == nil && s.tail == nil
		default:
			shead := s.head
			for r = true; r && shead != nil && o != nil; shead, o = shead.stackCell, o.stackCell {
				r = MatchValue(shead, o)
			}
			if r {
				r = o.Equal(s.tail.Reverse())
			}
		}
	case Sequence:
		switch {
		case s == nil && o == nil:
			r = true
		case s == nil && o != nil:
			r = false
		case s != nil && o == nil:
			r = s.head == nil && s.tail == nil
		default:
			shead := s.head
			for r = true; r && shead != nil && o != nil; shead, o = shead.stackCell, o.Next() {
				r = shead.data == o.Peek()
			}
			if r {
				r = s.tail.Reverse().Equal(o)
			}
		}
	case nil:
		r = s == nil || (s.head == nil && s.tail == nil)
	}
	return
}

func (s *Fifo) Append(item interface{}) (r *Fifo) {
	if s == nil {
		r = new(Fifo)
	} else {
		r = s.copyHeader()
	}
	r.tail = r.tail.Push(item)
	r.length++
	return
}

func (s *Fifo) Peek() (v interface{}) {
	if s.length == 0 {
		panic(LIST_EMPTY)
	}
	s.reverseTail()
	return s.head.data
}

func (s *Fifo) Pop() (v interface{}, r *Fifo) {
	if s.length == 0 {
		panic(LIST_EMPTY)
	}
	s.reverseTail()
	r = s.copyHeader()
	r.length--
	v, r.head = r.head.Pop()
	return
}

func (s *Fifo) Len() (r int) {
	if s != nil {
	 	r = s.length
	}
	return
}

func (s *Fifo) IsNil() (r bool) {
	return s == nil
}

func (s *Fifo) Drop() (r *Fifo) {
	if s != nil && s.length > 0 {
		s = s.reverseTail()
		if s.head != nil && s.length > 0 {
			r = &Fifo{ head: s.head.stackCell, tail: s.tail, length: s.length - 1 }
		}
	}
	return
}

func (s *Fifo) Dup() *Fifo {
	if s.length == 0 {
		panic(LIST_TOO_SHALLOW)
	}
	s.reverseTail()
	return s.Append(s.Peek())
}

func (s *Fifo) Swap() (r *Fifo) {
	switch {
	case s == nil:
		r = nil
	case s.length < 2:
		r = s
	case s.head == nil && s.tail == nil:
		r = new(Fifo)
	case s.head == nil:
		v, t := s.tail.Pop()
		r = &Fifo{ tail: t, length: s.length - 1 }
		r.reverseTail()
		r.tail = &stackCell{ data: r.head.data }
		r.length++
		r.head = &stackCell{ data: v, stackCell: r.head.stackCell }
	case s.tail == nil:
		v, h := s.head.Pop()
		t := h.Reverse()
		r = &Fifo{ head: &stackCell{ data: t.data }, tail: &stackCell{ data: v, stackCell: t.stackCell }, length: s.length }
		
	default:
		r = &Fifo{	head: s.head.stackCell.Push(s.tail.data),
								tail: s.tail.stackCell.Push(s.head.data),
								length: s.length,
							}
	}
	return
}

/*
	Make a new queue containing n cells where each cell contains the same value as is stored at the same depth
	in the existing queue.

	If the queue is shorter than n we copy the entire queue.
*/
func (s *Fifo) Copy(n int) (r *Fifo) {
	if s != nil {
		if n > s.length {
			n = s.length
		}
		r = new(Fifo)
		var v	interface{}
		for ; n > 0; n-- {
			v, s = s.Pop()
			r = r.Append(v)
		}
	}
 	return
}

/*
	Make a new queue in which the elements in the source queue are reversed.
*/
func (s *Fifo) Reverse() *Fifo {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}
	return &Fifo{ head: s.tail.Clone(), tail: s.head.Clone(), length: s.length }
}

/*
	Move to the Nth cell from the start of the queue, or return nil if there are fewer than N cells.
*/
func (s *Fifo) Move(n int) (r *Fifo) {
	switch {
	case s == nil, n >= s.length:
		r = nil
	default:
		for r = s.copyHeader(); n > 0; n-- {
			r = r.Drop()
		}
	}
	return
}

/*
	Move to the Nth cell from the front of the queue and create a new cell with the same value which is appended to the queue.
*/
func (s *Fifo) Pick(n int) (r *Fifo) {
	switch {
	case s == nil, n >= s.length:
		r = s
	default:
		r = s.Move(n)
		r = s.Append(r.Peek())
	}
	return
}

/*
	Move to the Nth cell from the front of the queue and copy its value, then make a new queue in which this is the first element
	and the succeeding elements are the section 0..N-1 followed by N+1 onwards.
*/
func (s *Fifo) Roll(n int) (r *Fifo) {
	switch {
	case s == nil, n >= s.length, n == 0:
		r = s
	default:
		var v	interface{}

		for ; n > 0; n-- {
			v, s = s.Pop()
			r = r.Append(v)
		}
		r = r.reverseTail()

		v, s = s.Pop()
		r = &Fifo{ head: &stackCell{ data: v, stackCell: r.head }, tail: r.tail, length: r.length + 1 }
		for ; s.head != nil || s.tail != nil; {
			v, s = s.Pop()
			r = r.Append(v)
		}
	}
 	return
}