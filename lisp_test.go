package greenspun

import(
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	ConfirmLen := func(v LispPair, r int) {
		if l := Len(v); l != r {
			t.Fatalf("Len(%v) should be %v but is %v", v, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(&cell{}, 1)
	ConfirmLen(Cons(nil, nil), 1)
	ConfirmLen(List(), 0)
	ConfirmLen(Cons(0, nil), 1)
	ConfirmLen(List(0), 1)
	ConfirmLen(Cons(0, 1), 2)
	ConfirmLen(List(0, 1), 2)
	ConfirmLen(List(0, 1, 2), 3)
}

func TestIsList(t *testing.T) {
	ConfirmIsList := func(v interface{}, r bool) {
		if a := IsList(v); a != r {
			t.Fatalf("IsList(%v) should be %v [%T] but is %v [%T]", v, r, r, a, a)
		}
	}

	ConfirmIsList(nil, false)
	ConfirmIsList(1, false)
	ConfirmIsList([]int{}, false)
	ConfirmIsList(&cell{}, true)
	ConfirmIsList(Cons(nil, nil), true)
	ConfirmIsList(List(), true)
	ConfirmIsList(List(0), true)
}

func TestIsAtom(t *testing.T) {
	ConfirmIsAtom := func(v interface{}, r bool) {
		if a := IsAtom(v); a != r {
			t.Fatalf("IsAtom(%v) should be %v but is %v", v, r, a)
		}
	}

	ConfirmIsAtom(nil, true)
	ConfirmIsAtom(1, true)
	ConfirmIsAtom([]int{}, true)
	ConfirmIsAtom(&cell{}, false)
	ConfirmIsAtom(Cons(nil, nil), false)
	ConfirmIsAtom(List(), false)
	ConfirmIsAtom(List(0), false)
}

func TestCar(t *testing.T) {
	ConfirmCar := func(v LispPair, r interface{}) {
		if car := Car(v); !areEqual(car, r) {
			t.Fatalf("Car(%v) should be %v but is %v", v, r, car)
		}
	}

	ConfirmCar(nil, nil)
	ConfirmCar(&cell{ Head: 0 }, 0)
	ConfirmCar(Cons(0, nil), 0)
	ConfirmCar(Cons(1, 0), 1)
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(2, 1), 0), Cons(2, 1))
	ConfirmCar(Cons(List(1, nil, nil), 0), List(1, nil, nil))
}

func TestCdr(t *testing.T) {
	ConfirmCdr := func(v LispPair, r interface{}) {
		if cdr := Cdr(v); !areEqual(cdr, r) {
			t.Fatalf("Cdr(%v) should be %v but is %v", v, r, cdr)
		}
	}

	ConfirmCdr(nil, nil)
	ConfirmCdr(Cons(0, nil), nil)
	ConfirmCdr(Cons(0, 1), 1)
	ConfirmCdr(Cons(0, Cons(1, nil)), Cons(1, nil))
	ConfirmCdr(Cons(0, Cons(1, 2)), Cons(1, 2))
}

func TestCaar(t *testing.T) {
	ConfirmCaar := func(v LispPair, r interface{}) {
		if caar := Caar(v); caar != r {
			t.Fatalf("Caar(%v) should be %v but is %v", v, r, caar)
		}
	}

	ConfirmCaar(nil, nil)
	ConfirmCaar(Cons(0, nil), nil)
	ConfirmCaar(Cons(0, 1), nil)
	ConfirmCaar(Cons(Cons(0, 1), nil), 0)
}

func TestCadr(t *testing.T) {
	ConfirmCadr := func(v LispPair, r interface{}) {
		if cadr := Cadr(v); cadr != r {
			t.Fatalf("Cadr(%v) should be %v but is %v", v, r, cadr)
		}
	}

	ConfirmCadr(nil, nil)
	ConfirmCadr(Cons(0, nil), nil)
	ConfirmCadr(Cons(0, 1), nil)
	ConfirmCadr(Cons(Cons(0, 1), nil), 1)
}

func TestCdar(t *testing.T) {
	ConfirmCdar := func(v LispPair, r interface{}) {
		if cdar := Cdar(v); cdar != r {
			t.Fatalf("Cdar(%v) should be %v but is %v", v, r, cdar)
		}
	}

	ConfirmCdar(nil, nil)
	ConfirmCdar(Cons(0, nil), nil)
	ConfirmCdar(Cons(0, Cons(1, nil)), 1)
	ConfirmCdar(Cons(0, Cons(1, 2)), 1)
}

func TestCddr(t *testing.T) {
	ConfirmCddr := func(v LispPair, r interface{}) {
		if cddr := Cddr(v); cddr != r {
			t.Fatalf("Cddr(%v) should be %v but is %v", v, r, cddr)
		}
	}

	ConfirmCddr(nil, nil)
	ConfirmCddr(Cons(0, nil), nil)
	ConfirmCddr(Cons(0, Cons(1, nil)), nil)
	ConfirmCddr(Cons(0, Cons(1, 2)), 2)
}

func TestOffset(t *testing.T) {
	ConfirmOffset := func(v LispPair, i int, r LispPair) {
		if offset := Offset(v, i); !Equal(offset, r) {
			t.Fatalf("Offset(%v) should be %v but is %v", v, r, offset)
		}
	}

	ConfirmOffset(nil, -1, nil)
	ConfirmOffset(nil, 0, nil)
	ConfirmOffset(nil, 1, nil)

	ConfirmOffset(Cons(0, nil), -1, nil)
	ConfirmOffset(Cons(0, nil), 0, Cons(0, nil))
	ConfirmOffset(Cons(0, nil), 1, nil)

	ConfirmOffset(Cons(0, Cons(1, nil)), -1, nil)
	ConfirmOffset(Cons(0, Cons(1, nil)), 0, Cons(0, Cons(1, nil)))
	ConfirmOffset(Cons(0, Cons(1, nil)), 1, Cons(1, nil))
	ConfirmOffset(Cons(0, Cons(1, nil)), 2, nil)

	ConfirmOffset(Cons(0, Cons(1, 2)), -1, nil)
	ConfirmOffset(Cons(0, Cons(1, 2)), 0, Cons(0, Cons(1, 2)))
	ConfirmOffset(Cons(0, Cons(1, 2)), 1, Cons(1, 2))
	ConfirmOffset(Cons(0, Cons(1, 2)), 2, nil)
}

func TestEnd(t *testing.T) {
	ConfirmEnd := func(v, o LispPair) {
		if end := End(v); !Equal(end, o) {
			t.Fatalf("End(%v) should be %v but is %v", v, o, end)
		}
	}

	ConfirmEnd(nil, nil)
	ConfirmEnd(Cons(0, nil), Cons(0, nil))
	ConfirmEnd(Cons(0, 1), Cons(0, 1))
	ConfirmEnd(List(0, 1, 2), Cons(1, 2))
}

func TestAppend(t *testing.T) {
	ConfirmAppend := func(c LispPair, v interface{}, r interface{}) {
		cs := fmt.Sprintf("%v", c)
		if x := Append(c, v); !Equal(c, r) {
			t.Fatalf("%v.Append(%v) should have tail %v but has %v", cs, v, r, x)
		}
	}

	ConfirmAppend(List(), 1, 1)
	ConfirmAppend(List(), 1, List(1))
	ConfirmAppend(List(), List(1), List(1))
	ConfirmAppend(List(1), 2, 2)
	ConfirmAppend(List(1), 2, List(2))
	ConfirmAppend(List(1), List(2), List(2))
}

func TestEach(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmEach := func(c LispPair, f interface{}) {
		count = 0
		Each(c, f)
		if l := Len(c); l != count {
			t.Fatalf("Each(%v) should have iterated %v times not %v times", c, l, count)
		}
	}

	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("1: Each(%v) element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	ConfirmEach(list, func(index int, i interface{}) {
		if i != index {
			t.Fatalf("2: Each(%v) element %v erroneously reported as %v", list, index, i)
		}
		count++
	})

	ConfirmEach(list, func(key, i interface{}) {
		if i.(int) != key.(int) {
			t.Fatalf("3: Each(%v) element %v erroneously reported as %v", list, key, i)
		}
		count++
	})
}

func TestWhile(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c LispPair, l int, f interface{}) {
		if count := While(c, f); count != l {
			t.Fatalf("While(%v, %v) should have iterated %v times not %v times", c, l, l, count)
		}
	}

	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		return i != limit
	})

	limit = 6
	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	limit = 7
	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key != limit
	})
}

func TestUntil(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c LispPair, l int, f interface{}) {
		if count := Until(c, f); count != l {
			t.Fatalf("Until(%v, %v) should have iterated %v times not %v times", c, l, l, count)
		}
	}

	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		return i == limit
	})

	limit = 6
	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	limit = 7
	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key == limit
	})
}