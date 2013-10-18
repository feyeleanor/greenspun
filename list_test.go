package greenspun

import(
	"testing"
)

func TestListString(t *testing.T) {
	ConfirmString := func(c *cell, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v", c, r)
		}
	}

	ConfirmString(List(0, 1, 2), "(0 (1 (2 . <nil>)))")
	ConfirmString(List(0, 1, 2, 3), "(0 (1 (2 (3 . <nil>))))")
}

func TestListLen(t *testing.T) {
	ConfirmLen := func(c *cell, r int) {
		if l := c.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", c, r, l)
		}
	}

	ConfirmLen(Cons(nil, nil), 1)
	ConfirmLen(Cons(0, nil), 1)
	ConfirmLen(Cons(0, 1), 2)
	ConfirmLen(List(0, 1, 2), 3)
}

func TestListequal(t *testing.T) {
	ConfirmEqual := func(l *cell, r cell, ok bool) {
		if x := l.equal(r); x != ok {
			t.Fatalf("1: %v.equal(%v) should be %v but is %v", l, r, ok, x)
		}
		if x := r.equal(*l); x != ok {
			t.Fatalf("2: %v.equal(%v) should be %v but is %v", r, l, ok, x)
		}
	}

	ConfirmEqual(List(1), cell{ nil, nil }, false)
	ConfirmEqual(List(1), cell{ 1, nil }, true)
	ConfirmEqual(List(nil, 1), cell{ 1, nil }, false)
	ConfirmEqual(List(nil, 1), *List(nil, 1), true)
	ConfirmEqual(List(nil, 1), cell{ nil, 1 }, false)

	ConfirmEqual(List(nil, 1), cell{ nil, 1 }, false)
	ConfirmEqual(List(nil, 1), cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), cell{ nil, &cell{ 1, nil } }, true)

	ConfirmEqual(List(nil, 1), cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(1, 1), cell{ 1, &cell{ 1, nil } }, true)

	ConfirmEqual(&cell{ cell{ 1, 1 }, nil }, cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, cell{ 1, 1 } }, cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, cell{ cell{ 1, 1 }, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, cell{ nil, cell{ 1, 1 } }, false)
	ConfirmEqual(&cell{ cell{ 1, 1 }, nil }, cell{ cell{ 1, 1 }, nil }, true)
	ConfirmEqual(&cell{ nil, cell{ 1, 1 } }, cell{ nil, cell{ 1, 1 } }, true)
	ConfirmEqual(&cell{ cell{ 1, 1 }, cell{ 1, 1 } }, cell{ cell{ 1, 1 }, cell{ 1, 1 } }, true)

	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, &cell{ 1, 1 } }, cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, cell{ &cell{ 1, 1 }, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, cell{ nil, &cell{ 1, 1 } }, false)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, cell{ &cell{ 1, 1 }, nil }, true)
	ConfirmEqual(&cell{ nil, &cell{ 1, 1 } }, cell{ nil, &cell{ 1, 1 } }, true)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, true)
}

func TestListEqual(t *testing.T) {
	ConfirmEqual := func(l *cell, r interface{}, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(List(1), cell{ 1, nil }, true)
	ConfirmEqual(List(1), &cell{ 1, nil }, true)
	ConfirmEqual(List(1), Cons(1, nil), true)

	ConfirmEqual(List(1), List(), false)
	ConfirmEqual(List(1), nil, false)
	ConfirmEqual(List(1), Cons(nil, nil), false)

	ConfirmEqual(List(nil, 1), cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), Cons(nil, Cons(1, nil)), true)

	ConfirmEqual(List(nil, 1), cell{ nil, nil }, false)
	ConfirmEqual(List(nil, 1), &cell{ nil, nil }, false)
	ConfirmEqual(List(nil, 1), Cons(nil, nil), false)

	ConfirmEqual(List(Cons(0, 1), 2), cell{ &cell{ 0, 1 }, 2 }, false)
	ConfirmEqual(List(Cons(0, 1), 2), &cell{ &cell{ 0, 1 }, 2 }, false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(&cell{ 0, 1 }, 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), List(Cons(0, 1), 2), true)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), Cons(2, nil)), true)
}