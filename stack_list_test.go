package greenspun

import (
	"fmt"
	"testing"
)

func TestStackListString(t *testing.T) {
	ConfirmString := func(s *StackList, r string) {
		if s.String() != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(nil, "0:<]")
	ConfirmString(new(StackList), "0:<]")
	ConfirmString(&StackList{ depth: 1, stackCell: stack(1) }, "1:<1]")
	ConfirmString(&StackList{ depth: 2, stackCell: stack(1, 2) }, "2:<1 2]")
	ConfirmString(&StackList{ depth: 3, stackCell: stack(1, 2, 3) }, "3:<1 2 3]")
	ConfirmString(Stack(1), "1:<1]")
	ConfirmString(Stack(1, 2), "2:<1 2]")
	ConfirmString(Stack(1, 2, 3), "3:<1 2 3]")
}

func TestStackListEqual(t *testing.T) {
	ConfirmEqual := func(x *StackList, y Equatable, r bool) {
		if z := x.Equal(y); z != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", x, y, r, z)
		}
	}

	ConfirmEqual(nil, (*StackList)(nil), true)
	ConfirmEqual(Stack(1), (*StackList)(nil), false)
	ConfirmEqual(nil, Stack(), true)
	ConfirmEqual(Stack(), nil, true)
	ConfirmEqual(nil, Stack(1), false)
	ConfirmEqual(Stack(1), Stack(1), true)

	ConfirmEqual(Stack(1, 2), (*StackList)(nil), false)
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

func TestStackListPush(t *testing.T) {
	ConfirmPush := func(s *StackList, v interface{}, r *StackList) {
		vs := fmt.Sprintf("%v", s)
		if x := s.Push(v); !x.Equal(r) {
			t.Fatalf("%v.Push(%v) should be %v but is %v", vs, v, r, x)
		}
	}

	ConfirmPush(nil, nil, Stack(nil))
	ConfirmPush(nil, 1, Stack(1))
	ConfirmPush(Stack(1), 1, Stack(1, 1))
	ConfirmPush(Stack(1, 2), 1, Stack(1, 1, 2))
}

func TestStackListTop(t *testing.T) {
	ConfirmTop := func(s *StackList, r interface{}) {
		if x := s.Top(); x != r {
			t.Fatalf("%v.Top() should be %v but is %v", s, r, x)
		}
	}

	ConfirmTop(nil, nil)
	ConfirmTop(Stack(), nil)
	ConfirmTop(Stack(0), 0)
	ConfirmTop(Stack(1, 0), 1)
	ConfirmTop(Stack(2, 1, 0), 2)
}

func TestStackListPop(t *testing.T) {
	ConfirmPop := func(s *StackList, r interface{}, n *StackList) {
		vs := s.String()
		switch x := s.Pop(); {
		case x != r:
			t.Fatalf("%v.Pop() should be %v but is %v", vs, r, x)
		case !s.Equal(n):
			t.Fatalf("%v.Pop() should leave %v but leaves %v", vs, n, s)
		}
	}

	ConfirmPop(nil, nil, nil)
	ConfirmPop(Stack(), nil, nil)
	ConfirmPop(Stack(), nil, Stack())
	ConfirmPop(Stack(0), 0, Stack())
	ConfirmPop(Stack(1, 0), 1, Stack(0))
	ConfirmPop(Stack(2, 1, 0), 2, Stack(1, 0))
}

func TestStackListLen(t * testing.T) {
	ConfirmLen := func(s *StackList, r int) {
		if l := s.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", s, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(Stack(), 0)
	ConfirmLen(Stack(0), 1)
	ConfirmLen(Stack(1, 0), 2)
}

func TestStackListDrop(t *testing.T) {
	ConfirmDrop := func(s, r *StackList) {
		vs := s.String()
		if s.Drop(); !s.Equal(r) {
			t.Fatalf("%v.Drop() should leave %v but leaves %v", vs, r, s)
		}
	}

	ConfirmDrop(nil, nil)
	ConfirmDrop(Stack(), nil)
	ConfirmDrop(Stack(), Stack())
	ConfirmDrop(Stack(0), Stack())
	ConfirmDrop(Stack(1, 0), Stack(0))
	ConfirmDrop(Stack(2, 1, 0), Stack(1, 0))
}

func TestStackListSwap(t *testing.T) {
	ConfirmSwap := func(s, r *StackList) {
		vs := s.String()
		if s.Swap(); !s.Equal(r) {
			t.Fatalf("%v.Swap() should be %v but is %v", vs, r, s)
		}
	}

	ConfirmSwap(nil, nil)
	ConfirmSwap(Stack(), Stack())
	ConfirmSwap(Stack(0), Stack(0))
	ConfirmSwap(Stack(1, 0), Stack(0, 1))
	ConfirmSwap(Stack(2, 1, 0), Stack(1, 2, 0))
}

func TestStackListRoll(t *testing.T) {
	ConfirmRoll := func(s *StackList, n int, r *StackList) {
		vs := s.String()
		if s.Roll(n); !s.Equal(r) {
			t.Fatalf("%v.Roll(%v) should be %v but is %v", vs, n, r, s)
		}
	}

	ConfirmRoll(nil, 0, nil)
	ConfirmRoll(nil, 1, nil)
	ConfirmRoll(nil, 2, nil)

	ConfirmRoll(Stack(), 0, Stack())
	ConfirmRoll(Stack(), 1, Stack())
	ConfirmRoll(Stack(), 2, Stack())

	ConfirmRoll(Stack(0), 0, Stack(0))
	ConfirmRoll(Stack(0), 1, Stack(0))
	ConfirmRoll(Stack(0), 2, Stack(0))

	ConfirmRoll(Stack(1, 0), 0, Stack(1, 0))
	ConfirmRoll(Stack(1, 0), 1, Stack(0, 1))
	ConfirmRoll(Stack(1, 0), 2, Stack(1, 0))

	ConfirmRoll(Stack(2, 1, 0), 0, Stack(2, 1, 0))
	ConfirmRoll(Stack(2, 1, 0), 1, Stack(1, 2, 0))
	ConfirmRoll(Stack(2, 1, 0), 2, Stack(0, 2, 1))
	ConfirmRoll(Stack(2, 1, 0), 3, Stack(2, 1, 0))
}

func TestStackListReplace(t *testing.T) {
	ConfirmReplace := func(s *StackList, n int, v interface{}, r *StackList) {
		vs := s.String()
		if s.Replace(n, v); !s.Equal(r) {
			t.Fatalf("%v.Roll(%v, %v) should be %v but is", vs, n, v, r, s)
		}
	}

	ConfirmReplace(nil, 0, 0, nil)
	ConfirmReplace(nil, 1, 0, nil)

	ConfirmReplace(Stack(), 0, 0, nil)
	ConfirmReplace(Stack(), 1, 0, nil)

	ConfirmReplace(Stack(0), 0, 1, Stack(1))
	ConfirmReplace(Stack(0), 1, 1, Stack(0))

	ConfirmReplace(Stack(1, 0), 0, 2, Stack(2, 0))
	ConfirmReplace(Stack(1, 0), 1, 2, Stack(1, 2))
	ConfirmReplace(Stack(1, 0), 2, 2, Stack(1, 0))
}