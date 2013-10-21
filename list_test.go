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

	ConfirmString(List(0, 1, 2), "(0 1 2)")
	ConfirmString(List(0, 1, 2, 3), "(0 1 2 3)")
	ConfirmString(List(0, List(1, 2, 3), List(2, 3), 3), "(0 (1 2 3) (2 3) 3)")
}

func TestListequal(t *testing.T) {
	ConfirmEqual := func(l *cell, r *cell, ok bool) {
		if x := l.equal(r); x != ok {
			t.Fatalf("1: %v.equal(%v) should be %v but is %v", l, r, ok, x)
		}
		if x := r.equal(l); x != ok {
			t.Fatalf("2: %v.equal(%v) should be %v but is %v", r, l, ok, x)
		}
	}

	ConfirmEqual(List(1), &cell{ nil, nil }, false)
	ConfirmEqual(List(1), &cell{ 1, nil }, true)
	ConfirmEqual(List(nil, 1), &cell{ 1, nil }, false)
	ConfirmEqual(List(nil, 1), &cell{ nil, 1 }, false)
	ConfirmEqual(List(nil, 1), Cons(nil, 1), false)
	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil} }, true)

	ConfirmEqual(List(nil, 1), &cell{ nil, 1 }, false)
	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil } }, true)

	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(1, 1), &cell{ 1, &cell{ 1, nil } }, true)

	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, &cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, cell{ 1, 1 } }, &cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, &cell{ &cell{ 1, 1 }, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, &cell{ nil, &cell{ 1, 1 } }, false)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, &cell{ &cell{ 1, 1 }, nil }, true)
	ConfirmEqual(&cell{ nil, &cell{ 1, 1 } }, &cell{ nil, &cell{ 1, 1 } }, true)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, &cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, true)

	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, &cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, &cell{ 1, 1 } }, &cell{ nil, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, &cell{ &cell{ 1, 1 }, nil }, false)
	ConfirmEqual(&cell{ nil, nil }, &cell{ nil, &cell{ 1, 1 } }, false)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, nil }, &cell{ &cell{ 1, 1 }, nil }, true)
	ConfirmEqual(&cell{ nil, &cell{ 1, 1 } }, &cell{ nil, &cell{ 1, 1 } }, true)
	ConfirmEqual(&cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, &cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, true)
}

func TestListEqual(t *testing.T) {
	ConfirmEqual := func(l *cell, r interface{}, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(List(1), &cell{ 1, nil }, true)
	ConfirmEqual(List(1), Cons(1, nil), true)

	ConfirmEqual(List(1), List(), false)
	ConfirmEqual(List(1), nil, false)
	ConfirmEqual(List(1), Cons(nil, nil), false)

	ConfirmEqual(List(nil, 1), &cell{ nil, &cell{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), Cons(nil, Cons(1, nil)), true)

	ConfirmEqual(List(nil, 1), &cell{ nil, nil }, false)
	ConfirmEqual(List(nil, 1), Cons(nil, nil), false)

	ConfirmEqual(List(Cons(0, 1), 2), &cell{ &cell{ 0, 1 }, 2 }, false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(&cell{ 0, 1 }, 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), List(Cons(0, 1), 2), true)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), Cons(2, nil)), true)
}

func TestListEach(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmEach := func(c *cell, f interface{}) {
		count = 0
		c.Each(f)
		if l := Len(c); l != count {
			t.Fatalf("%v.Each() should have iterated %v times not %v times", c, l, count)
		}
	}

	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("1: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	ConfirmEach(list, func(index int, i interface{}) {
		if i != index {
			t.Fatalf("2: %v.Each() element %v erroneously reported as %v", list, index, i)
		}
		count++
	})

	ConfirmEach(list, func(key, i interface{}) {
		if i.(int) != key.(int) {
			t.Fatalf("3: %v.Each() element %v erroneously reported as %v", list, key, i)
		}
		count++
	})

	list = List()
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("4: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	list = nil
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("5: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})
}