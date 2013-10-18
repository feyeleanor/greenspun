package greenspun

import(
	"testing"
)

func TestcellString(t *testing.T) {
	ConfirmString := func(c *cell, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v", c, r)
		}
	}

	ConfirmString(Cons(nil, nil), "()")
	ConfirmString(Cons(0, nil), "(0 . <nil>)")
	ConfirmString(Cons(0, 1), "(0 . 1)")
	ConfirmString(Cons(0, Cons(1, 2)), "(0 (1 . 2))")
	ConfirmString(Cons(0, Cons(1, Cons(2, 3))), "(0 (1 (2 . 3)))")
}

func Testcellequal(t *testing.T) {
	ConfirmEqual := func(l, r cell, ok bool) {
		if x := l.equal(r); x != ok {
			t.Fatalf("%v.equal(%v) should be %v but is %v", l, r, ok, x)
		}
		if x := r.equal(l); x != ok {
			t.Fatalf("%v.equal(%v) should be %v but is %v", r, l, ok, x)
		}
	}

	ConfirmEqual(cell{ nil, nil }, cell{ nil, nil }, true)
	ConfirmEqual(cell{ 1, nil }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, 1 }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ 1, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ nil, 1 }, false)
	ConfirmEqual(cell{ 1, nil }, cell{ 1, nil }, true)
	ConfirmEqual(cell{ nil, 1 }, cell{ nil, 1 }, true)
	ConfirmEqual(cell{ 1, 1 }, cell{ 1, 1 }, true)

	ConfirmEqual(cell{ cell{ 1, 1 }, nil }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, cell{ 1, 1 } }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ cell{ 1, 1 }, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ nil, cell{ 1, 1 } }, false)
	ConfirmEqual(cell{ cell{ 1, 1 }, nil }, cell{ cell{ 1, 1 }, nil }, true)
	ConfirmEqual(cell{ nil, cell{ 1, 1 } }, cell{ nil, cell{ 1, 1 } }, true)
	ConfirmEqual(cell{ cell{ 1, 1 }, cell{ 1, 1 } }, cell{ cell{ 1, 1 }, cell{ 1, 1 } }, true)

	ConfirmEqual(cell{ &cell{ 1, 1 }, nil }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, &cell{ 1, 1 } }, cell{ nil, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ &cell{ 1, 1 }, nil }, false)
	ConfirmEqual(cell{ nil, nil }, cell{ nil, &cell{ 1, 1 } }, false)
	ConfirmEqual(cell{ &cell{ 1, 1 }, nil }, cell{ &cell{ 1, 1 }, nil }, true)
	ConfirmEqual(cell{ nil, &cell{ 1, 1 } }, cell{ nil, &cell{ 1, 1 } }, true)
	ConfirmEqual(cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, cell{ &cell{ 1, 1 }, &cell{ 1, 1 } }, true)
}

func TestcellEqual(t *testing.T) {
	ConfirmEqual := func(l *cell, r interface{}, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(Cons(nil, nil), cell{ nil, nil }, true)
	ConfirmEqual(Cons(nil, nil), &cell{ nil, nil }, true)
	ConfirmEqual(Cons(nil, nil), Cons(nil, nil), true)

	ConfirmEqual(Cons(1, nil), cell{ 1, nil }, true)
	ConfirmEqual(Cons(1, nil), &cell{ 1, nil }, true)
	ConfirmEqual(Cons(1, nil), Cons(1, nil), true)

	ConfirmEqual(Cons(nil, 1), cell{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), &cell{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), Cons(nil, 1), true)

	ConfirmEqual(Cons(1, nil), cell{ nil, nil }, false)
	ConfirmEqual(Cons(1, nil), &cell{ nil, nil }, false)
	ConfirmEqual(Cons(1, nil), Cons(nil, nil), false)

	ConfirmEqual(Cons(nil, 1), cell{ nil, nil }, false)
	ConfirmEqual(Cons(nil, 1), &cell{ nil, nil }, false)
	ConfirmEqual(Cons(nil, 1), Cons(nil, nil), false)

	ConfirmEqual(Cons(nil, 1), cell{ 1, nil }, false)
	ConfirmEqual(Cons(nil, 1), &cell{ 1, nil }, false)
	ConfirmEqual(Cons(nil, 1), Cons(1, nil), false)

	ConfirmEqual(Cons(1, nil), cell{ nil, 1 }, false)
	ConfirmEqual(Cons(1, nil), &cell{ nil, 1 }, false)
	ConfirmEqual(Cons(1, nil), Cons(nil, 1), false)

	ConfirmEqual(Cons(nil, 1), cell{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), &cell{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), Cons(nil, 1), true)

	ConfirmEqual(Cons(Cons(0, 1), 2), cell{ &cell{ 0, 1 }, 2 }, true)
	ConfirmEqual(Cons(Cons(0, 1), 2), &cell{ &cell{ 0, 1 }, 2 }, true)
	ConfirmEqual(Cons(Cons(0, 1), 2), Cons( &cell{ 0, 1 }, 2 ), true)
}

func TestcellCar(t *testing.T) {
	ConfirmCar := func(c *cell, r interface{}) {
		if car := c.Car(); !areEqual(car, r) {
			t.Fatalf("%v.Car() should be %v but is %v", c, r, car)
		}
	}

	ConfirmCar(nil, nil)
	ConfirmCar(Cons(0, nil), 0)
	ConfirmCar(Cons(1, 0), 1)
	ConfirmCar(Cons(List(1), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(2, 1), 0), Cons(2, 1))
	ConfirmCar(Cons(List(1, nil, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(List(1, nil, nil), 0), List(1, nil, nil))
}

func TestcellCdr(t *testing.T) {
	ConfirmCdr := func(c *cell, r interface{}) {
		if cdr := c.Cdr(); !areEqual(cdr, r) {
			t.Fatalf("%v.Cdr() should be %v but is %v", c, r, cdr)
		}
	}

	ConfirmCdr(nil, nil)
	ConfirmCdr(Cons(0, nil), nil)
	ConfirmCdr(Cons(0, 1), 1)
	ConfirmCdr(Cons(0, Cons(1, nil)), Cons(1, nil))
	ConfirmCdr(Cons(0, Cons(1, 2)), Cons(1, 2))
}

func TestcellRplaca(t *testing.T) {
	ConfirmRplaca := func(c *cell, v interface{}, r *cell) {
		cs := c.String()
		c.Rplaca(v)
		if x := c.Equal(r); !x {
			t.Fatalf("%v.Rplaca(%v) should be %v but is %v", cs, v, r, c)
		}
	}

	ConfirmRplaca(Cons(nil, nil), 1, Cons(1, nil))
	ConfirmRplaca(Cons(Cons(0, 1), 2), 1, Cons(1, 2))
}

func TestcellRplacd(t *testing.T) {
	ConfirmRplacd := func(c *cell, v interface{}, r *cell) {
		cs := c.String()
		c.Rplacd(v)
		if x := c.Equal(r); !x {
			t.Fatalf("%v.Rplacd(%v) should be %v but is %v", cs, v, r, c)
		}
	}

	ConfirmRplacd(Cons(nil, nil), 1, Cons(nil, 1))
	ConfirmRplacd(Cons(0, Cons(1, 2)), 1, Cons(0, 1))
	ConfirmRplacd(Cons(Cons(0, 1), 2), 1, Cons(Cons(0, 1), 1))
}