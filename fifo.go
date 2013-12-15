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
	head		*stackCell
	tail		*stackCell
	length	int
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

func (s *Fifo) copyHeader() *Fifo {
	return &Fifo{ head: s.head, tail: s.tail, length: s.length }
}

func (s *Fifo) reverseTail() (r *Fifo) {
	if s == nil {
		panic(LIST_UNINITIALIZED)
	}

	if s.head == nil {
		r = &Fifo{ length: s.length }
		s.tail.Each(func(v interface{}) {
			r.head = r.head.Push(v)
		})
	} else {
		r = s.copyHeader()
	}
	return
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
				r = shead.MatchValue(ohead)
			}
			for ; r && stail != nil && otail != nil; stail , otail = stail.stackCell, otail.stackCell {
				r = stail.MatchValue(otail)
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
				r = shead.MatchValue(o)
			}
			if r {
				r = o.Equal(s.tail.Reverse())
			}
		}
	case nil:
		r = s == nil || (s.head == nil && s.tail == nil)
	}
	return
}

func (s *Fifo) Put(item interface{}) (r *Fifo) {
	if s == nil {
		r = new(Fifo)
	} else {
		r = s.copyHeader()
	}
	r.tail.Push(item)
	r.length++
	return
}

func (s *Fifo) Peek() (v interface{}) {
	if r := s.reverseTail(); r.length > 0 {
		*s = *r
		v = s.head.Peek()
	}
	return
}

func (s *Fifo) Pop() (v interface{}, r *Fifo) {
	if r = s.reverseTail(); r.length == 0 {
		panic(LIST_EMPTY)
	}
	v, r.head = r.head.Pop()
	r.length--
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
	if r = s.reverseTail(); r.length > 0 {
		r.head = r.head.stackCell
		r.length--
	}
	return
}

func (s *Fifo) Dup() (r *Fifo) {
	if r = s.reverseTail(); r.length == 0 {
		panic(LIST_TOO_SHALLOW)
	}
	r.tail.Push(r.head.Peek())
	return
}

func (s *Fifo) Swap() (r *Fifo) {
	switch {
	case s == nil:
		panic(LIST_UNINITIALIZED)
	case s.length == 1:
		panic(LIST_TOO_SHALLOW)
	case s.head == nil:
		r = &Fifo{ tail: s.tail.stackCell, length: s.length }
		r = r.reverseTail()
		r.tail = &stackCell{ data: s.tail.data }
		r.head.data, r.tail.data = r.tail.data, r.head.data
	case s.tail == nil:
		r = &Fifo{ head: &stackCell{ data: s.head.data }, length: 1 }
		for c := s.head.stackCell; c != nil; c = c.stackCell {
			r.Put(c.data)
		}
		r.head.data, r.tail.data = r.tail.data, r.head.data
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
			r.Put(v)
		}
	}
 	return
}

/*
	Make a new queue in which the elements in the source queue are reversed.
*/
func (s *Fifo) Reverse() (r *Fifo) {
	switch {
	case s == nil:
		panic(LIST_UNINITIALIZED)
	case s.tail == nil:
		r = s.copyHeader()
		r.head, r.tail = r.tail, r.head
		r.reverseTail()
		r.head, r.tail = r.tail, r.head
	case s.head == nil:
		r = s.reverseTail()
	default:
		r.head, r.tail = r.tail, r.head
	}
	return
}

/*
	Move to the Nth cell from the start of the queue, or return an error if there are fewer than N cells.
*/
func (s *Fifo) Move(n int) (r *Fifo) {
	switch {
	case s == nil:
		panic(LIST_UNINITIALIZED)
	case n > s.length:
		panic(LIST_TOO_SHALLOW)
	}
	for r = s.copyHeader(); n > 0; n-- {
		r = r.Drop()
	}
	return
}

/*
	Move to the Nth cell from the front of the queue and create a new cell with the same value which is appended to the queue.
*/
func (s *Fifo) Pick(n int) (r *Fifo) {
	switch {
	case s == nil:
		panic(LIST_UNINITIALIZED)
	case n > s.length:
		panic(LIST_TOO_SHALLOW)
	}
	r = s.Move(n)
	return s.Put(r.Peek())
}

/*
	Create a new queue common with the current queue from the Nth+1 element. The Nth item of the current queue becames
	the first item of the new queue and then successive elements are filled with corresponding values starting with
	that at the front of the current queue.
*/
func (s *Fifo) Roll(n int) (r *Fifo) {




	return
}