package greenspun

import(
	"testing"
)

func TestLen(t *testing.T) {
	ConfirmLen := func(v LispPair, r int) {
		if l := Len(v); l != r {
			t.Fatalf("Len(%v) should be %v but is %v", v, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(Cons(), 0)
	ConfirmLen(Cons(0), 1)
	ConfirmLen(Cons(0, 1), 2)
	ConfirmLen(Cons(0, 1, 2), 3)
}

func TestCar(t *testing.T) {
	ConfirmCar := func(v LispPair, r interface{}) {
		if car := Car(v); car != r {
			t.Fatalf("Car(%v) should be %v but is %v", v, r, car)
		}
	}

	ConfirmCar(nil, nil)
	ConfirmCar(Cons(0, 1), 0)
	ConfirmCar(Cons(1), 1)
}

func TestCdr(t *testing.T) {
	ConfirmCdr := func(v LispPair, r interface{}) {
		if cdr := Cdr(v); cdr != r {
			t.Fatalf("Cdr(%v) should be %v but is %v", v, r, cdr)
		}
	}

	ConfirmCdr(nil, nil)
	ConfirmCdr(Cons(0), nil)
	ConfirmCdr(Cons(0, 1), 1)
	ConfirmCdr(Cons(0, Cons(1)), Cons(1))
	ConfirmCdr(Cons(0, Cons(1, 2)), Cons(1, 2))
}

func TestCaar(t *testing.T) {
	ConfirmCaar := func(v LispPair, r interface{}) {
		if caar := Caar(v); caar != r {
			t.Fatalf("Caar(%v) should be %v but is %v", v, r, caar)
		}
	}

	ConfirmCaar(nil, nil)
	ConfirmCaar(Cons(0), nil)
	ConfirmCaar(Cons(0, 1), nil)
	ConfirmCaar(Cons(Cons(0, 1)), 0)
}

func TestCadr(t *testing.T) {
	ConfirmCadr := func(v LispPair, r interface{}) {
		if cadr := Cadr(v); cadr != r {
			t.Fatalf("Cadr(%v) should be %v but is %v", v, r, cadr)
		}
	}

	ConfirmCadr(nil, nil)
	ConfirmCadr(Cons(0), nil)
	ConfirmCadr(Cons(0, 1), nil)
	ConfirmCadr(Cons(Cons(0, 1)), 1)
}

func TestCdar(t *testing.T) {
	ConfirmCdar := func(v LispPair, r interface{}) {
		if cdar := Cdar(v); cdar != r {
			t.Fatalf("Cdar(%v) should be %v but is %v", v, r, cdar)
		}
	}

	ConfirmCdar(nil, nil)
	ConfirmCdar(Cons(0), nil)
	ConfirmCdar(Cons(0, Cons(1)), 1)
	ConfirmCdar(Cons(0, Cons(1, 2)), 1)
}

func TestCddr(t *testing.T) {
	ConfirmCddr := func(v LispPair, r interface{}) {
		if cddr := Cddr(v); cddr != r {
			t.Fatalf("Cddr(%v) should be %v but is %v", v, r, cddr)
		}
	}

	ConfirmCddr(nil, nil)
	ConfirmCddr(Cons(0), nil)
	ConfirmCddr(Cons(0, Cons(1)), nil)
	ConfirmCddr(Cons(0, Cons(1, 2)), 2)
}

func TestEnd(t *testing.T) {
	ConfirmEnd := func(v, o LispPair) {
		if end := End(v); !Equal(end, o) {
			t.Fatalf("End(%v) should be %v but is %v", v, o, end)
		}
	}

	ConfirmEnd(nil, nil)
	ConfirmEnd(Cons(0), Cons(0))
	ConfirmEnd(Cons(0, 1), Cons(0, 1))
	ConfirmEnd(Cons(0, 1, 2), Cons(1, 2))
}