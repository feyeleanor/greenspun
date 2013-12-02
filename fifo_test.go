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

func TestFifoPeek(t *testing.T) {
	RefutePeek := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Peek() should panic", vs)()
		s.Peek()
	}

	ConfirmPeek := func(s *Fifo, r interface{}) {
		if x := s.Peek(); x != r {
			t.Fatalf("%v.Peek() should be %v but is %v", s, r, x)
		}
	}

	RefutePeek(nil)
	RefutePeek(Queue())
	ConfirmPeek(Queue(0), 0)
	ConfirmPeek(Queue(1, 0), 1)
	ConfirmPeek(Queue(2, 1, 0), 2)
}

func TestFifoPop(t *testing.T) {
	RefutePop := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pop() should panic", vs)()
		s.Pop()
	}

	ConfirmPop := func(s *Fifo, r interface{}, n *Fifo) {
		vs := s.String()
		switch x := s.Pop(); {
		case x != r:
			t.Fatalf("%v.Pop() should be %v but is %v", vs, r, x)
		case !s.Equal(n):
			t.Fatalf("%v.Pop() should leave %v but leaves %v", vs, n, s)
		}
	}

	RefutePop(nil)
	RefutePop(Queue())
	ConfirmPop(Queue(0), 0, Queue())
	ConfirmPop(Queue(1, 0), 1, Queue(0))
	ConfirmPop(Queue(2, 1, 0), 2, Queue(1, 0))
}

func TestFifoLen(t * testing.T) {
	ConfirmLen := func(s *Fifo, r int) {
		if l := s.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(Queue(), 0)
	ConfirmLen(Queue(0), 1)
	ConfirmLen(Queue(1, 0), 2)
}

func TestFifoDrop(t *testing.T) {
	RefuteDrop := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Drop() should panic", vs)()
		s.Drop()
	}

	ConfirmDrop := func(s, r *Fifo) {
		vs := s.String()
		if s.Drop(); !s.Equal(r) {
			t.Fatalf("%v.Drop() should leave %v but leaves %v", vs, r, s)
		}
	}

	RefuteDrop(nil)
	ConfirmDrop(Queue(), nil)
	ConfirmDrop(Queue(), Queue())
	ConfirmDrop(Queue(0), Queue())
	ConfirmDrop(Queue(1, 0), Queue(0))
	ConfirmDrop(Queue(2, 1, 0), Queue(1, 0))
}

func TestFifoDup(t *testing.T) {
	RefuteDup := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Dup() should panic", vs)()
		s.Dup()
	}

	ConfirmDup := func(s, r *Fifo) {
		vs := s.String()
		if s.Dup(); !s.Equal(r) {
			t.Fatalf("%v.Dup() should be %v but is %v", vs, r, s)
		}
	}

	RefuteDup(nil)
	RefuteDup(Queue())
	ConfirmDup(Queue(1), Queue(1, 1))
	ConfirmDup(Queue(1, 2), Queue(1, 2, 1))
}

func TestFifoSwap(t *testing.T) {
	RefuteSwap := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Swap() should panic", vs)()
		s.Swap()
	}

	ConfirmSwap := func(s, r *Fifo) {
		vs := s.String()
		if s.Swap(); !s.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", vs, r, s)
		}
	}

	RefuteSwap(nil)
	RefuteSwap(Queue())
	RefuteSwap(Queue(0))
	ConfirmSwap(Queue(1, 0), Queue(0, 1))
	ConfirmSwap(Queue(2, 1, 0), Queue(0, 1, 2))
	ConfirmSwap(Queue(3, 2, 1, 0), Queue(0, 2, 1, 3))
}

func TestFifoCopy(t *testing.T) {
	ConfirmCopy := func(s *Fifo, n int, r *Fifo) {
		if x := s.Copy(n); !x.Equal(r) {
			t.Fatalf("%v.Copy(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmCopy(nil, 0, nil)
	ConfirmCopy(nil, 1, nil)

	ConfirmCopy(Queue(), 0, Queue())
	ConfirmCopy(Queue(), 1, Queue())

	ConfirmCopy(Queue(0), 0, Queue())
	ConfirmCopy(Queue(0), 1, Queue(0))
	ConfirmCopy(Queue(0), 2, Queue(0))

	ConfirmCopy(Queue(0, 1), 0, Queue())
	ConfirmCopy(Queue(0, 1), 1, Queue(0))
	ConfirmCopy(Queue(0, 1), 2, Queue(0, 1))
	ConfirmCopy(Queue(0, 1), 3, Queue(0, 1))

	ConfirmCopy(Queue(0, 1, 2), 0, Queue())
	ConfirmCopy(Queue(0, 1, 2), 1, Queue(0))
	ConfirmCopy(Queue(0, 1, 2), 2, Queue(0, 1))
	ConfirmCopy(Queue(0, 1, 2), 3, Queue(0, 1, 2))
	ConfirmCopy(Queue(0, 1, 2), 4, Queue(0, 1, 2))
}

func TestFifoMove(t *testing.T) {
	RefuteMove := func(s *Fifo, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Move(%v) should panic", vs, x)()
		s.Move(x)
	}

	ConfirmMove := func(s *Fifo, n int, r *Fifo) {
		vs := s.String()
		if s.Move(n); !s.Equal(r) {
			t.Fatalf("%v.Move(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	RefuteMove(nil, 0)
	RefuteMove(nil, 1)

	RefuteMove(Queue(), 0)
	RefuteMove(Queue(), 1)

	ConfirmMove(Queue(0), 0, Queue(0))
	RefuteMove(Queue(0), 1)

	ConfirmMove(Queue(0, 1), 0, Queue(0, 1))
	ConfirmMove(Queue(0, 1), 1, Queue(1))
	RefuteMove(Queue(0, 1), 2)
}

func TestFifoPick(t *testing.T) {
	RefutePick := func(s *Fifo, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pick(%v) should panic", vs, x)()
		s.Pick(x)
	}

	ConfirmPick := func(s *Fifo, n int, r *Fifo) {
		vs := s.String()
		if s.Pick(n); !s.Equal(r) {
			t.Fatalf("%v.Pick(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	RefutePick(nil, 0)
	RefutePick(nil, 1)

	RefutePick(Queue(), 0)
	RefutePick(Queue(), 1)

	ConfirmPick(Queue(0), 0, Queue(0, 0))
	RefutePick(Queue(0), 1)

	ConfirmPick(Queue(0, 1), 0, Queue(0, 1, 0))
	ConfirmPick(Queue(0, 1), 1, Queue(0, 1, 1))
	RefutePick(Queue(0, 1), 2)

	ConfirmPick(Queue(0, 1, 2), 0, Queue(0, 1, 2, 0))
	ConfirmPick(Queue(0, 1, 2), 1, Queue(0, 1, 2, 1))
	ConfirmPick(Queue(0, 1, 2), 2, Queue(0, 1, 2, 2))
	RefutePick(Queue(0, 1, 2), 3)
}