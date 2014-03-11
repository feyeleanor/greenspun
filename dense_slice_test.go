package greenspun

import "testing"

func TestMakeDenseSlice(t *testing.T) {
	ConfirmNewSlice := func(n, c int) {
		switch x := MakeDenseSlice(n, c).elements; {
		case len(x) != n:
			t.Fatalf("MakeDenseSlice(%[1]v, %[2]v) should have %[1]v elements but has %[3]v elements", n, c, len(x))
		case cap(x) != c:
			t.Fatalf("MakeDenseSlice(%[1]v, %[2]v) should have %[2]v capacity but has %[3]v capacity", n, c, c, cap(x))
		}
	}

	ConfirmNewSlice(0, 0)
	ConfirmNewSlice(0, 1)
	ConfirmNewSlice(0, 2)
	ConfirmNewSlice(1, 1)
	ConfirmNewSlice(1, 2)
}

func TestNewDenseSlice(t *testing.T) {
	ConfirmNewSlice := func(v []interface{}, r *DenseSlice) {
		if x := NewDenseSlice(v...); !x.Equal(r) {
			t.Fatalf("NewDenseSlice(%v) should be %v but is %v", v, r, x)
		}
	}

	ConfirmNewSlice(nil, &DenseSlice{ elements: make(sliceSlice, 0), currentVersion: 0 })

	ConfirmNewSlice([]interface{}{ &versionedValue{ data: 0 } },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 } }, currentVersion: 0 })

	ConfirmNewSlice([]interface{}{ &versionedValue{ data: 0 }, nil, nil, 3: &versionedValue{ data: 0 } },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 0 } }, currentVersion: 0 })

	ConfirmNewSlice([]interface{}{ 0, nil, nil, 1, nil, nil, nil, nil, nil, 0 },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 1 }, 9: &versionedValue{ data: 0 } }, currentVersion: 0 })
}

func TestDenseSliceString(t *testing.T) {
	ConfirmString := func(l *DenseSlice, r string) {
		if x := l.String(); x != r {
			t.Fatalf("%v.String() should be %v", x, r)
		}
	}

	ConfirmString(nil, "<nil>")
	ConfirmString(NewDenseSlice(), "[]")
	ConfirmString(NewDenseSlice(0), "[0]")
	ConfirmString(NewDenseSlice(0, 1), "[0 1]")
	ConfirmString(NewDenseSlice(0, 1, 2), "[0 1 2]")
}

func TestDenseSliceLen(t *testing.T) {
	ConfirmLen := func(l *DenseSlice, r int) {
		if x := l.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", l, r, x)
		}
	}

	ConfirmLen(NewDenseSlice(), 0)
	ConfirmLen(NewDenseSlice(0), 1)
	ConfirmLen(NewDenseSlice(0, 1, 2), 3)
	ConfirmLen(NewDenseSlice(0, 1, 2, 3), 4)
}

func TestDenseSliceEqual(t *testing.T) {
	ConfirmEqual := func(l *DenseSlice, v interface{}) {
		if x := l.Equal(v); !x {
			t.Fatalf("%v.Equal(%v) should be true", l, v)
		}
	}

	RefuteEqual := func(l *DenseSlice, v interface{}) {
		if x := l.Equal(v); x {
			t.Fatalf("%v.Equal(%v) should be false", l, v)
		}
	}

	ConfirmEqual(nil, nil)
	RefuteEqual(nil, new(DenseSlice))
	RefuteEqual(new(DenseSlice), nil)
	ConfirmEqual(new(DenseSlice), new(DenseSlice))
	ConfirmEqual(&DenseSlice{}, new(DenseSlice))
	ConfirmEqual(new(DenseSlice), &DenseSlice{})

	ConfirmEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 } } })

	ConfirmEqual(NewDenseSlice(0), NewDenseSlice(0))
	RefuteEqual(NewDenseSlice(1), NewDenseSlice(0))
	RefuteEqual(NewDenseSlice(0), NewDenseSlice(1))
	ConfirmEqual(NewDenseSlice(1), NewDenseSlice(1))

	ConfirmEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 1 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 }, &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 }, &versionedValue{ data: 1 } } })

	ConfirmEqual(NewDenseSlice(0, 0), NewDenseSlice(0, 0))
	RefuteEqual(NewDenseSlice(0, 0), NewDenseSlice(0, 1))
	RefuteEqual(NewDenseSlice(0, 0), NewDenseSlice(1, 0))
	RefuteEqual(NewDenseSlice(0, 0), NewDenseSlice(1, 1))

	RefuteEqual(NewDenseSlice(0, 1), NewDenseSlice(0, 0))
	ConfirmEqual(NewDenseSlice(0, 1), NewDenseSlice(0, 1))
	RefuteEqual(NewDenseSlice(0, 1), NewDenseSlice(1, 0))
	RefuteEqual(NewDenseSlice(0, 1), NewDenseSlice(1, 1))

	RefuteEqual(NewDenseSlice(1, 0), NewDenseSlice(0, 0))
	RefuteEqual(NewDenseSlice(1, 0), NewDenseSlice(0, 1))
	ConfirmEqual(NewDenseSlice(1, 0), NewDenseSlice(1, 0))
	RefuteEqual(NewDenseSlice(1, 0), NewDenseSlice(1, 1))

	RefuteEqual(NewDenseSlice(1, 1), NewDenseSlice(0, 0))
	RefuteEqual(NewDenseSlice(1, 1), NewDenseSlice(0, 1))
	RefuteEqual(NewDenseSlice(1, 1), NewDenseSlice(1, 0))
	ConfirmEqual(NewDenseSlice(1, 1), NewDenseSlice(1, 1))
}

func TestDenseSliceAt(t *testing.T) {
	ConfirmOutOfBounds := func(l *DenseSlice, i int, r bool) {
		defer func() {
			if x := recover() == ARGUMENT_OUT_OF_BOUNDS; x != r {
				t.Fatalf("%v.At(%v) out of bounds should be %v but is %v", l, i, r, x)
			}
		}()
		l.At(i)
	}

	ConfirmOutOfBounds(nil, -1, true)
	ConfirmOutOfBounds(nil, 0, true)
	ConfirmOutOfBounds(nil, 1, true)

	elements := NewDenseSlice(0, 1, 2)
	ConfirmOutOfBounds(elements, -1, true)
	ConfirmOutOfBounds(elements, 0, false)
	ConfirmOutOfBounds(elements, 1, false)
	ConfirmOutOfBounds(elements, 2, false)
	ConfirmOutOfBounds(elements, 3, true)

	ConfirmAt := func(l *DenseSlice, i int, r interface{}) {
		if x := l.At(i); x != r {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, i, r, x)
		}
	}

	elements = NewDenseSlice([]interface{}{ 0, 1, 2, nil, nil }...)
	ConfirmAt(elements, 0, 0)
	ConfirmAt(elements, 1, 1)
	ConfirmAt(elements, 2, 2)
	ConfirmAt(elements, 3, nil)
	ConfirmAt(elements, 4, nil)

	elements = NewDenseSlice([]interface{}{ 0: 0, 4: 1 }...)
	ConfirmAt(elements, 0, 0)
	ConfirmAt(elements, 1, nil)
	ConfirmAt(elements, 2, nil)
	ConfirmAt(elements, 3, nil)
	ConfirmAt(elements, 4, 1)
}

func TestDenseSliceSet(t *testing.T) {
	ConfirmOutOfBounds := func(l *DenseSlice, i int, r bool) {
		defer func() {
			if x := recover() == ARGUMENT_OUT_OF_BOUNDS; x != r {
				t.Fatalf("%v.Set(%v, <nil>) out of bounds should be %v but is %v", l, i, r, x)
			}
		}()
		l.Set(i, nil)
	}

	ConfirmOutOfBounds(nil, -1, true)
	ConfirmOutOfBounds(nil, 0, false)
	ConfirmOutOfBounds(nil, 1, false)
	ConfirmOutOfBounds(NewDenseSlice(0, 1, 2), -1, true)
	ConfirmOutOfBounds(NewDenseSlice(0, 1, 2), 0, false)
	ConfirmOutOfBounds(NewDenseSlice(0, 1, 2), 1, false)
	ConfirmOutOfBounds(NewDenseSlice(0, 1, 2), 2, false)
	ConfirmOutOfBounds(NewDenseSlice(0, 1, 2), 3, false)

	ConfirmSet := func(l *DenseSlice, i int, v interface{}, r *DenseSlice) {
		if x := l.Set(i, v); !x.Equal(r) {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", l, i, v, r, x)
		}
	}
	ConfirmSet(MakeDenseSlice(3, 3), 0, -1, NewDenseSlice(-1, nil, nil))
	ConfirmSet(MakeDenseSlice(3, 3), 1, -1, NewDenseSlice(nil, -1, nil))
	ConfirmSet(MakeDenseSlice(3, 3), 2, -1, NewDenseSlice(nil, nil, -1))
}

func TestDenseSliceAppend(t *testing.T) {
	ConfirmAppend := func(l *DenseSlice, v []interface{}, r *DenseSlice) {
		if x := l.Append(v...); !x.Equal(r){
			t.Fatalf("%v.Append(%v) should be %v but is %v", l, v, r, x)
		}
	}

	RefuteAppend := func(l *DenseSlice, v []interface{}, r *DenseSlice) {
		if x := l.Append(v...); x.Equal(r){
			t.Fatalf("%v.Append(%v) should be %v but is %v", l, v, r, x)
		}
	}

	ConfirmAppend(nil, nil, NewDenseSlice())
	RefuteAppend(nil, nil, NewDenseSlice(0))

	RefuteAppend(nil, []interface{}{ 0 }, NewDenseSlice())
	ConfirmAppend(nil, []interface{}{ 0 }, NewDenseSlice(0))
	RefuteAppend(nil, []interface{}{ 0 }, NewDenseSlice(0, 0))

	RefuteAppend(nil, []interface{}{ 0, 0 }, NewDenseSlice())
	RefuteAppend(nil, []interface{}{ 0, 0 }, NewDenseSlice(0))
	ConfirmAppend(nil, []interface{}{ 0, 0 }, NewDenseSlice(0, 0))
	RefuteAppend(nil, []interface{}{ 0, 0 }, NewDenseSlice(0, 0, 0))

	ConfirmAppend(NewDenseSlice(0), []interface{}{ 0 }, NewDenseSlice(0, 0))
	ConfirmAppend(NewDenseSlice(0, 1), []interface{}{ 0 }, NewDenseSlice(0, 1, 0))
	ConfirmAppend(NewDenseSlice(0, 1), []interface{}{ 0, 1 }, NewDenseSlice(0, 1, 0, 1))
	ConfirmAppend(NewDenseSlice(0, 1), []interface{}{ 0, 1, 2 }, NewDenseSlice(0, 1, 0, 1, 2))
	ConfirmAppend(NewDenseSlice(0, 1, 2), []interface{}{ 0 }, NewDenseSlice(0, 1, 2, 0))
	ConfirmAppend(NewDenseSlice(0, 1, 2), []interface{}{ 0, 1 }, NewDenseSlice(0, 1, 2, 0, 1))
	ConfirmAppend(NewDenseSlice(0, 1, 2), []interface{}{ 0, 1, 2 }, NewDenseSlice(0, 1, 2, 0, 1, 2))
}

func TestDenseSliceEach(t *testing.T) {
	s := NewDenseSlice(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmEach := func(c *DenseSlice, f interface{}) {
		count = 0
		c.Each(f)
		if l := c.Len(); l != count {
			t.Fatalf("%v.Each() should have iterated %v times not %v times", c, l, count)
		}
	}

	ConfirmEach(s, func() { count++ })

	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("1: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	ConfirmEach(s, func(index int, i interface{}) {
		if i != index {
			t.Fatalf("2: %v.Each() element %v erroneously reported as %v", s, index, i)
		}
		count++
	})

	ConfirmEach(s, func(key, i interface{}) {
		if i.(int) != key.(int) {
			t.Fatalf("3: %v.Each() element %v erroneously reported as %v", s, key, i)
		}
		count++
	})

	s = &DenseSlice{}
	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("4: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})

	s = nil
	ConfirmEach(s, func(i interface{}) {
		if i != count {
			t.Fatalf("5: %v.Each() element %v erroneously reported as %v", s, count, i)
		}
		count++
	})
}
