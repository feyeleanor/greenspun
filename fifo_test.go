package greenspun

import (
//	"fmt"
	"testing"
)

func TestFifoString(t *testing.T) {
	ConfirmString := func(s *Fifo, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "()")
	ConfirmString(new(Fifo), "()")
	ConfirmString(&Fifo{ length: 1, head: stack(1) }, "(1)")
	ConfirmString(&Fifo{ length: 2, head: stack(1, 2) }, "(1 2)")
	ConfirmString(&Fifo{ length: 3, head: stack(1, 2, 3) }, "(1 2 3)")
	ConfirmString(Queue(1), "(1)")
	ConfirmString(Queue(1, 2), "(1 2)")
	ConfirmString(Queue(1, 2, 3), "(1 2 3)")
}

func TestFifoEqual(t *testing.T) {
	ConfirmEqual := func(x *Fifo, y Equatable, r bool) {
		if z := x.Equal(y); z != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", x, y, r, z)
		}
	}

	ConfirmEqual(nil, (*Fifo)(nil), true)
	ConfirmEqual(Queue(1), (*Fifo)(nil), false)
	ConfirmEqual(nil, Queue(), true)
	ConfirmEqual(Queue(), nil, true)
	ConfirmEqual(nil, Queue(1), false)
	ConfirmEqual(Queue(1), Queue(1), true)

	ConfirmEqual(Queue(1, 2), (*Fifo)(nil), false)
	ConfirmEqual(nil, Queue(1, 2), false)
	ConfirmEqual(Queue(1, 2), Queue(1, 2), true)
	ConfirmEqual(Queue(1, 2), Queue(2, 1), false)
	ConfirmEqual(Queue(2, 1), Queue(1, 2), false)
	ConfirmEqual(Queue(2, 1), Queue(2, 1), true)

	ConfirmEqual(Queue(1, 2, 3), Queue(1, 2, 3), true)
	ConfirmEqual(Queue(2, 1, 3), Queue(1, 2, 3), false)
	ConfirmEqual(Queue(3, 2, 1), Queue(3, 2, 1), true)

	ConfirmEqual(nil, stack(1), false)
	ConfirmEqual(Queue(1), stack(1), true)

	ConfirmEqual(Queue(1, 2), (*stackCell)(nil), false)
	ConfirmEqual(nil, stack(1, 2), false)
	ConfirmEqual(Queue(1, 2), stack(1, 2), true)
	ConfirmEqual(Queue(1, 2), stack(2, 1), false)
	ConfirmEqual(Queue(2, 1), stack(1, 2), false)
	ConfirmEqual(Queue(2, 1), stack(2, 1), true)

	ConfirmEqual(Queue(1, 2, 3), stack(1, 2, 3), true)
	ConfirmEqual(Queue(2, 1, 3), stack(1, 2, 3), false)
	ConfirmEqual(Queue(3, 2, 1), stack(3, 2, 1), true)
}

func TestFifoPut(t *testing.T) {
	RefutePut := func(s *Fifo, v interface{}) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Put(%v) should panic", vs, v)()
		s.Put(v)
	}

	ConfirmPut := func(s *Fifo, v interface{}, r *Fifo) {
		vs := s.String()
		if s.Put(v); !s.Equal(r) {
			t.Fatalf("%v.Put(%v) should be %v but is %v", vs, v, r, s)
		}
	}

	RefutePut(nil, nil)
	RefutePut(nil, 1)
	ConfirmPut(Queue(1), 1, Queue(1, 1))
	ConfirmPut(Queue(1, 2), 1, Queue(1, 2, 1))
	ConfirmPut(Queue(1, 2, 3), 1, Queue(1, 2, 3, 1))
}