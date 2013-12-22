package greenspun

import (
//	"fmt"
	"testing"
)

func TestStackString(t *testing.T) {
	ConfirmString := func(s *stackCell, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "()")
	ConfirmString(&stackCell{ data: 1 }, "(1)")
	ConfirmString(&stackCell{ data: 1, stackCell: &stackCell{ data: 2 } }, "(1 2)")
	ConfirmString(&stackCell{ data: 1, stackCell: &stackCell{ data: 2, stackCell: &stackCell{ data: 3 } } }, "(1 2 3)")
	ConfirmString(stack(1), "(1)")
	ConfirmString(stack(1, 2), "(1 2)")
	ConfirmString(stack(1, 2, 3), "(1 2 3)")
}

func TestStackEqual(t *testing.T) {
	ConfirmEqual := func(x, y *stackCell, r bool) {
		switch {
		case x.Equal(y) != r:
			t.Fatalf("1: %v.Equal(%v) should be %v", x, y, r)
		case y.Equal(x) != r:
			t.Fatalf("2: %v.Equal(%v) should be %v", y, x, r)
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(stack(1), nil, false)
	ConfirmEqual(nil, stack(1), false)
	ConfirmEqual(stack(1), stack(1), true)
}

func TestStackEach(t *testing.T) {
	s := stack(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmEach := func(c *stackCell, f interface{}) {
		count = 0
		c.Each(f)
		if l := c.Len(); l != count {
			t.Fatalf("%v.Each() should have iterated %v times not %v times", c, l, count)
		}
	}

	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("1: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	ConfirmEach(s, func(index int, i interface{}) {
		if i != index {
			t.Fatalf("2: %v.Each() element %v erroneously reported as %v", s, index, i)
		}
		count++
	})

	ConfirmEach(s, func(key, i interface{}) {
		if i.(int) != key.(int) {
			t.Fatalf("3: %v.Each() element %v erroneously reported as %v", s, key, i)
		}
		count++
	})

	s = stack()
	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("4: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	s = nil
	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("5: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})
}

func TestStackPush(t *testing.T) {
	ConfirmPush := func(s *stackCell, v interface{}, r *stackCell) {
		if x := s.Push(v); !x.Equal(r) {
			t.Fatalf("%v.Push(%v) should be %v but is %v", s, v, r, x)
		}
	}

	ConfirmPush(nil, nil, stack(nil))
	ConfirmPush(nil, 1, stack(1))
	ConfirmPush(stack(1), 1, stack(1, 1))
}

func TestStackPeek(t *testing.T) {
	RefutePeek := func(s *stackCell) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Peek() should panic", vs)()
		s.Peek()
	}

	ConfirmPeek := func(s *stackCell, r interface{}) {
		if x := s.Peek(); x != r {
			t.Fatalf("%v.Peek() should be %v but is %v", s, r, x)
		}
	}

	RefutePeek(nil)
	RefutePeek(stack())
	ConfirmPeek(stack(1), 1)
	ConfirmPeek(stack(1, 2), 1)
}

func TestStackPop(t *testing.T) {
	RefutePop := func(s *stackCell) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pop() should panic", vs)()
		s.Pop()
	}

	ConfirmPop := func(s *stackCell, r interface{}, z *stackCell) {
		switch x, n := s.Pop(); {
		case x != r:
			t.Fatalf("%v.Pop() should be %v but is %v", s, r, x)
		case !n.Equal(z):
			t.Fatalf("%v.Pop() should be [%v, %v] but is [%v, %v]", s, r, z, x, n)
		}
	}

	RefutePop(nil)
	RefutePop(stack())
	ConfirmPop(stack(1), 1, nil)
	ConfirmPop(stack(1, 2), 1, stack(2))
}

func TestStackLen(t *testing.T) {
	ConfirmLen := func(s *stackCell, r int) {
		if l := s.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(stack(nil), 1)
	ConfirmLen(stack(nil, nil), 2)
}

func TestStackDrop(t *testing.T) {
	ConfirmDrop := func(s, r *stackCell) {
		vs := s.String()
		if x := s.Drop(); !x.Equal(r) {
			t.Fatalf("%v.Drop() should leave %v but leaves %v", vs, r, x)
		}
	}

	ConfirmDrop(nil, nil)
	ConfirmDrop(stack(), nil)
	ConfirmDrop(stack(), stack())
	ConfirmDrop(stack(0), stack())
	ConfirmDrop(stack(1, 0), stack(0))
	ConfirmDrop(stack(2, 1, 0), stack(1, 0))
}

func TestStackDup(t *testing.T) {
	RefuteDup := func(s *stackCell) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Dup() should panic", vs)()
		s.Dup()
	}

	ConfirmDup := func(s, r *stackCell) {
		if x := s.Dup(); !x.Equal(r) {
			t.Fatalf("%v.Dup() should be %v but is %v", s, r, x)
		}
	}

	RefuteDup(nil)
	RefuteDup(stack())
	ConfirmDup(stack(1), stack(1, 1))
	ConfirmDup(stack(1, 2), stack(1, 1, 2))
}

func TestStackSwap(t *testing.T) {
	RefuteSwap := func(s *stackCell) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Swap() should panic", vs)()
		s.Swap()
	}

	ConfirmSwap := func(s, r *stackCell) {
		if x := s.Swap(); !x.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", s, r, x)
		}
	}

	RefuteSwap(nil)
	RefuteSwap(stack())
	RefuteSwap(stack(1))
	ConfirmSwap(stack(1, 2), stack(2, 1))
	ConfirmSwap(stack(1, 2, 3), stack(2, 1, 3))
}

func TestStackCopy(t *testing.T) {
	ConfirmCopy := func(s *stackCell, n int, r *stackCell) {
		if x := s.Copy(n); !x.Equal(r) {
			t.Fatalf("%v.Copy(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmCopy(nil, 0, nil)
	ConfirmCopy(nil, 1, nil)

	ConfirmCopy(stack(), 0, stack())
	ConfirmCopy(stack(), 1, stack())

	ConfirmCopy(stack(0), 0, stack())
	ConfirmCopy(stack(0), 1, stack(0))
	ConfirmCopy(stack(0), 2, stack(0))

	ConfirmCopy(stack(0, 1), 0, stack())
	ConfirmCopy(stack(0, 1), 1, stack(0))
	ConfirmCopy(stack(0, 1), 2, stack(0, 1))
	ConfirmCopy(stack(0, 1), 3, stack(0, 1))

	ConfirmCopy(stack(0, 1, 2), 0, stack())
	ConfirmCopy(stack(0, 1, 2), 1, stack(0))
	ConfirmCopy(stack(0, 1, 2), 2, stack(0, 1))
	ConfirmCopy(stack(0, 1, 2), 3, stack(0, 1, 2))
	ConfirmCopy(stack(0, 1, 2), 4, stack(0, 1, 2))
}

func TestStackClone(t *testing.T) {
	ConfirmClone := func(s, r *stackCell) {
		if x := s.Clone(); !x.Equal(r) {
			t.Fatalf("%v.Clone() should be %v but is %v", s, r, x)
		}
	}

	ConfirmClone(nil, nil)
	ConfirmClone(stack(), stack())
	ConfirmClone(stack(0), stack(0))
	ConfirmClone(stack(0, 1), stack(0, 1))
	ConfirmClone(stack(0, 1, 2), stack(0, 1, 2))
}

func TestStackMove(t *testing.T) {
	RefuteMove := func(s *stackCell, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Move(%v) should panic", vs, x)()
		s.Move(x)
	}

	ConfirmMove := func(s *stackCell, n int, r *stackCell) {
		if x := s.Move(n); !x.Equal(r) {
			t.Fatalf("%v.Move(%v) should be %v but is %v", s, n, r, x)
		}
	}

	RefuteMove(nil, 0)
	RefuteMove(nil, 1)

	RefuteMove(stack(), 0)
	RefuteMove(stack(), 1)

	ConfirmMove(stack(0), 0, stack(0))
	RefuteMove(stack(0), 1)

	ConfirmMove(stack(0, 1), 0, stack(0, 1))
	ConfirmMove(stack(0, 1), 1, stack(1))
	RefuteMove(stack(0, 1), 2)
}

func TestStackPick(t *testing.T) {
	RefutePick := func(s *stackCell, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Pick(%v) should panic", vs, x)()
		s.Pick(x)
	}

	ConfirmPick := func(s *stackCell, n int, r *stackCell) {
		if x := s.Pick(n); !x.Equal(r) {
			t.Fatalf("%v.Pick(%v) should be %v but is %v", s, n, r, x)
		}
	}

	RefutePick(nil, 0)
	RefutePick(nil, 1)

	RefutePick(stack(), 0)
	RefutePick(stack(), 1)

	ConfirmPick(stack(0), 0, stack(0, 0))
	RefutePick(stack(0), 1)

	ConfirmPick(stack(0, 1), 0, stack(0, 0, 1))
	ConfirmPick(stack(0, 1), 1, stack(1, 0, 1))
	RefutePick(stack(0, 1), 2)

	ConfirmPick(stack(0, 1, 2), 0, stack(0, 0, 1, 2))
	ConfirmPick(stack(0, 1, 2), 1, stack(1, 0, 1, 2))
	ConfirmPick(stack(0, 1, 2), 2, stack(2, 0, 1, 2))
	RefutePick(stack(0, 1, 2), 3)
}

func TestStackRoll(t *testing.T) {
	RefuteRoll := func(s *stackCell, x int) {
		vs := s.String()
		defer ConfirmPanic(t, "%v.Roll(%v) should panic", vs, x)()
		s.Roll(x)
	}

	ConfirmRoll := func(s *stackCell, n int, r *stackCell) {
		if x := s.Roll(n); !x.Equal(r) {
			t.Fatalf("%v.Roll(%v) should be %v but is %v", s, n, r, x)
		}
	}

	RefuteRoll(nil, 0)
	RefuteRoll(nil, 1)

	RefuteRoll(stack(), 0)
	RefuteRoll(stack(), 1)

	ConfirmRoll(stack(0), 0, stack(0))
	RefuteRoll(stack(0), 1)

	ConfirmRoll(stack(0, 1), 0, stack(0, 1))
	ConfirmRoll(stack(0, 1), 1, stack(1, 0))
	RefuteRoll(stack(0, 1), 2)

	ConfirmRoll(stack(0, 1, 2), 0, stack(0, 1, 2))
	ConfirmRoll(stack(0, 1, 2), 1, stack(1, 0, 2))
	ConfirmRoll(stack(0, 1, 2), 2, stack(2, 0, 1))
	RefuteRoll(stack(0, 1, 2), 3)

	ConfirmRoll(stack(0, 1, 2, 3), 0, stack(0, 1, 2, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 1, stack(1, 0, 2, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 2, stack(2, 0, 1, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 3, stack(3, 0, 1, 2))
	RefuteRoll(stack(0, 1, 2, 3), 4)
}

func TestStackReverse(t *testing.T) {
	ConfirmReverse := func(s, r *stackCell) {
		if x := s.Reverse(); !x.Equal(r) {
			t.Fatalf("%v.Reverse() should be %v but is %v", s, r, x)
		}
	}

	ConfirmReverse(nil, nil)
	ConfirmReverse(stack(0), stack(0))
	ConfirmReverse(stack(0, 1), stack(1, 0))
	ConfirmReverse(stack(0, 1, 2), stack(2, 1, 0))
}