package greenspun

import "testing"

func TestFifoQueue(t *testing.T) {
	ConfirmQueue := func(r *Fifo, s ...interface{}) {
		if q := Queue(s...); !q.Equal(r) {
			t.Fatalf("Queue(%v) should be %v but is %v", s, r, q)
		}
	}

	ConfirmQueue(new(Fifo))
	ConfirmQueue(&Fifo{ head: stack(1), tail: nil, length: 1 }, 1)
	ConfirmQueue(&Fifo{ head: stack(1, 2), tail: nil, length: 2 }, 1, 2)
	ConfirmQueue(&Fifo{ head: stack(1, 2, 3), tail: nil, length: 3 }, 1, 2, 3)
	ConfirmQueue(&Fifo{ head: stack(1), tail: stack(2), length: 2 }, 1, 2)
	ConfirmQueue(&Fifo{ head: stack(1, 2), tail: stack(3), length: 3 }, 1, 2, 3)
}

func TestCopyHeader(t *testing.T) {
	if x := (*Fifo)(nil).copyHeader(); x != nil {
		t.Fatalf("a nil header should produce a nil header")
	}

	ConfirmCopyHeader := func(s *Fifo) {
		switch x := s.copyHeader(); {
		case !x.head.Equal(s.head):
			t.Fatalf("%v.head != %v", s, x.head)
		case !x.tail.Equal(s.tail):
			t.Fatalf("%v.tail != %v", s, x.tail)
		case x.length != s.length:
			t.Fatalf("%v.length != %v", s, x.length)
		case !x.Equal(s):
			t.Fatalf("%v.copyHeader() != %v", s, x)
		}
	}

	ConfirmCopyHeader(Queue())

	ConfirmCopyHeader(Queue(0))
	ConfirmCopyHeader(&Fifo{ head: stack(0), length: 1 })
	ConfirmCopyHeader(&Fifo{ tail: stack(0), length: 1 })

	ConfirmCopyHeader(Queue(0, 1))
	ConfirmCopyHeader(&Fifo{ head: stack(0, 1), length: 2 })
	ConfirmCopyHeader(&Fifo{ head: stack(0), tail: stack(1), length: 2 })
	ConfirmCopyHeader(&Fifo{ tail: stack(1, 0), length: 2 })

	ConfirmCopyHeader(Queue(0, 1, 2))
	ConfirmCopyHeader(&Fifo{ head: stack(0, 1, 2), length: 3 })
	ConfirmCopyHeader(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 })
	ConfirmCopyHeader(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 })
	ConfirmCopyHeader(&Fifo{ tail: stack(2, 1, 0), length: 3 })
}

func TestFifoReverseTail(t *testing.T) {
	ConfirmReverseTail := func(f, r *Fifo) {
		fs := f.String()
		if f.reverseTail(); !f.Equal(r) {
			t.Fatalf("%v.reverseTail() should be %v but is %v", fs, r, f)
		}
	}

	ConfirmReverseTail(&Fifo{}, &Fifo{})
	ConfirmReverseTail(&Fifo{ head: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 })
	ConfirmReverseTail(&Fifo{ tail: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 })
	ConfirmReverseTail(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ head: stack(0), tail: stack(1), length: 2 })
	ConfirmReverseTail(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, &Fifo{ head: stack(0, 1), tail: stack(2), length: 3 })
	ConfirmReverseTail(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, &Fifo{ head: stack(0), tail: stack(2, 1), length: 3 })
}

func TestFifoBalance(t *testing.T) {
	ConfirmBalance := func(f, r *Fifo) {
		switch x := f.balance(); {
		case x.length != r.length:
			t.Fatalf("%v.balance() should be %v long but is %v long", f, r.length, x.length)
		case !x.head.Equal(r.head):
			t.Fatalf("%v.balance() head should be %v but is %v", f, r.head, x.head)
		case !x.tail.Equal(r.tail):
			t.Fatalf("%v.balance() tail should be %v but is %v", f, r.tail, x.tail)
		}
	}

	ConfirmBalance(&Fifo{}, &Fifo{})

	ConfirmBalance(&Fifo{ head: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 })
	ConfirmBalance(&Fifo{ tail: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 })

	ConfirmBalance(&Fifo{ head: stack(0, 1), length: 2 }, &Fifo{ head: stack(0), tail: stack(1), length: 2 })
	ConfirmBalance(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ head: stack(0), tail: stack(1), length: 2 })
	ConfirmBalance(&Fifo{ tail: stack(1, 0), length: 2 }, &Fifo{ head: stack(0), tail: stack(1), length: 2 })

	ConfirmBalance(&Fifo{ head: stack(0, 1, 2), length: 3 }, &Fifo{ head: stack(0), tail: stack(2, 1), length: 3 })
	ConfirmBalance(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, &Fifo{ head: stack(0, 1), tail: stack(2), length: 3 })
	ConfirmBalance(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, &Fifo{ head: stack(0), tail: stack(2, 1), length: 3 })
	ConfirmBalance(&Fifo{ tail: stack(2, 1, 0), length: 3 }, &Fifo{ head: stack(0, 1), tail: stack(2), length: 3 })

	ConfirmBalance(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, &Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 })
	ConfirmBalance(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, &Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 })
	ConfirmBalance(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, &Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 })
	ConfirmBalance(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, &Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 })

	ConfirmBalance(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, &Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 })
	ConfirmBalance(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, &Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 })
	ConfirmBalance(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, &Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 })
	ConfirmBalance(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, &Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 })
	ConfirmBalance(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, &Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 })
}

func TestFifoString(t *testing.T) {
	ConfirmString := func(s *Fifo, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "()")
	ConfirmString(new(Fifo), "()")

	ConfirmString(Queue(1), "(1)")
	ConfirmString(&Fifo{ length: 1, head: stack(1) }, "(1)")
	ConfirmString(&Fifo{ length: 1, tail: stack(1) }, "(1)")

	ConfirmString(Queue(1, 2), "(1 2)")
	ConfirmString(&Fifo{ length: 2, head: stack(1, 2) }, "(1 2)")
	ConfirmString(&Fifo{ length: 2, head: stack(1), tail: stack(2) }, "(1 2)")
	ConfirmString(&Fifo{ length: 2, tail: stack(2, 1) }, "(1 2)")

	ConfirmString(Queue(1, 2, 3), "(1 2 3)")
	ConfirmString(&Fifo{ length: 3, tail: stack(3, 2, 1) }, "(1 2 3)")
	ConfirmString(&Fifo{ length: 3, head: stack(1), tail: stack(3, 2) }, "(1 2 3)")
	ConfirmString(&Fifo{ length: 3, head: stack(1, 2), tail: stack(3) }, "(1 2 3)")
	ConfirmString(&Fifo{ length: 3, head: stack(1, 2, 3) }, "(1 2 3)")
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

	ConfirmEqual(&Fifo{ head: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 }, true)
	ConfirmEqual(&Fifo{ head: stack(0), length: 1 }, Queue(0), true)
	ConfirmEqual(&Fifo{ tail: stack(0), length: 1 }, &Fifo{ head: stack(0), length: 1 }, true)
	ConfirmEqual(&Fifo{ tail: stack(0), length: 1 }, Queue(0), true)

	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ head: stack(0, 1), length: 2 }, true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ head: stack(1, 0), length: 2 }, false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, Queue(0, 1), true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, Queue(1, 0), false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ tail: stack(1, 0), length: 2 }, true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, &Fifo{ tail: stack(0, 1), length: 2 }, false)

	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, &Fifo{ head: stack(0, 1, 2), length: 3 }, true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, Queue(0, 1, 2), true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, &Fifo{ tail: stack(2, 1, 0), length: 3 }, true)

	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, &Fifo{ head: stack(0, 1, 2, 3), length: 4 }, true)
	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, Queue(0, 1, 2, 3), true)
	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, &Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, true)

	ConfirmEqual(&Fifo{ head: stack(0), length: 1 }, stack(), false)
	ConfirmEqual(&Fifo{ head: stack(0), length: 1 }, stack(0), true)
	ConfirmEqual(&Fifo{ head: stack(0), length: 1 }, stack(0, 1), false)

	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, stack(0), false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, stack(0, 1), true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, stack(1, 0), false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, stack(0, 1, 2), false)

	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, stack(0, 1), false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, stack(0, 1, 2), true)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, stack(1, 0, 2), false)
	ConfirmEqual(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, stack(0, 1, 2, 3), false)

	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, stack(0, 1, 2), false)
	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, stack(0, 1, 2, 3), true)
	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, stack(0, 2, 1, 3), false)
	ConfirmEqual(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, stack(0, 1, 2, 3, 4), false)
}

func TestFifoAppend(t *testing.T) {
	ConfirmAppend := func(s *Fifo, v interface{}, r *Fifo) {
		if x := s.Append(v); !x.Equal(r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmAppend(nil, 1, Queue(1))
	ConfirmAppend(Queue(1), 1, Queue(1, 1))
	ConfirmAppend(&Fifo{ head: stack(1), length: 1 }, 1, Queue(1, 1))
	ConfirmAppend(&Fifo{ tail: stack(1), length: 1 }, 1, Queue(1, 1))

	ConfirmAppend(Queue(1, 2), 1, Queue(1, 2, 1))
	ConfirmAppend(&Fifo{ head: stack(1, 2), length: 2 }, 1, Queue(1, 2, 1))
	ConfirmAppend(&Fifo{ head: stack(1), tail: stack(2), length: 2 }, 1, Queue(1, 2, 1))
	ConfirmAppend(&Fifo{ tail: stack(2, 1), length: 2 }, 1, Queue(1, 2, 1))

	ConfirmAppend(Queue(1, 2, 3), 1, Queue(1, 2, 3, 1))
	ConfirmAppend(&Fifo{ head: stack(1, 2, 3), length: 3 }, 1, Queue(1, 2, 3, 1))
	ConfirmAppend(&Fifo{ head: stack(1, 2), tail: stack(3), length: 3 }, 1, Queue(1, 2, 3, 1))
	ConfirmAppend(&Fifo{ head: stack(1), tail: stack(3, 2), length: 3 }, 1, Queue(1, 2, 3, 1))
	ConfirmAppend(&Fifo{ tail: stack(3, 2, 1), length: 3 }, 1, Queue(1, 2, 3, 1))
}

func TestFifoPeek(t *testing.T) {
	RefutePeek := func(s *Fifo) {
		defer ConfirmPanic(t, "%v.Peek() should panic", s)()
		s.Peek()
	}

	ConfirmPeek := func(s *Fifo, r interface{}) {
		if x := s.Peek(); x != r {
			t.Fatalf("%v.Peek() should be %v but is %v", s, r, x)
		}
	}

	RefutePeek(nil)
	RefutePeek(Queue())

	ConfirmPeek(Queue(1), 1)
	ConfirmPeek(&Fifo{ head: stack(1), length: 1 }, 1)
	ConfirmPeek(&Fifo{ tail: stack(1), length: 1 }, 1)

	ConfirmPeek(Queue(1, 2), 1)
	ConfirmPeek(&Fifo{ head: stack(1, 2), length: 2 }, 1)
	ConfirmPeek(&Fifo{ head: stack(1), tail: stack(2), length: 2 }, 1)
	ConfirmPeek(&Fifo{ tail: stack(2, 1), length: 2 }, 1)

	ConfirmPeek(Queue(1, 2, 3), 1)
	ConfirmPeek(&Fifo{ head: stack(1, 2, 3), length: 3 }, 1)
	ConfirmPeek(&Fifo{ head: stack(1, 2), tail: stack(3), length: 3 }, 1)
	ConfirmPeek(&Fifo{ head: stack(1), tail: stack(3, 2), length: 3 }, 1)
	ConfirmPeek(&Fifo{ tail: stack(3, 2, 1), length: 3 }, 1)
}

func TestFifoPop(t *testing.T) {
	RefutePop := func(s *Fifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pop() should panic", vs)()
		s.Pop()
	}

	ConfirmPop := func(s *Fifo, r interface{}, n *Fifo) {
		switch v, x := s.Pop(); {
		case v != r:
			t.Fatalf("%v.Pop() should be %v but is %v", s, r, v)
		case !x.Equal(n):
			t.Fatalf("%v.Pop() should leave %v but leaves %v", s, n, x)
		}
	}

	RefutePop(nil)
	RefutePop(Queue())
	ConfirmPop(Queue(0), 0, Queue())
	ConfirmPop(&Fifo{ head: stack(0), length: 1 }, 0, Queue())
	ConfirmPop(&Fifo{ tail: stack(0), length: 1 }, 0, Queue())

	ConfirmPop(Queue(0, 1), 0, Queue(1))
	ConfirmPop(&Fifo{ head: stack(0, 1), length: 2 }, 0, Queue(1))
	ConfirmPop(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 0, Queue(1))
	ConfirmPop(&Fifo{ tail: stack(1, 0), length: 2 }, 0, Queue(1))

	ConfirmPop(Queue(0, 1, 2), 0, Queue(1, 2))
	ConfirmPop(&Fifo{ head: stack(0, 1, 2), length: 3 }, 0, Queue(1, 2))
	ConfirmPop(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 0, Queue(1, 2))
	ConfirmPop(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 0, Queue(1, 2))
	ConfirmPop(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 0, Queue(1, 2))

	ConfirmPop(Queue(0, 1, 2, 3), 0, Queue(1, 2, 3))
	ConfirmPop(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 0, Queue(1, 2, 3))
	ConfirmPop(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 0, Queue(1, 2, 3))
	ConfirmPop(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 0, Queue(1, 2, 3))
	ConfirmPop(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 0, Queue(1, 2, 3))
	ConfirmPop(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 0, Queue(1, 2, 3))
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
	ConfirmDrop := func(s, r *Fifo) {
		if x := s.Drop(); !x.Equal(r) {
			t.Fatalf("%v.Drop() should leave %v but leaves %v", s, r, x)
		}
	}

	ConfirmDrop(Queue(), nil)
	ConfirmDrop(Queue(), Queue())
	ConfirmDrop(Queue(0), Queue())
	ConfirmDrop(&Fifo{ head: stack(0), length: 1 }, Queue())
	ConfirmDrop(&Fifo{ tail: stack(0), length: 1 }, Queue())

	ConfirmDrop(Queue(0, 1), Queue(1))
	ConfirmDrop(&Fifo{ head: stack(0, 1), length: 2 }, Queue(1))
	ConfirmDrop(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, Queue(1))
	ConfirmDrop(&Fifo{ tail: stack(1, 0), length: 2 }, Queue(1))

	ConfirmDrop(Queue(0, 1, 2), Queue(1, 2))
	ConfirmDrop(&Fifo{ head: stack(0, 1, 2), length: 3 }, Queue(1, 2))
	ConfirmDrop(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, Queue(1, 2))
	ConfirmDrop(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, Queue(1, 2))
	ConfirmDrop(&Fifo{ tail: stack(2, 1, 0), length: 3 }, Queue(1, 2))

	ConfirmDrop(Queue(0, 1, 2, 3), Queue(1, 2, 3))
	ConfirmDrop(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, Queue(1, 2, 3))
	ConfirmDrop(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, Queue(1, 2, 3))
	ConfirmDrop(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, Queue(1, 2, 3))
	ConfirmDrop(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, Queue(1, 2, 3))
	ConfirmDrop(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, Queue(1, 2, 3))
}

func TestFifoDup(t *testing.T) {
	RefuteDup := func(s *Fifo) {
		defer ConfirmPanic(t, "%v.Dup() should panic", s)()
		s.Dup()
	}

	ConfirmDup := func(s, r *Fifo) {
		if x := s.Dup(); !x.Equal(r) {
			t.Fatalf("%v.Dup() should be %v but is %v", s, r, x)
		}
	}

	RefuteDup(nil)
	RefuteDup(Queue())

	ConfirmDup(Queue(0), Queue(0, 0))
	ConfirmDup(&Fifo{ head: stack(0), length: 1 }, Queue(0, 0))
	ConfirmDup(&Fifo{ tail: stack(0), length: 1 }, Queue(0, 0))

	ConfirmDup(Queue(0, 1), Queue(0, 1, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1), length: 2 }, Queue(0, 1, 0))
	ConfirmDup(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, Queue(0, 1, 0))
	ConfirmDup(&Fifo{ tail: stack(1, 0), length: 2 }, Queue(0, 1, 0))

	ConfirmDup(Queue(0, 1, 2), Queue(0, 1, 2, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1, 2), length: 3 }, Queue(0, 1, 2, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, Queue(0, 1, 2, 0))
	ConfirmDup(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, Queue(0, 1, 2, 0))
	ConfirmDup(&Fifo{ tail: stack(2, 1, 0), length: 3 }, Queue(0, 1, 2, 0))

	ConfirmDup(Queue(0, 1, 2, 3), Queue(0, 1, 2, 3, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, Queue(0, 1, 2, 3, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, Queue(0, 1, 2, 3, 0))
	ConfirmDup(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, Queue(0, 1, 2, 3, 0))
	ConfirmDup(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, Queue(0, 1, 2, 3, 0))
	ConfirmDup(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, Queue(0, 1, 2, 3, 0))
}

func TestFifoSwap(t *testing.T) {
	ConfirmSwap := func(s, r *Fifo) {
		vs := s.String()
		if x := s.Swap(); !x.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", vs, r, x)
		}
	}

	ConfirmSwap(nil, nil)
	ConfirmSwap(nil, Queue())

	ConfirmSwap(Queue(), Queue())
	ConfirmSwap(Queue(), nil)
	ConfirmSwap(Queue(0), Queue(0))
	ConfirmSwap(Queue(0, 1), Queue(1, 0))
	ConfirmSwap(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, Queue(1, 0))

	ConfirmSwap(Queue(0, 1, 2), Queue(2, 1, 0))
	ConfirmSwap(&Fifo{ head: stack(0, 1, 2), length: 3 }, Queue(2, 1, 0))
	ConfirmSwap(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, Queue(2, 1, 0))
	ConfirmSwap(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, Queue(2, 1, 0))
	ConfirmSwap(&Fifo{ tail: stack(2, 1, 0), length: 3 }, Queue(2, 1, 0))

	ConfirmSwap(Queue(0, 1, 2, 3), Queue(3, 1, 2, 0))
	ConfirmSwap(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, Queue(3, 1, 2, 0))
	ConfirmSwap(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, Queue(3, 1, 2, 0))
	ConfirmSwap(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, Queue(3, 1, 2, 0))
	ConfirmSwap(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, Queue(3, 1, 2, 0))
	ConfirmSwap(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, Queue(3, 1, 2, 0))
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
	ConfirmCopy(&Fifo{ head: stack(0), length: 1 }, 0, Queue())
	ConfirmCopy(&Fifo{ tail: stack(0), length: 1 }, 0, Queue())

	ConfirmCopy(Queue(0), 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0), length: 1 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ tail: stack(0), length: 1 }, 1, Queue(0))

	ConfirmCopy(Queue(0), 2, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0), length: 1 }, 2, Queue(0))
	ConfirmCopy(&Fifo{ tail: stack(0), length: 1 }, 2, Queue(0))

	ConfirmCopy(Queue(0, 1), 0, Queue())
	ConfirmCopy(&Fifo{ head: stack(0, 1), length: 2 }, 0, Queue())
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 0, Queue())
	ConfirmCopy(&Fifo{ tail: stack(1, 0), length: 2 }, 0, Queue())

	ConfirmCopy(Queue(0, 1), 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0, 1), length: 2 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ tail: stack(1, 0), length: 2 }, 1, Queue(0))

	ConfirmCopy(Queue(0, 1), 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0, 1), length: 2 }, 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ tail: stack(1, 0), length: 2 }, 2, Queue(0, 1))

	ConfirmCopy(Queue(0, 1), 3, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0, 1), length: 2 }, 3, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 3, Queue(0, 1))
	ConfirmCopy(&Fifo{ tail: stack(1, 0), length: 2 }, 3, Queue(0, 1))

	ConfirmCopy(Queue(0, 1, 2), 0, Queue())
	ConfirmCopy(&Fifo{ head: stack(0, 1, 2), length: 3 }, 0, Queue())
	ConfirmCopy(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 0, Queue())
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 0, Queue())
	ConfirmCopy(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 0, Queue())

	ConfirmCopy(Queue(0, 1, 2), 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0, 1, 2), length: 3 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 1, Queue(0))
	ConfirmCopy(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 1, Queue(0))

	ConfirmCopy(Queue(0, 1, 2), 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0, 1, 2), length: 3 }, 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 2, Queue(0, 1))
	ConfirmCopy(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 2, Queue(0, 1))

	ConfirmCopy(Queue(0, 1, 2), 3, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0, 1, 2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 3, Queue(0, 1, 2))

	ConfirmCopy(Queue(0, 1, 2), 4, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0, 1, 2), length: 3 }, 4, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 4, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 4, Queue(0, 1, 2))
	ConfirmCopy(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 4, Queue(0, 1, 2))
}

func TestFifoMove(t *testing.T) {
	ConfirmMove := func(s *Fifo, n int, r *Fifo) {
		x := s.Move(n)
		if !x.Equal(r) {
			t.Fatalf("%v.Move(%v) should be %v but is %v", s, n, r, x)
		}

		if s != nil {
			l := s.length - n
			if l < 0 {
				l = 0
			}
			switch {
			case x == nil && l == 0:
			case x == nil && l != 0:
				t.Fatalf("%v.Move(%v).length should be %v but is 0", s, n, l)
			case x.length != l:
				t.Fatalf("%v.Move(%v).length should be %v but is %v", s, n, l, x.length)
			}
		}
	}

	ConfirmMove(nil, 0, nil)
	ConfirmMove(nil, 0, Queue())
	ConfirmMove(nil, 1, nil)
	ConfirmMove(nil, 1, Queue())

	ConfirmMove(Queue(), 0, nil)
	ConfirmMove(Queue(), 0, Queue())
	ConfirmMove(Queue(), 1, nil)
	ConfirmMove(Queue(), 1, Queue())
	
	ConfirmMove(Queue(0), 0, Queue(0))
	ConfirmMove(&Fifo{ head: stack(0), length: 1 }, 0, Queue(0))
	ConfirmMove(&Fifo{ tail: stack(0), length: 1 }, 0, Queue(0))
	
	ConfirmMove(Queue(0), 1, Queue())
	ConfirmMove(&Fifo{ head: stack(0), length: 1 }, 1, Queue())
	ConfirmMove(&Fifo{ tail: stack(0), length: 1 }, 1, Queue())

	ConfirmMove(Queue(0, 1), 0, Queue(0, 1))
	ConfirmMove(&Fifo{ head: stack(0, 1), length: 2 }, 0, Queue(0, 1))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 0, Queue(0, 1))
	ConfirmMove(&Fifo{ tail: stack(1, 0), length: 2 }, 0, Queue(0, 1))

	ConfirmMove(Queue(0, 1), 1, Queue(1))
	ConfirmMove(&Fifo{ head: stack(0, 1), length: 2 }, 1, Queue(1))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 1, Queue(1))
	ConfirmMove(&Fifo{ tail: stack(1, 0), length: 2 }, 1, Queue(1))

	ConfirmMove(Queue(0, 1), 2, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1), length: 2 }, 2, Queue())
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 2, Queue())
	ConfirmMove(&Fifo{ tail: stack(1, 0), length: 2 }, 2, Queue())

	ConfirmMove(Queue(0, 1, 2), 0, Queue(0, 1, 2))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmMove(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 0, Queue(0, 1, 2))

	ConfirmMove(Queue(0, 1, 2), 1, Queue(1, 2))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), length: 3 }, 1, Queue(1, 2))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 1, Queue(1, 2))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 1, Queue(1, 2))
	ConfirmMove(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 1, Queue(1, 2))

	ConfirmMove(Queue(0, 1, 2), 2, Queue(2))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), length: 3 }, 2, Queue(2))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 2, Queue(2))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 2, Queue(2))
	ConfirmMove(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 2, Queue(2))

	ConfirmMove(Queue(0, 1, 2), 3, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), length: 3 }, 3, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 3, Queue())
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 3, Queue())
	ConfirmMove(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 3, Queue())

	
	ConfirmMove(Queue(0, 1, 2, 3), 0, Queue(0, 1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmMove(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 0, Queue(0, 1, 2, 3))

	ConfirmMove(Queue(0, 1, 2, 3), 1, Queue(1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 1, Queue(1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 1, Queue(1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 1, Queue(1, 2, 3))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 1, Queue(1, 2, 3))
	ConfirmMove(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 1, Queue(1, 2, 3))

	ConfirmMove(Queue(0, 1, 2, 3), 2, Queue(2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 2, Queue(2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 2, Queue(2, 3))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 2, Queue(2, 3))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 2, Queue(2, 3))
	ConfirmMove(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 2, Queue(2, 3))

	ConfirmMove(Queue(0, 1, 2, 3), 3, Queue(3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 3, Queue(3))
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 3, Queue(3))
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 3, Queue(3))
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 3, Queue(3))
	ConfirmMove(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 3, Queue(3))

	ConfirmMove(Queue(0, 1, 2, 3), 4, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 4, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 4, Queue())
	ConfirmMove(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 4, Queue())
	ConfirmMove(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 4, Queue())
	ConfirmMove(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 4, Queue())
}

func TestFifoPick(t *testing.T) {
	ConfirmPick := func(s *Fifo, n int, r *Fifo) {
		if x := s.Pick(n); !x.Equal(r) {
			t.Fatalf("%v.Pick(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmPick(nil, 0, nil)
	ConfirmPick(nil, 1, nil)

	ConfirmPick(Queue(), 0, nil)
	ConfirmPick(Queue(), 1, nil)

	ConfirmPick(Queue(0), 0, Queue(0, 0))
	ConfirmPick(&Fifo{ head: stack(0), length: 1 }, 0, Queue(0, 0))
	ConfirmPick(Queue(0), 1, Queue(0))
	
	ConfirmPick(Queue(0, 1), 0, Queue(0, 1, 0))
	ConfirmPick(&Fifo{ head: stack(0, 1), length: 2 }, 0, Queue(0, 1, 0))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 0, Queue(0, 1, 0))
	ConfirmPick(&Fifo{ tail: stack(1, 0), length: 2 }, 0, Queue(0, 1, 0))

	ConfirmPick(Queue(0, 1), 1, Queue(0, 1, 1))
	ConfirmPick(&Fifo{ head: stack(0, 1), length: 2 }, 1, Queue(0, 1, 1))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 1, Queue(0, 1, 1))
	ConfirmPick(&Fifo{ tail: stack(1, 0), length: 2 }, 1, Queue(0, 1, 1))
	
	ConfirmPick(Queue(0, 1), 2, Queue(0, 1))
	ConfirmPick(&Fifo{ head: stack(0, 1), length: 2 }, 2, Queue(0, 1))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 2, Queue(0, 1))
	ConfirmPick(&Fifo{ tail: stack(1, 0), length: 2 }, 2, Queue(0, 1))

	ConfirmPick(Queue(0, 1, 2), 0, Queue(0, 1, 2, 0))
	ConfirmPick(&Fifo{ head: stack(0, 1, 2), length: 3 }, 0, Queue(0, 1, 2, 0))
	ConfirmPick(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 0, Queue(0, 1, 2, 0))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 0, Queue(0, 1, 2, 0))
	ConfirmPick(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 0, Queue(0, 1, 2, 0))

	ConfirmPick(Queue(0, 1, 2), 1, Queue(0, 1, 2, 1))
	ConfirmPick(&Fifo{ head: stack(0, 1, 2), length: 3 }, 1, Queue(0, 1, 2, 1))
	ConfirmPick(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 1, Queue(0, 1, 2, 1))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 1, Queue(0, 1, 2, 1))
	ConfirmPick(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 1, Queue(0, 1, 2, 1))

	ConfirmPick(Queue(0, 1, 2), 2, Queue(0, 1, 2, 2))
	ConfirmPick(&Fifo{ head: stack(0, 1, 2), length: 3 }, 2, Queue(0, 1, 2, 2))
	ConfirmPick(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 2, Queue(0, 1, 2, 2))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 2, Queue(0, 1, 2, 2))
	ConfirmPick(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 2, Queue(0, 1, 2, 2))
	
	ConfirmPick(Queue(0, 1, 2), 3, Queue(0, 1, 2))
	ConfirmPick(&Fifo{ head: stack(0, 1, 2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmPick(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmPick(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmPick(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 3, Queue(0, 1, 2))
}

func TestFifoRoll(t *testing.T) {
	ConfirmRoll := func(s *Fifo, n int, r *Fifo) {
		if x := s.Roll(n); !x.Equal(r) {
			t.Fatalf("%v.Roll(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmRoll(nil, 0, nil)
	ConfirmRoll(nil, 1, nil)

	ConfirmRoll(Queue(), 0, nil)
	ConfirmRoll(Queue(), 1, nil)

	ConfirmRoll(Queue(0), 0, Queue(0))
	ConfirmRoll(&Fifo{ head: stack(0), length: 1 }, 0, Queue(0))
	ConfirmRoll(&Fifo{ tail: stack(0), length: 1 }, 0, Queue(0))

	ConfirmRoll(Queue(0), 1, Queue(0))
	ConfirmRoll(&Fifo{ head: stack(0), length: 1 }, 1, Queue(0))
	ConfirmRoll(&Fifo{ tail: stack(0), length: 1 }, 1, Queue(0))
	
	ConfirmRoll(Queue(0, 1), 0, Queue(0, 1))
	ConfirmRoll(&Fifo{ head: stack(0, 1), length: 2 }, 0, Queue(0, 1))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 0, Queue(0, 1))
	ConfirmRoll(&Fifo{ tail: stack(1, 0), length: 2 }, 0, Queue(0, 1))

	ConfirmRoll(Queue(0, 1), 1, Queue(1, 0))
	ConfirmRoll(&Fifo{ head: stack(0, 1), length: 2 }, 1, Queue(1, 0))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 1, Queue(1, 0))
	ConfirmRoll(&Fifo{ tail: stack(1, 0), length: 2 }, 1, Queue(1, 0))
	
	ConfirmRoll(Queue(0, 1), 2, Queue(0, 1))
	ConfirmRoll(&Fifo{ head: stack(0, 1), length: 2 }, 2, Queue(0, 1))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(1), length: 2 }, 2, Queue(0, 1))
	ConfirmRoll(&Fifo{ tail: stack(1, 0), length: 2 }, 2, Queue(0, 1))

	ConfirmRoll(Queue(0, 1, 2), 0, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 0, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 0, Queue(0, 1, 2))

	ConfirmRoll(Queue(0, 1, 2), 1, Queue(1, 0, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), length: 3 }, 1, Queue(1, 0, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 1, Queue(1, 0, 2))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 1, Queue(1, 0, 2))
	ConfirmRoll(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 1, Queue(1, 0, 2))

	ConfirmRoll(Queue(0, 1, 2), 2, Queue(2, 0, 1))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), length: 3 }, 2, Queue(2, 0, 1))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 2, Queue(2, 0, 1))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 2, Queue(2, 0, 1))
	ConfirmRoll(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 2, Queue(2, 0, 1))
	
	ConfirmRoll(Queue(0, 1, 2), 3, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(2), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(2, 1), length: 3 }, 3, Queue(0, 1, 2))
	ConfirmRoll(&Fifo{ tail: stack(2, 1, 0), length: 3 }, 3, Queue(0, 1, 2))

	ConfirmRoll(Queue(0, 1, 2, 3), 0, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 0, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 0, Queue(0, 1, 2, 3))

	ConfirmRoll(Queue(0, 1, 2, 3), 1, Queue(1, 0, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 1, Queue(1, 0, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 1, Queue(1, 0, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 1, Queue(1, 0, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 1, Queue(1, 0, 2, 3))
	ConfirmRoll(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 1, Queue(1, 0, 2, 3))

	ConfirmRoll(Queue(0, 1, 2, 3), 2, Queue(2, 0, 1, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 2, Queue(2, 0, 1, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 2, Queue(2, 0, 1, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 2, Queue(2, 0, 1, 3))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 2, Queue(2, 0, 1, 3))
	ConfirmRoll(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 2, Queue(2, 0, 1, 3))

	ConfirmRoll(Queue(0, 1, 2, 3), 3, Queue(3, 0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 3, Queue(3, 0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 3, Queue(3, 0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 3, Queue(3, 0, 1, 2))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 3, Queue(3, 0, 1, 2))
	ConfirmRoll(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 3, Queue(3, 0, 1, 2))

	ConfirmRoll(Queue(0, 1, 2, 3), 4, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), length: 4 }, 4, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(3), length: 4 }, 4, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(3, 2), length: 4 }, 4, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0), tail: stack(3, 2, 1), length: 4 }, 4, Queue(0, 1, 2, 3))
	ConfirmRoll(&Fifo{ tail: stack(3, 2, 1, 0), length: 4 }, 4, Queue(0, 1, 2, 3))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 0, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 0, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 0, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 0, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 0, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 0, Queue(0, 1, 2, 3, 4))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 1, Queue(1, 0, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 1, Queue(1, 0, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 1, Queue(1, 0, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 1, Queue(1, 0, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 1, Queue(1, 0, 2, 3, 4))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 1, Queue(1, 0, 2, 3, 4))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 2, Queue(2, 0, 1, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 2, Queue(2, 0, 1, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 2, Queue(2, 0, 1, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 2, Queue(2, 0, 1, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 2, Queue(2, 0, 1, 3, 4))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 2, Queue(2, 0, 1, 3, 4))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 3, Queue(3, 0, 1, 2, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 3, Queue(3, 0, 1, 2, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 3, Queue(3, 0, 1, 2, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 3, Queue(3, 0, 1, 2, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 3, Queue(3, 0, 1, 2, 4))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 3, Queue(3, 0, 1, 2, 4))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 4, Queue(4, 0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 4, Queue(4, 0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 4, Queue(4, 0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 4, Queue(4, 0, 1, 2, 3))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 4, Queue(4, 0, 1, 2, 3))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 4, Queue(4, 0, 1, 2, 3))

	ConfirmRoll(Queue(0, 1, 2, 3, 4), 5, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3, 4), length: 5 }, 5, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2, 3), tail: stack(4), length: 5 }, 5, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1, 2), tail: stack(4, 3), length: 5 }, 5, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ head: stack(0, 1), tail: stack(4, 3, 2), length: 5 }, 5, Queue(0, 1, 2, 3, 4))
	ConfirmRoll(&Fifo{ tail: stack(4, 3, 2, 1, 0), length: 5 }, 5, Queue(0, 1, 2, 3, 4))
}