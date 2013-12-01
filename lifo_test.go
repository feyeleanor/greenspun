package greenspun

import (
//	"fmt"
	"testing"
)

func TestLifoString(t *testing.T) {
	ConfirmString := func(s *Lifo, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "()")
	ConfirmString(new(Lifo), "()")
	ConfirmString(&Lifo{ depth: 1, stackCell: stack(1) }, "(1)")
	ConfirmString(&Lifo{ depth: 2, stackCell: stack(1, 2) }, "(1 2)")
	ConfirmString(&Lifo{ depth: 3, stackCell: stack(1, 2, 3) }, "(1 2 3)")
	ConfirmString(Stack(1), "(1)")
	ConfirmString(Stack(1, 2), "(1 2)")
	ConfirmString(Stack(1, 2, 3), "(1 2 3)")
}

func TestLifoEqual(t *testing.T) {
	ConfirmEqual := func(x *Lifo, y Equatable, r bool) {
		if z := x.Equal(y); z != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", x, y, r, z)
		}
	}

	ConfirmEqual(nil, (*Lifo)(nil), true)
	ConfirmEqual(Stack(1), (*Lifo)(nil), false)
	ConfirmEqual(nil, Stack(), true)
	ConfirmEqual(Stack(), nil, true)
	ConfirmEqual(nil, Stack(1), false)
	ConfirmEqual(Stack(1), Stack(1), true)

	ConfirmEqual(Stack(1, 2), (*Lifo)(nil), false)
	ConfirmEqual(nil, Stack(1, 2), false)
	ConfirmEqual(Stack(1, 2), Stack(1, 2), true)
	ConfirmEqual(Stack(1, 2), Stack(2, 1), false)
	ConfirmEqual(Stack(2, 1), Stack(1, 2), false)
	ConfirmEqual(Stack(2, 1), Stack(2, 1), true)

	ConfirmEqual(Stack(1, 2, 3), Stack(1, 2, 3), true)
	ConfirmEqual(Stack(2, 1, 3), Stack(1, 2, 3), false)
	ConfirmEqual(Stack(3, 2, 1), Stack(3, 2, 1), true)

	ConfirmEqual(nil, stack(1), false)
	ConfirmEqual(Stack(1), stack(1), true)

	ConfirmEqual(Stack(1, 2), (*stackCell)(nil), false)
	ConfirmEqual(nil, stack(1, 2), false)
	ConfirmEqual(Stack(1, 2), stack(1, 2), true)
	ConfirmEqual(Stack(1, 2), stack(2, 1), false)
	ConfirmEqual(Stack(2, 1), stack(1, 2), false)
	ConfirmEqual(Stack(2, 1), stack(2, 1), true)

	ConfirmEqual(Stack(1, 2, 3), stack(1, 2, 3), true)
	ConfirmEqual(Stack(2, 1, 3), stack(1, 2, 3), false)
	ConfirmEqual(Stack(3, 2, 1), stack(3, 2, 1), true)
}

func TestLifoPush(t *testing.T) {
	RefutePush := func(s *Lifo, v interface{}) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Push(%v) should panic", vs, v)()
		s.Push(v)
	}

	ConfirmPush := func(s *Lifo, v interface{}, r *Lifo) {
		vs := s.String()
		if s.Push(v); !s.Equal(r) {
			t.Fatalf("%v.Push(%v) should be %v but is %v", vs, v, r, s)
		}
	}

	RefutePush(nil, nil)
	RefutePush(nil, 1)
	ConfirmPush(Stack(1), 1, Stack(1, 1))
	ConfirmPush(Stack(1, 2), 1, Stack(1, 1, 2))
}

func TestLifoPeek(t *testing.T) {
	RefutePeek := func(s *Lifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Peek() should panic", vs)()
		s.Peek()
	}

	ConfirmPeek := func(s *Lifo, r interface{}) {
		if x := s.Peek(); x != r {
			t.Fatalf("%v.Peek() should be %v but is %v", s, r, x)
		}
	}

	RefutePeek(nil)
	RefutePeek(Stack())
	ConfirmPeek(Stack(0), 0)
	ConfirmPeek(Stack(1, 0), 1)
	ConfirmPeek(Stack(2, 1, 0), 2)
}

func TestLifoPop(t *testing.T) {
	RefutePop := func(s *Lifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pop() should panic", vs)()
		s.Pop()
	}

	ConfirmPop := func(s *Lifo, r interface{}, n *Lifo) {
		vs := s.String()
		switch x := s.Pop(); {
		case x != r:
			t.Fatalf("%v.Pop() should be %v but is %v", vs, r, x)
		case !s.Equal(n):
			t.Fatalf("%v.Pop() should leave %v but leaves %v", vs, n, s)
		}
	}

	RefutePop(nil)
	RefutePop(Stack())
	ConfirmPop(Stack(0), 0, Stack())
	ConfirmPop(Stack(1, 0), 1, Stack(0))
	ConfirmPop(Stack(2, 1, 0), 2, Stack(1, 0))
}

func TestLifoLen(t * testing.T) {
	ConfirmLen := func(s *Lifo, r int) {
		if l := s.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(Stack(), 0)
	ConfirmLen(Stack(0), 1)
	ConfirmLen(Stack(1, 0), 2)
}

func TestLifoDrop(t *testing.T) {
	RefuteDrop := func(s *Lifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Drop() should panic", vs)()
		s.Drop()
	}

	ConfirmDrop := func(s, r *Lifo) {
		vs := s.String()
		if s.Drop(); !s.Equal(r) {
			t.Fatalf("%v.Drop() should leave %v but leaves %v", vs, r, s)
		}
	}

	RefuteDrop(nil)
	ConfirmDrop(Stack(), nil)
	ConfirmDrop(Stack(), Stack())
	ConfirmDrop(Stack(0), Stack())
	ConfirmDrop(Stack(1, 0), Stack(0))
	ConfirmDrop(Stack(2, 1, 0), Stack(1, 0))
}

func TestLifoDup(t *testing.T) {
	RefuteDup := func(s *Lifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Dup() should panic", vs)()
		s.Dup()
	}

	ConfirmDup := func(s, r *Lifo) {
		vs := s.String()
		if s.Dup(); !s.Equal(r) {
			t.Fatalf("%v.Dup() should be %v but is %v", vs, r, s)
		}
	}

	RefuteDup(nil)
	RefuteDup(Stack())
	ConfirmDup(Stack(1), Stack(1, 1))
	ConfirmDup(Stack(1, 2), Stack(1, 1, 2))
}

func TestLifoSwap(t *testing.T) {
	RefuteSwap := func(s *Lifo) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Swap() should panic", vs)()
		s.Swap()
	}

	ConfirmSwap := func(s, r *Lifo) {
		vs := s.String()
		if s.Swap(); !s.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", vs, r, s)
		}
	}

	RefuteSwap(nil)
	RefuteSwap(Stack())
	RefuteSwap(Stack(0))
	ConfirmSwap(Stack(1, 0), Stack(0, 1))
	ConfirmSwap(Stack(2, 1, 0), Stack(1, 2, 0))
}

func TestLifoCopy(t *testing.T) {
	ConfirmCopy := func(s *Lifo, n int, r *Lifo) {
		if x := s.Copy(n); !x.Equal(r) {
			t.Fatalf("%v.Copy(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmCopy(nil, 0, nil)
	ConfirmCopy(nil, 1, nil)

	ConfirmCopy(Stack(), 0, Stack())
	ConfirmCopy(Stack(), 1, Stack())

	ConfirmCopy(Stack(0), 0, Stack())
	ConfirmCopy(Stack(0), 1, Stack(0))
	ConfirmCopy(Stack(0), 2, Stack(0))

	ConfirmCopy(Stack(0, 1), 0, Stack())
	ConfirmCopy(Stack(0, 1), 1, Stack(0))
	ConfirmCopy(Stack(0, 1), 2, Stack(0, 1))
	ConfirmCopy(Stack(0, 1), 3, Stack(0, 1))

	ConfirmCopy(Stack(0, 1, 2), 0, Stack())
	ConfirmCopy(Stack(0, 1, 2), 1, Stack(0))
	ConfirmCopy(Stack(0, 1, 2), 2, Stack(0, 1))
	ConfirmCopy(Stack(0, 1, 2), 3, Stack(0, 1, 2))
	ConfirmCopy(Stack(0, 1, 2), 4, Stack(0, 1, 2))
}

func TestLifoMove(t *testing.T) {
	RefuteMove := func(s *Lifo, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Move(%v) should panic", vs, x)()
		s.Move(x)
	}

	ConfirmMove := func(s *Lifo, n int, r *Lifo) {
		vs := s.String()
		if s.Move(n); !s.Equal(r) {
			t.Fatalf("%v.Move(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	RefuteMove(nil, 0)
	RefuteMove(nil, 1)

	RefuteMove(Stack(), 0)
	RefuteMove(Stack(), 1)

	ConfirmMove(Stack(0), 0, Stack(0))
	RefuteMove(Stack(0), 1)

	ConfirmMove(Stack(0, 1), 0, Stack(0, 1))
	ConfirmMove(Stack(0, 1), 1, Stack(1))
	RefuteMove(Stack(0, 1), 2)
}

func TestLifoPick(t *testing.T) {
	RefutePick := func(s *Lifo, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pick(%v) should panic", vs, x)()
		s.Pick(x)
	}

	ConfirmPick := func(s *Lifo, n int, r *Lifo) {
		vs := s.String()
		if s.Pick(n); !s.Equal(r) {
			t.Fatalf("%v.Pick(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	RefutePick(nil, 0)
	RefutePick(nil, 1)

	RefutePick(Stack(), 0)
	RefutePick(Stack(), 1)

	ConfirmPick(Stack(0), 0, Stack(0, 0))
	RefutePick(Stack(0), 1)

	ConfirmPick(Stack(0, 1), 0, Stack(0, 0, 1))
	ConfirmPick(Stack(0, 1), 1, Stack(1, 0, 1))
	RefutePick(Stack(0, 1), 2)

	ConfirmPick(Stack(0, 1, 2), 0, Stack(0, 0, 1, 2))
	ConfirmPick(Stack(0, 1, 2), 1, Stack(1, 0, 1, 2))
	ConfirmPick(Stack(0, 1, 2), 2, Stack(2, 0, 1, 2))
	RefutePick(Stack(0, 1, 2), 3)
}

func TestLifoRoll(t *testing.T) {
	RefuteRoll := func(s *Lifo, n int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Roll(%v) should panic", vs, n)()
		s.Roll(n)
	}

	ConfirmRoll := func(s *Lifo, n int, r *Lifo) {
		vs := s.String()
		if s.Roll(n); !s.Equal(r) {
			t.Fatalf("%v.Roll(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	RefuteRoll(nil, 0)
	RefuteRoll(nil, 1)

	RefuteRoll(Stack(), 0)
	RefuteRoll(Stack(), 1)

	ConfirmRoll(Stack(0), 0, Stack(0))
	RefuteRoll(Stack(0), 1)
	RefuteRoll(Stack(0), 2)

	ConfirmRoll(Stack(1, 0), 0, Stack(1, 0))
	ConfirmRoll(Stack(1, 0), 1, Stack(0, 1))
	RefuteRoll(Stack(1, 0), 2)

	ConfirmRoll(Stack(2, 1, 0), 0, Stack(2, 1, 0))
	ConfirmRoll(Stack(2, 1, 0), 1, Stack(1, 2, 0))
	ConfirmRoll(Stack(2, 1, 0), 2, Stack(0, 2, 1))
	RefuteRoll(Stack(2, 1, 0), 3)
}

func TestLifoRplaca(t *testing.T) {
	RefuteRplaca := func(s *Lifo, x interface{}) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Rplaca(%v) should panic", vs, x)()
		s.Rplaca(x)
	}

	ConfirmRplaca := func(s *Lifo, v interface{}, r *Lifo) {
		vs := s.String()
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("%v.Rplaca(%v) should be %v but is", vs, v, r, s)
		}
	}

	RefuteRplaca(nil, 0)
	RefuteRplaca(Stack(), 0)
	ConfirmRplaca(Stack(0), 1, Stack(1))
	ConfirmRplaca(Stack(1, 0), 2, Stack(2, 0))
}

func TestLifoRplacd(t *testing.T) {
	RefuteRplacd := func(s *Lifo, x interface{}) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Rplacd(%v) should panic", vs, x)()
		s.Rplacd(x)
	}

	ConfirmRplacd := func(s *Lifo, v interface{}, r *Lifo) {
		vs := s.String()
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("%v.Rplacd(%v) should be %v but is %v", vs, v, r, s)
		}
	}

	RefuteRplacd(nil, nil)
	RefuteRplacd(nil, 0)
	RefuteRplacd(nil, stack())
	RefuteRplacd(nil, Stack())

	RefuteRplacd(Stack(), nil)
	RefuteRplacd(Stack(), 0)
	RefuteRplacd(Stack(), stack())
	RefuteRplacd(Stack(), Stack())
	RefuteRplacd(Stack(), stack(0))
	RefuteRplacd(Stack(), Stack(0))

	ConfirmRplacd(Stack(0), nil, Stack(0))
	ConfirmRplacd(Stack(0), stack(1), Stack(0, 1))
	ConfirmRplacd(Stack(0), Stack(1), Stack(0, 1))

	ConfirmRplacd(Stack(1, 0), nil, Stack(1))
	ConfirmRplacd(Stack(1, 0), stack(), Stack(1))
	ConfirmRplacd(Stack(1, 0), Stack(), Stack(1))
	ConfirmRplacd(Stack(1, 0), stack(2), Stack(1, 2))
	ConfirmRplacd(Stack(1, 0), Stack(2), Stack(1, 2))
}