package greenspun

import "testing"

func TestStackString(t *testing.T) {
	ConfirmString := func(s *stackCell, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "||")
	ConfirmString(&stackCell{ data: 1 }, "|1|")
	ConfirmString(&stackCell{ data: 1, stackCell: &stackCell{ data: 2 } }, "|1 2|")
	ConfirmString(&stackCell{ data: 1, stackCell: &stackCell{ data: 2, stackCell: &stackCell{ data: 3 } } }, "|1 2 3|")
	ConfirmString(stack(1), "|1|")
	ConfirmString(stack(1, 2), "|1 2|")
	ConfirmString(stack(1, 2, 3), "|1 2 3|")
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

func TestStackTop(t *testing.T) {
	ConfirmTop := func(s *stackCell, r interface{}) {
		if x := s.Top(); x != r {
			t.Fatalf("%v.Top() should be %v but is %v", s, r, x)
		}
	}

	ConfirmTop(nil, nil)
	ConfirmTop(stack(), nil)
	ConfirmTop(stack(1), 1)
	ConfirmTop(stack(1, 2), 1)
}

func TestStackPop(t *testing.T) {
	ConfirmPop := func(s *stackCell, r interface{}, z *stackCell) {
		switch x, n := s.Pop(); {
		case x != r:
			t.Fatalf("%v.Pop() should be %v but is %v", s, r, x)
		case !n.Equal(z):
			t.Fatalf("%v.Pop() should be [%v, %v] but is [%v, %v]", s, r, z, x, n)
		}
	}

	ConfirmPop(nil, nil, nil)
	ConfirmPop(stack(), nil, nil)
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

func TestStackDup(t *testing.T) {
	ConfirmDup := func(s, r *stackCell) {
		if x := s.Dup(); !x.Equal(r) {
			t.Fatalf("%v.Dup() should be %v but is %v", s, r, x)
		}
	}

	ConfirmDup(nil, nil)
	ConfirmDup(stack(), stack())
	ConfirmDup(stack(1), stack(1, 1))
	ConfirmDup(stack(1, 2), stack(1, 1, 2))
}

func TestStackSwap(t *testing.T) {
	ConfirmSwap := func(s, r *stackCell) {
		if x := s.Swap(); !x.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", s, r, x)
		}
	}

	ConfirmSwap(nil, nil)
	ConfirmSwap(stack(), stack())
	ConfirmSwap(stack(1), stack(1))
	ConfirmSwap(stack(1, 2), stack(2, 1))
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

func TestStackEnd(t *testing.T) {
	ConfirmEnd := func(s, r *stackCell) {
		if x := s.end(); !x.Equal(r) {
			t.Fatalf("%v.end() should be %v but is %v", s, r, x)
		}
	}

	ConfirmEnd(nil, nil)
	ConfirmEnd(stack(), stack())
	ConfirmEnd(stack(1), stack(1))
	ConfirmEnd(stack(1, 2), stack(2))
	ConfirmEnd(stack(1, 2, 3), stack(3))
}

func TestStackPickCell(t *testing.T) {
	ConfirmPickCell := func(s *stackCell, n int, r *stackCell) {
		if x := s.pickCell(n); !x.Equal(r) {
			t.Fatalf("%v.pickCell(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmPickCell(nil, 0, nil)
	ConfirmPickCell(nil, 1, nil)

	ConfirmPickCell(stack(), 0, stack())
	ConfirmPickCell(stack(), 1, nil)

	ConfirmPickCell(stack(0), 0, stack(0))
	ConfirmPickCell(stack(0), 1, nil)

	ConfirmPickCell(stack(0, 1), 0, stack(0, 1))
	ConfirmPickCell(stack(0, 1), 1, stack(1))
	ConfirmPickCell(stack(0, 1), 2, nil)
}

func TestStackPick(t *testing.T) {
	ConfirmPick := func(s *stackCell, n int, r *stackCell) {
		if x := s.Pick(n); !x.Equal(r) {
			t.Fatalf("%v.Pick(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmPick(nil, 0, nil)
	ConfirmPick(nil, 1, nil)

	ConfirmPick(stack(), 0, stack())
	ConfirmPick(stack(), 1, nil)

	ConfirmPick(stack(0), 0, stack(0, 0))
	ConfirmPick(stack(0), 1, stack(0))

	ConfirmPick(stack(0, 1), 0, stack(0, 0, 1))
	ConfirmPick(stack(0, 1), 1, stack(1, 0, 1))
	ConfirmPick(stack(0, 1), 2, stack(0, 1))

	ConfirmPick(stack(0, 1, 2), 0, stack(0, 0, 1, 2))
	ConfirmPick(stack(0, 1, 2), 1, stack(1, 0, 1, 2))
	ConfirmPick(stack(0, 1, 2), 2, stack(2, 0, 1, 2))
	ConfirmPick(stack(0, 1, 2), 3, stack(0, 1, 2))
}

func TestStackRoll(t *testing.T) {
	ConfirmRoll := func(s *stackCell, n int, r *stackCell) {
		if x := s.Roll(n); !x.Equal(r) {
			t.Fatalf("%v.Roll(%v) should be %v but is %v", s, n, r, x)
		}
	}

	ConfirmRoll(nil, 0, nil)
	ConfirmRoll(nil, 1, nil)

	ConfirmRoll(stack(), 0, stack())
	ConfirmRoll(stack(), 1, stack())

	ConfirmRoll(stack(0), 0, stack(0))
	ConfirmRoll(stack(0), 1, stack(0))

	ConfirmRoll(stack(0, 1), 0, stack(0, 1))
	ConfirmRoll(stack(0, 1), 1, stack(1, 0))
	ConfirmRoll(stack(0, 1), 2, stack(0, 1))

	ConfirmRoll(stack(0, 1, 2), 0, stack(0, 1, 2))
	ConfirmRoll(stack(0, 1, 2), 1, stack(1, 0, 2))
	ConfirmRoll(stack(0, 1, 2), 2, stack(2, 0, 1))
	ConfirmRoll(stack(0, 1, 2), 3, stack(0, 1, 2))

	ConfirmRoll(stack(0, 1, 2, 3), 0, stack(0, 1, 2, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 1, stack(1, 0, 2, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 2, stack(2, 0, 1, 3))
	ConfirmRoll(stack(0, 1, 2, 3), 3, stack(3, 0, 1, 2))
	ConfirmRoll(stack(0, 1, 2, 3), 4, stack(0, 1, 2, 3))
}