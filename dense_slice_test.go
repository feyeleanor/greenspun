package greenspun

import "testing"

func TestNewDenseSlice(t *testing.T) {
	ConfirmNewSlice := func(n int, v []interface{}, r *DenseSlice) {
		if x := NewDenseSlice(n, v...); !x.Equal(r) {
			t.Fatalf("NewDenseSlice(%v, %v) should be %v but is %v", n, v, r, x)
		}
	}

	ConfirmNewSlice(0, nil, &DenseSlice{ elements: make(sliceSlice, 0), currentVersion: 0 })

	ConfirmNewSlice(0, []interface{}{ &versionedValue{ data: 0 } },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 } }, currentVersion: 0 })

	ConfirmNewSlice(0, []interface{}{ &versionedValue{ data: 0 }, nil, nil, 3: &versionedValue{ data: 0 } },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 0 } }, currentVersion: 0 })

	ConfirmNewSlice(0, []interface{}{ 0, nil, nil, 1, nil, nil, nil, nil, nil, 0 },
									&DenseSlice{ elements: sliceSlice{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 1 }, 9: &versionedValue{ data: 0 } }, currentVersion: 0 })
}

func TestDenseSliceString(t *testing.T) {
	t.Fatalf("implement String()")
}

func TestDenseSliceLen(t *testing.T) {
	ConfirmLen := func(l *DenseSlice, r int) {
		if x := l.Len(); x != r {
			t.Fatalf("%v.Len() should be %v but is %v", l, r, x)
		}
	}

	ConfirmLen(NewDenseSlice(0), 0)
	ConfirmLen(NewDenseSlice(0, 0), 1)
	ConfirmLen(NewDenseSlice(0, 0, 1, 2, 3), 4)
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

	ConfirmEqual(NewDenseSlice(1, 0), NewDenseSlice(1, 0))
	RefuteEqual(NewDenseSlice(1, 1), NewDenseSlice(1, 0))
	RefuteEqual(NewDenseSlice(1, 0), NewDenseSlice(1, 1))

	ConfirmEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 1 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 }, &versionedValue{ data: 0 } } })
	RefuteEqual(&DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 0 }, &versionedValue{ data: 0 } } }, &DenseSlice{ elements: sliceSlice{ &versionedValue{ data: 1 }, &versionedValue{ data: 1 } } })

	ConfirmEqual(NewDenseSlice(2, 0, 0), NewDenseSlice(2, 0, 0))
	RefuteEqual(NewDenseSlice(2, 0, 0), NewDenseSlice(2, 0, 1))
	RefuteEqual(NewDenseSlice(2, 0, 0), NewDenseSlice(2, 1, 0))
	RefuteEqual(NewDenseSlice(2, 0, 0), NewDenseSlice(2, 1, 1))

	RefuteEqual(NewDenseSlice(2, 0, 1), NewDenseSlice(2, 0, 0))
	ConfirmEqual(NewDenseSlice(2, 0, 1), NewDenseSlice(2, 0, 1))
	RefuteEqual(NewDenseSlice(2, 0, 1), NewDenseSlice(2, 1, 0))
	RefuteEqual(NewDenseSlice(2, 0, 1), NewDenseSlice(2, 1, 1))

	RefuteEqual(NewDenseSlice(2, 1, 0), NewDenseSlice(2, 0, 0))
	RefuteEqual(NewDenseSlice(2, 1, 0), NewDenseSlice(2, 0, 1))
	ConfirmEqual(NewDenseSlice(2, 1, 0), NewDenseSlice(2, 1, 0))
	RefuteEqual(NewDenseSlice(2, 1, 0), NewDenseSlice(2, 1, 1))

	RefuteEqual(NewDenseSlice(2, 1, 1), NewDenseSlice(2, 0, 0))
	RefuteEqual(NewDenseSlice(2, 1, 1), NewDenseSlice(2, 0, 1))
	RefuteEqual(NewDenseSlice(2, 1, 1), NewDenseSlice(2, 1, 0))
	ConfirmEqual(NewDenseSlice(2, 1, 1), NewDenseSlice(2, 1, 1))
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
	ConfirmOutOfBounds(NewDenseSlice(3), -1, true)
	ConfirmOutOfBounds(NewDenseSlice(3), 0, false)
	ConfirmOutOfBounds(NewDenseSlice(3), 1, false)
	ConfirmOutOfBounds(NewDenseSlice(3), 2, false)
	ConfirmOutOfBounds(NewDenseSlice(3), 3, true)

	ConfirmAt := func(l *DenseSlice, i int, r interface{}) {
		if x := l.At(i); x != r {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, i, r, x)
		}
	}
	ConfirmAt(NewDenseSlice(3), 0, nil)
	ConfirmAt(NewDenseSlice(3), 1, nil)
	ConfirmAt(NewDenseSlice(3), 2, nil)

	ConfirmAt(NewDenseSlice(3, 1), 0, 1)
	ConfirmAt(NewDenseSlice(3, 1), 1, nil)
	ConfirmAt(NewDenseSlice(3, 1), 2, nil)

	ConfirmAt(NewDenseSlice(5, 0, 1, 2), 0, 0)
	ConfirmAt(NewDenseSlice(5, 0, 1, 2), 1, 1)
	ConfirmAt(NewDenseSlice(5, 0, 1, 2), 2, 2)
	ConfirmAt(NewDenseSlice(5, 0, 1, 2), 3, nil)
	ConfirmAt(NewDenseSlice(5, 0, 1, 2), 4, nil)

	elements := []interface{}{ 0: 2, 1: 0, 3: 1 }
	ConfirmAt(NewDenseSlice(5, elements...), 0, 2)
	ConfirmAt(NewDenseSlice(5, elements...), 1, 0)
	ConfirmAt(NewDenseSlice(5, elements...), 2, nil)
	ConfirmAt(NewDenseSlice(5, elements...), 3, 1)
	ConfirmAt(NewDenseSlice(5, elements...), 4, nil)
}