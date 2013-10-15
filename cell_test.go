package greenspun

import(
	"testing"
)

func TestCellString(t *testing.T) {
	ConfirmString := func(c *Cell, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v", c, r)
		}
	}

	ConfirmString(nil, "()")
	ConfirmString(Cons(nil, nil), "()")
	ConfirmString(Cons(0, nil), "(0 . <nil>)")
	ConfirmString(Cons(0, 1), "(0 . 1)")
	ConfirmString(List(0, 1, 2), "(0 (1 . 2))")
	ConfirmString(List(0, 1, 2, 3), "(0 (1 (2 . 3)))")
}

func TestCellLen(t *testing.T) {
	ConfirmLen := func(c *Cell, r int) {
		if l := c.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", c, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(Cons(nil, nil), 0)
	ConfirmLen(Cons(0, nil), 1)
	ConfirmLen(Cons(0, 1), 2)
	ConfirmLen(List(0, 1, 2), 3)
}

func TestCellCar(t *testing.T) {
	ConfirmCar := func(c *Cell, r interface{}) {
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

func TestCellCdr(t *testing.T) {
	ConfirmCdr := func(c *Cell, r interface{}) {
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