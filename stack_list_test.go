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

	ConfirmString(nil, "nil:||")
	ConfirmString(new(StackList), "0:||")
	ConfirmString(&StackList{ depth: 1, stackCell: stack(1) }, "1:|1|")
	ConfirmString(&StackList{ depth: 2, stackCell: stack(1, 2) }, "2:|1 2|")
	ConfirmString(&StackList{ depth: 3, stackCell: stack(1, 2, 3) }, "3:|1 2 3|")
	ConfirmString(Stack(1), "1:|1|")
	ConfirmString(Stack(1, 2), "2:|1 2|")
	ConfirmString(Stack(1, 2, 3), "3:|1 2 3|")
}

func TestStackListEqual(t *testing.T) {
	ConfirmEqual := func(x *StackList, y Equatable, r bool) {
		if z := x.Equal(y); z != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", x, y, r, z)
		}
	}

	ConfirmEqual(nil, (*StackList)(nil), true)
	ConfirmEqual(Stack(1), (*StackList)(nil), false)
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
