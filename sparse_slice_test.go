package greenspun

import "testing"

func TestNewSparseSlice(t *testing.T) {
	ConfirmNewArray := func(n int, d interface{}, v []sliceHash, r *SparseSlice) {
		if x := NewSparseSlice(n, d, v...); !x.Equal(r) {
			t.Fatalf("NewSparseSlice(%v, %v, %v) should be %v but is %v", n, d, v, r, x)
		}
	}

	ConfirmNewArray(0, nil, []sliceHash{}, &SparseSlice{ elements: make(sliceHash), length: 0, currentVersion: 0, defaultValue: nil })

	ConfirmNewArray(0, nil, []sliceHash{ sliceHash{ 0: &versionedValue{ data: 0 } } },
									&SparseSlice{ elements: sliceHash{ 0: &versionedValue{ data: 0 } }, length: 1, currentVersion: 0, defaultValue: nil })

	ConfirmNewArray(0, 10, []sliceHash{ sliceHash{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 0 } } },
									&SparseSlice{ elements: sliceHash{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 0 } }, length: 4, currentVersion: 0, defaultValue: &versionedValue{ data: 10 } })

	ConfirmNewArray(0, 10, []sliceHash{ sliceHash{ 0: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 0 } }, sliceHash{ 0: &versionedValue{ data: 1 }, 9: &versionedValue{ data: 0 } } },
									&SparseSlice{ elements: sliceHash{ 0: &versionedValue{ data: 1 }, 3: &versionedValue{ data: 0 }, 9: &versionedValue{ data: 0 } }, length: 10, currentVersion: 0, defaultValue: &versionedValue{ data: 10 } })
}

func TestSparseSliceString(t *testing.T) {
}

func TestSparseSliceLen(t *testing.T) {
}

func TestSparseSliceEqual(t *testing.T) {
	ConfirmEqual := func(l *SparseSlice, v interface{}, r bool) {
		if x := l.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, v, r, x)
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(nil, new(SparseSlice), false)
	ConfirmEqual(new(SparseSlice), nil, false)
	ConfirmEqual(new(SparseSlice), new(SparseSlice), true)
	ConfirmEqual(&SparseSlice{}, new(SparseSlice), true)
	ConfirmEqual(new(SparseSlice), &SparseSlice{}, true)

	ConfirmEqual(&SparseSlice{ length: 1, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 1, defaultValue: &versionedValue{ data: 0 } }, true)
	ConfirmEqual(&SparseSlice{ length: 1, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 1, defaultValue: &versionedValue{ data: 1 } }, false)

	ConfirmEqual(&SparseSlice{ length: 2, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 2, defaultValue: &versionedValue{ data: 0 } }, true)
	ConfirmEqual(&SparseSlice{ length: 2, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 2, defaultValue: &versionedValue{ data: 1 } }, false)

	ConfirmEqual(&SparseSlice{ length: 3, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 3, defaultValue: &versionedValue{ data: 0 } }, true)
	ConfirmEqual(&SparseSlice{ length: 3, defaultValue: &versionedValue{ data: 0 } }, &SparseSlice{ length: 3, defaultValue: &versionedValue{ data: 1 } }, false)

	ConfirmEqual(	&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 0 },
								false )

	ConfirmEqual(	&SparseSlice{	length: 0 },
								&SparseSlice{	length: 1, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 3 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 3 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 1 }, elements:	sliceHash{ 0: &versionedValue{ data: 0 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 1 }, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 2 }, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 1 }, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 2 }, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, defaultValue: &versionedValue{ data: 1 }, elements: sliceHash{ 0: &versionedValue{ data: 0 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 2 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 3 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 3 } } },
								&SparseSlice{	length: 2, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 3, defaultValue: &versionedValue{ data: 1 }, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 2: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 3, defaultValue: &versionedValue{ data: 2 }, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 2: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								false )

	ConfirmEqual(	&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 3, defaultValue: &versionedValue{ data: 1 }, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 2: &versionedValue{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseSlice{	length: 3, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 1: &versionedValue{ data: 1 }, 2: &versionedValue{ data: 2 } } },
								&SparseSlice{	length: 3, defaultValue: &versionedValue{ data: 2 }, elements: sliceHash{ 0: &versionedValue{ data: 0 }, 2: &versionedValue{ data: 2 } } },
								false )
}

func TestSparseSliceAt(t *testing.T) {
	ConfirmOutOfBounds := func(l *SparseSlice, i int, r bool) {
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
	ConfirmOutOfBounds(NewSparseSlice(3, nil), -1, true)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 0, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 1, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 2, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 3, true)

	ConfirmAt := func(l *SparseSlice, i int, r interface{}) {
		if x := l.At(i); x != r {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, i, r, x)
		}
	}
	ConfirmAt(NewSparseSlice(3, nil), 0, nil)
	ConfirmAt(NewSparseSlice(3, nil), 1, nil)
	ConfirmAt(NewSparseSlice(3, nil), 2, nil)
	ConfirmAt(NewSparseSlice(3, 1), 0, 1)
	ConfirmAt(NewSparseSlice(3, 1), 1, 1)
	ConfirmAt(NewSparseSlice(3, 1), 2, 1)

	ConfirmAt(NewSparseSlice(5, -1, denseSliceHash(0, 1, 2)), 0, 0)
	ConfirmAt(NewSparseSlice(5, -1, denseSliceHash(0, 1, 2)), 1, 1)
	ConfirmAt(NewSparseSlice(5, -1, denseSliceHash(0, 1, 2)), 2, 2)
	ConfirmAt(NewSparseSlice(5, -1, denseSliceHash(0, 1, 2)), 3, -1)
	ConfirmAt(NewSparseSlice(5, -1, denseSliceHash(0, 1, 2)), 4, -1)

	elements := sliceHash{ 0: &versionedValue{ data: 2 }, 1: &versionedValue{ data: 0 }, 3: &versionedValue{ data: 1 } }
	ConfirmAt(NewSparseSlice(5, -1, elements), 0, 2)
	ConfirmAt(NewSparseSlice(5, -1, elements), 1, 0)
	ConfirmAt(NewSparseSlice(5, -1, elements), 2, -1)
	ConfirmAt(NewSparseSlice(5, -1, elements), 3, 1)
	ConfirmAt(NewSparseSlice(5, -1, elements), 4, -1)
}

func TestSparseSliceSet(t *testing.T) {
	ConfirmOutOfBounds := func(l *SparseSlice, i int, r bool) {
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
	ConfirmOutOfBounds(NewSparseSlice(3, nil), -1, true)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 0, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 1, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 2, false)
	ConfirmOutOfBounds(NewSparseSlice(3, nil), 3, false)

	ConfirmSet := func(l *SparseSlice, i int, v interface{}, r *SparseSlice) {
		if x := l.Set(i, v); !x.Equal(r) {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", l, i, v, r, x)
		}
	}
	ConfirmSet(NewSparseSlice(3, nil), 0, -1, NewSparseSlice(3, nil, sliceHash{ 0: &versionedValue{ data: -1 } }))
	ConfirmSet(NewSparseSlice(3, nil), 1, -1, NewSparseSlice(3, nil, sliceHash{ 1: &versionedValue{ data: -1 } }))
	ConfirmSet(NewSparseSlice(3, nil), 2, -1, NewSparseSlice(3, nil, sliceHash{ 2: &versionedValue{ data: -1 } }))
	ConfirmSet(NewSparseSlice(3, nil), 3, -1, NewSparseSlice(4, nil, sliceHash{ 3: &versionedValue{ data: -1 } }))
	ConfirmSet(NewSparseSlice(3, nil), 4, -1, NewSparseSlice(5, nil, sliceHash{ 4: &versionedValue{ data: -1 } }))

	ConfirmSet(NewSparseSlice(3, nil), 3, nil, NewSparseSlice(4, nil))
	ConfirmSet(NewSparseSlice(3, nil), 4, nil, NewSparseSlice(5, nil))
}

func TestSparseSliceEach(t *testing.T) {
	s := NewSparseSlice(10, nil, denseSliceHash(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	count := 0

	ConfirmEach := func(c *SparseSlice, f interface{}) {
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

	s = &SparseSlice{}
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

func TestSparseSliceMove(t *testing.T) {
}

func TestSparseSliceInsert(t *testing.T) {
	ConfirmInsert := func(l *SparseSlice, i int, v []interface{}, r *SparseSlice) {
		if x := l.Insert(i, v...); !r.Equal(x) {
			t.Fatalf("%v.Insert(%v, %v) should be %v but is %v", l, i, v, r, x)
		}
	}

	ConfirmInsert(nil, 0, []interface{}{ 0 }, NewSparseSlice(1, nil, denseSliceHash(0)))
	ConfirmInsert(nil, 1, []interface{}{ 0 }, NewSparseSlice(2, nil, denseSliceHash(nil, 0)))
	ConfirmInsert(nil, 2, []interface{}{ 0 }, NewSparseSlice(3, nil, denseSliceHash(nil, nil, 0)))

	ConfirmInsert(nil, 0, []interface{}{ 0, 1 }, NewSparseSlice(2, nil, denseSliceHash(0, 1)))
	ConfirmInsert(nil, 1, []interface{}{ 0, 1 }, NewSparseSlice(3, nil, denseSliceHash(nil, 0, 1)))
	ConfirmInsert(nil, 2, []interface{}{ 0, 1 }, NewSparseSlice(4, nil, denseSliceHash(nil, nil, 0, 1)))

	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 0, []interface{}{ 1 }, NewSparseSlice(2, 0, denseSliceHash(1, 0)))
	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 1, []interface{}{ 1 }, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 2, []interface{}{ 1 }, NewSparseSlice(3, 0, denseSliceHash(0, 0, 1)))

	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 0, []interface{}{ 1, 2 }, NewSparseSlice(3, 0, denseSliceHash(1, 2, 0)))
	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 1, []interface{}{ 1, 2 }, NewSparseSlice(3, 0, denseSliceHash(0, 1, 2)))
	ConfirmInsert(NewSparseSlice(1, 0, denseSliceHash(0)), 2, []interface{}{ 1, 2 }, NewSparseSlice(4, 0, denseSliceHash(0, 0, 1, 2)))

	ConfirmInsert(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 0, []interface{}{ 1, 2 }, NewSparseSlice(4, 0, denseSliceHash(1, 2, 0, 1)))
	ConfirmInsert(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 1, []interface{}{ 1, 2 }, NewSparseSlice(4, 0, denseSliceHash(0, 1, 2, 1)))
	ConfirmInsert(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 2, []interface{}{ 1, 2 }, NewSparseSlice(4, 0, denseSliceHash(0, 1, 1, 2)))
	ConfirmInsert(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 3, []interface{}{ 1, 2 }, NewSparseSlice(5, 0, denseSliceHash(0, 1, 0, 1, 2)))
}

func TestSparseSliceDelete(t *testing.T) {
	ConfirmDelete := func(l *SparseSlice, i, n int, r *SparseSlice) {
		if x := l.Delete(i, n); !r.Equal(x) {
			t.Fatalf("%v.Delete(%v, %v) should be %v but is %v", l, i, n, r, x)
		}
	}

	ConfirmDelete(nil, 0, 1, nil)
	ConfirmDelete(nil, 0, 2, nil)
	ConfirmDelete(nil, 1, 1, nil)
	ConfirmDelete(nil, 1, 2, nil)

	ConfirmDelete(NewSparseSlice(1, 0, denseSliceHash(0)), 0, 0, NewSparseSlice(0, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(1, 0, denseSliceHash(0)), 0, 1, NewSparseSlice(0, 0))

	ConfirmDelete(NewSparseSlice(1, 0, denseSliceHash(0)), 1, 0, NewSparseSlice(0, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(1, 0, denseSliceHash(0)), 1, 1, NewSparseSlice(0, 0, denseSliceHash(0)))

	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 0, 0, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 0, 1, NewSparseSlice(1, 0, denseSliceHash(1)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 0, 2, NewSparseSlice(0, 0))

	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 1, 0, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 1, 1, NewSparseSlice(1, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 1, 2, NewSparseSlice(1, 0, denseSliceHash(0)))

	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 2, 0, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 2, 1, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(2, 0, denseSliceHash(0, 1)), 2, 2, NewSparseSlice(2, 0, denseSliceHash(0, 1)))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 1, NewSparseSlice(2, 0, denseSliceHash(1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 2, NewSparseSlice(1, 0))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 3, NewSparseSlice(0, 0))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 1, NewSparseSlice(2, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 2, NewSparseSlice(1, 0))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 3, NewSparseSlice(1, 0))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 1, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 2, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 3, NewSparseSlice(2, 0, denseSliceHash(0, 1)))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 1, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 2, NewSparseSlice(3, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 3, NewSparseSlice(3, 0, denseSliceHash(0, 1)))


	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 1, NewSparseSlice(2, 0, denseSliceHash(1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 2, NewSparseSlice(1, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 0, 3, NewSparseSlice(0, 0))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 1, NewSparseSlice(2, 0, denseSliceHash(0, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 2, NewSparseSlice(1, 0, denseSliceHash(0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 1, 3, NewSparseSlice(1, 0, denseSliceHash(0)))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 1, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 2, NewSparseSlice(2, 0, denseSliceHash(0, 1)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 2, 3, NewSparseSlice(2, 0, denseSliceHash(0, 1)))

	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 0, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 1, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 2, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
	ConfirmDelete(NewSparseSlice(3, 0, denseSliceHash(0, 1)), 3, 3, NewSparseSlice(3, 0, denseSliceHash(0, 1, 0)))
}

func TestSparseSliceCopy(t *testing.T) {
	ConfirmCopy := func(l, r *SparseSlice) {
		x := l.Copy()
		if !r.Equal(x) {
			t.Fatalf("%v.Copy() should be %v but is %v", l, r, x)
		}
		if x != nil {
			if x.currentVersion != 0 {
				t.Fatalf("%v.Copy() currentVersion should be 0 but is %v", l, x.currentVersion)
			}
			for i, v := range x.elements {
				switch {
				case v.version != 0:
					t.Fatalf("%v.Copy()[%v] currentVersion should be 0 but is %v", l, i, v.version)
				case v.versionedValue != nil:
					t.Fatalf("%v.Copy()[%v] should be a terminal node but is", l, i, v.versionedValue)
				}
			}
		}
	}

	ConfirmCopy(nil, nil)

	ConfirmCopy(NewSparseSlice(3, 0, denseSliceHash(0, 1, 2)),
							NewSparseSlice(3, 0, denseSliceHash(0, 1, 2)))

	ConfirmCopy(NewSparseSlice(0, 0).Set(2, 1).Set(2, 3).Set(1, 4).Set(2, 5).Set(3, 2),
							NewSparseSlice(4, 0, denseSliceHash(0, 4, 5, 2)))
}

func TestSparseSliceCommit(t *testing.T) {
	ConfirmCommit := func(l, r *SparseSlice) {
		switch x := l.Commit(); {
		case !r.Equal(x):
			t.Fatalf("%v.Commit() should be %v but is %v", l, r, x)
		case x != nil && x.currentVersion != 0:
			t.Fatalf("%v.Commit() currentVersion should be 0 but is %v", l, x.currentVersion)
		}
	}

	ConfirmCommit(nil, nil)
	ConfirmCommit(NewSparseSlice(5, 0), NewSparseSlice(5, 0, denseSliceHash(0, 0, 0, 0, 0)))
	ConfirmCommit(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), NewSparseSlice(5, 0, denseSliceHash(0, 1, 2, 3, 4)))
}

func TestSparseSliceUndo(t *testing.T) {
	ConfirmUndo := func(l *SparseSlice, steps int, r *SparseSlice) {
		if x := l.Undo(steps); !r.Equal(x) {
			t.Fatalf("%v.Undo() should be %v but is %v", l, r, x)
		}
	}

	ConfirmUndo(nil, 0, nil)
	ConfirmUndo(NewSparseSlice(5, 0), 0, NewSparseSlice(5, 0, denseSliceHash(0, 0, 0, 0, 0)))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 0, NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 1, NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3))

	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 0, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2, 3, 4)))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 1, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2, 3)))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 2, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2)))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 3, NewSparseSlice(0, 0, denseSliceHash(0, 1)))
	ConfirmUndo(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 4, NewSparseSlice(0, 0))

	ConfirmUndo(NewSparseSlice(5, 0).Set(1, 1).Set(2, 2), 2, NewSparseSlice(0, 0, denseSliceHash(0, 0, 0, 0, 0)))
	ConfirmUndo(NewSparseSlice(5, 0).Set(1, 1).Set(2, 2), 1, NewSparseSlice(0, 0, denseSliceHash(0, 1, 0, 0, 0)))
	ConfirmUndo(NewSparseSlice(5, 0).Set(1, 1).Set(2, 2), 0, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2, 0, 0)))
}

func TestSparseSliceRollback(t *testing.T) {
	ConfirmRollback := func(l *SparseSlice, v int, r *SparseSlice) {
		switch x := l.Rollback(v); {
		case !r.Equal(x):
			t.Fatalf("%v.Rollback() should be %v but is %v", l, r, x)
		case x != nil && x.currentVersion != v:
			t.Fatalf("%[1]v.Rollback(%[2]v) currentVersion should be %[2]v but is %[3]v", l, v, x.currentVersion)
		}
	}

	ConfirmRollback(nil, 0, nil)
	ConfirmRollback(NewSparseSlice(5, 0), 0, NewSparseSlice(5, 0, denseSliceHash(0, 0, 0, 0, 0)))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 0, NewSparseSlice(0, 0))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 1, NewSparseSlice(0, 0).Set(1, 1))

	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 0, NewSparseSlice(0, 0))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 1, NewSparseSlice(0, 0, denseSliceHash(0, 1)))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 1, NewSparseSlice(0, 0, denseSliceHash(0, 1)))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 2, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2)))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 3, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2, 3)))
	ConfirmRollback(NewSparseSlice(0, 0).Set(1, 1).Set(2, 2).Set(3, 3).Set(4, 4), 4, NewSparseSlice(0, 0, denseSliceHash(0, 1, 2, 3, 4)))

	ConfirmRollback(NewSparseSlice(5, 0).Set(1, 1).Set(2, 2), 0, NewSparseSlice(0, 0, denseSliceHash(0, 0, 0, 0, 0)))
}