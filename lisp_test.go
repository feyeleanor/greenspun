package greenspun

import (
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
	ConfirmLen(List(), 0)
	ConfirmLen(&cell{}, 1)
	ConfirmLen(Cons(nil, nil), 1)
	ConfirmLen(Cons(0, nil), 1)
	ConfirmLen(List(0), 1)
	ConfirmLen(Cons(0, 1), 1)
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
	ConfirmIsList(Cons(0, nil), true)
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

func TestIsNil(t *testing.T) {
	ConfirmIsNil := func(v interface{}, r bool) {
		if n := IsNil(v); n != r {
			t.Fatalf("IsNil(%v) should be %v but is %v", v, r, n)
		}
	}

	ConfirmIsNil(nil, true)
	ConfirmIsNil(Cons(nil, nil), false)
	ConfirmIsNil(Cons(1, nil), false)
	ConfirmIsNil(List(), true)
	ConfirmIsNil(List(nil), false)
}

func TestEqual(t *testing.T) {
	ConfirmEqual := func(v, o LispPair, r bool) {
		if a := Equal(v, o); a != r {
			t.Fatalf("1: Equal(%v, %v) should be %v but is %v\n", v, o, r, a)
		}
		if a := Equal(o, v); a != r {
			t.Fatalf("2: Equal(%v, %v) should be %v but is %v\n", o, v, r, a)
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(List(), nil, true)
	ConfirmEqual(Cons(nil, nil), nil, false)
	ConfirmEqual(Cons(nil, nil), Cons(nil, nil), true)
	ConfirmEqual(List(nil), Cons(nil, nil), true)
	ConfirmEqual(List(nil), List(nil), true)
	ConfirmEqual(List(1, 2), Cons(1, Cons(2, nil)), true)
	ConfirmEqual(List(1, 2, 3), Cons(1, Cons(2, Cons(3, nil))), true)
	ConfirmEqual(List(1, List(2, 3), 4), List(1, Cons(2, Cons(3, nil)), 4), true)
	ConfirmEqual(List(1, List(2, 3), 4), List(1, List(2, 3), 4), true)
	ConfirmEqual(List(List(1, 3, 5), List(2, 3), 4), List(List(1, 3, 5), List(2, 3), 4), true)

	t.Logf("add tests for circular lists")
}

func TestCar(t *testing.T) {
	ConfirmCar := func(v LispPair, r interface{}) {
		if car := Car(v); !Equal(car, r) {
			t.Fatalf("Car(%v) should be %v but is %v", v, r, car)
		}
	}

	ConfirmCar(nil, nil)
	ConfirmCar(&cell{ head: 0 }, 0)
	ConfirmCar(Cons(0, nil), 0)
	ConfirmCar(Cons(1, 0), 1)
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(2, 1), 0), Cons(2, 1))
	ConfirmCar(Cons(List(1, nil, nil), 0), List(1, nil, nil))
}

func TestCdr(t *testing.T) {
	ConfirmCdr := func(v LispPair, r interface{}) {
		if cdr := Cdr(v); !Equal(cdr, r) {
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
		end := End(v)
		if ok := Equal(end, o); !ok {
			t.Fatalf("End(%v) should be %v but is %v", v, o, end)
		}
	}

	ConfirmEnd(nil, nil)
	ConfirmEnd(Cons(0, nil), Cons(0, nil))
	ConfirmEnd(Cons(0, 1), Cons(0, 1))
	ConfirmEnd(List(0, 1, 2), Cons(2, nil))
	ConfirmEnd(List(0, 1, 2, 3), Cons(3, nil))
	ConfirmEnd(Cons(0, Cons(1, Cons(2, Cons(3, nil)))), Cons(3, nil))
}

func TestAppend(t *testing.T) {
	ConfirmAppend := func(c LispPair, v interface{}, r interface{}) {
		cs := fmt.Sprintf("%v", c)
		if x := Append(c, v); !Equal(x, r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", cs, v, r, x)
		}
	}

	ConfirmAppend(List(), 1, List(1))
	ConfirmAppend(List(), List(1), List(1))
	ConfirmAppend(List(), List(1, 2), List(1, 2))
	ConfirmAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmAppend(List(1), 2, List(1, 2))
	ConfirmAppend(List(1), List(2), List(1, 2))
	ConfirmAppend(List(1), List(2, 3), List(1, 2, 3))

	ConfirmMultipleAppend := func(c LispPair, r interface{}, v... interface{}) {
		call := fmt.Sprintf("%v.Append(%v)", c, v)
		if x := Append(c, v...); !Equal(x, r) {
			t.Fatalf("%v should be %v but is %v", call, r, x)
		}
	}

	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1), List(2, 3))
	ConfirmMultipleAppend(List(), List(1, 2, 3, 4), List(1, 2), 3, List(4))
	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmMultipleAppend(List(1), List(1, 2), 2)
	ConfirmMultipleAppend(List(1), List(1, 2, 3), 2, 3)
	ConfirmMultipleAppend(List(1), List(1, 2), List(2))
	ConfirmMultipleAppend(List(1), List(1, 2, 3), List(2, 3))
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

	list = List()
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("4: Each(%v) element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	list = nil
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("5: Each(%v) element %v erroneously reported as %v", list, count, i)
		}
		count++
	})
}

func TestMap(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	doubles := List(0, 2, 4, 6, 8, 10, 12, 14, 16, 18)
	ConfirmMap := func(c, r LispPair, f interface{}) {
		switch m := Map(c, f); {
		case Len(c) != Len(m):
			t.Fatalf("Map(%v) should have iterated %v times not %v times", c, Len(c), Len(m))
		case !Equal(m, r):
			t.Fatalf("Map(%v) should be %v but is %v", c, r, m)
		}
	}

	ConfirmMap(list, doubles, func(i interface{}) interface{} {
		return i.(int) * 2
	})

	ConfirmMap(list, doubles, func(index int, i interface{}) interface{} {
		return index + i.(int)
	})

	ConfirmMap(list, doubles, func(key, i interface{}) interface{} {
		return key.(int) + i.(int)
	})
}

func TestReduce(t *testing.T) {
	ConfirmReduce := func(c LispPair, r, seed, f interface{}) {
		if x := Reduce(c, seed, f); !Equal(x, r) {
			t.Fatalf("Reduce(%v, %v) should be %v but is %v", c, seed, r, x)
		}
	}

	ConfirmReduce(List(1, 2, 3), 6, 1, func(seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 6, 1, func(index int, seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 6, 1, func(key, seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(index int, seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(key, seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(index int, seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(key, seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
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