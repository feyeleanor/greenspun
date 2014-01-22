package greenspun

import "testing"

func TestNewSparseArray(t *testing.T) {
	ConfirmNewArray := func(n int, d interface{}, v []arrayHash, r *SparseArray) {
		if x := NewSparseArray(n, d, v...); !x.Equal(r) {
			t.Fatalf("NewSparseArray(%v, %v, %v) should be %v but is %v", n, d, v, r, x)
		}
	}

	ConfirmNewArray(0, nil, []arrayHash{}, &SparseArray{ elements: make(arrayHash), length: 0, version: 0, Default: nil })

	ConfirmNewArray(0, nil, []arrayHash{ arrayHash{ 0: &arrayElement{ data: 0 } } },
									&SparseArray{ elements: arrayHash{ 0: &arrayElement{ data: 0 } }, length: 1, version: 0, Default: nil })

	ConfirmNewArray(0, 10, []arrayHash{ arrayHash{ 0: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 0 } } },
									&SparseArray{ elements: arrayHash{ 0: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 0 } } , length: 4, version: 0, Default: 10 })

	ConfirmNewArray(0, 10, []arrayHash{ arrayHash{ 0: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 0 } }, arrayHash{ 0: &arrayElement{ data: 1 }, 9: &arrayElement{ data: 0 } } },
									&SparseArray{ elements: arrayHash{ 0: &arrayElement{ data: 1 }, 3: &arrayElement{ data: 0 }, 9: &arrayElement{ data: 0 } } , length: 10, version: 0, Default: 10 })
}

func TestSparseArrayString(t *testing.T) {
}

func TestSparseArrayLen(t *testing.T) {
}

func TestSparseArrayEqual(t *testing.T) {
	ConfirmEqual := func(l *SparseArray, v interface{}, r bool) {
		if x := l.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, v, r, x)
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(nil, new(SparseArray), false)
	ConfirmEqual(new(SparseArray), nil, false)
	ConfirmEqual(new(SparseArray), new(SparseArray), true)
	ConfirmEqual(&SparseArray{}, new(SparseArray), true)
	ConfirmEqual(new(SparseArray), &SparseArray{}, true)

	ConfirmEqual(&SparseArray{ length: 1, Default: 0 }, &SparseArray{ length: 1, Default: 0 }, true)
	ConfirmEqual(&SparseArray{ length: 1, Default: 0 }, &SparseArray{ length: 1, Default: 1 }, false)

	ConfirmEqual(&SparseArray{ length: 2, Default: 0 }, &SparseArray{ length: 2, Default: 0 }, true)
	ConfirmEqual(&SparseArray{ length: 2, Default: 0 }, &SparseArray{ length: 2, Default: 1 }, false)

	ConfirmEqual(&SparseArray{ length: 3, Default: 0 }, &SparseArray{ length: 3, Default: 0 }, true)
	ConfirmEqual(&SparseArray{ length: 3, Default: 0 }, &SparseArray{ length: 3, Default: 1 }, false)

	ConfirmEqual(	&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 0 },
								false )

	ConfirmEqual(	&SparseArray{	length: 0 },
								&SparseArray{	length: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 3 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 3 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, Default: 1, elements:	arrayHash{ 0: &arrayElement{ data: 0 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 2, Default: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, Default: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 2, Default: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, Default: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, Default: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 2 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 3 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 3 } } },
								&SparseArray{	length: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 3, Default: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 3, Default: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, Default: 1, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 2: &arrayElement{ data: 2 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 3, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, Default: 2, elements: arrayHash{ 0: &arrayElement{ data: 0 }, 2: &arrayElement{ data: 2 } } },
								false )
}

func TestSparseArrayAt(t *testing.T) {
	ConfirmOutOfBounds := func(l *SparseArray, i int, r bool) {
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
	ConfirmOutOfBounds(NewSparseArray(3, nil), -1, true)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 0, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 1, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 2, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 3, true)

	ConfirmAt := func(l *SparseArray, i int, r interface{}) {
		if x := l.At(i); x != r {
			t.Fatalf("%v.At(%v) should be %v but is %v", l, i, r, x)
		}
	}
	ConfirmAt(NewSparseArray(3, nil), 0, nil)
	ConfirmAt(NewSparseArray(3, nil), 1, nil)
	ConfirmAt(NewSparseArray(3, nil), 2, nil)
	ConfirmAt(NewSparseArray(3, 1), 0, 1)
	ConfirmAt(NewSparseArray(3, 1), 1, 1)
	ConfirmAt(NewSparseArray(3, 1), 2, 1)

	ConfirmAt(NewSparseArray(5, -1, denseArrayHash(0, 1, 2)), 0, 0)
	ConfirmAt(NewSparseArray(5, -1, denseArrayHash(0, 1, 2)), 1, 1)
	ConfirmAt(NewSparseArray(5, -1, denseArrayHash(0, 1, 2)), 2, 2)
	ConfirmAt(NewSparseArray(5, -1, denseArrayHash(0, 1, 2)), 3, -1)
	ConfirmAt(NewSparseArray(5, -1, denseArrayHash(0, 1, 2)), 4, -1)

	elements := arrayHash{ 0: &arrayElement{ data: 2 }, 1: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 1 } }
	ConfirmAt(NewSparseArray(5, -1, elements), 0, 2)
	ConfirmAt(NewSparseArray(5, -1, elements), 1, 0)
	ConfirmAt(NewSparseArray(5, -1, elements), 2, -1)
	ConfirmAt(NewSparseArray(5, -1, elements), 3, 1)
	ConfirmAt(NewSparseArray(5, -1, elements), 4, -1)
}

func TestSparseArraySet(t *testing.T) {
	ConfirmOutOfBounds := func(l *SparseArray, i int, r bool) {
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
	ConfirmOutOfBounds(NewSparseArray(3, nil), -1, true)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 0, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 1, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 2, false)
	ConfirmOutOfBounds(NewSparseArray(3, nil), 3, false)

	ConfirmSet := func(l *SparseArray, i int, v interface{}, r *SparseArray) {
		if x := l.Set(i, v); !x.Equal(r) {
			t.Fatalf("%v.Set(%v, %v) should be %v but is %v", l, i, v, r, x)
		}
	}
	ConfirmSet(NewSparseArray(3, nil), 0, -1, NewSparseArray(3, nil, arrayHash{ 0: &arrayElement{ data: -1 } }))
	ConfirmSet(NewSparseArray(3, nil), 1, -1, NewSparseArray(3, nil, arrayHash{ 1: &arrayElement{ data: -1 } }))
	ConfirmSet(NewSparseArray(3, nil), 2, -1, NewSparseArray(3, nil, arrayHash{ 2: &arrayElement{ data: -1 } }))
	ConfirmSet(NewSparseArray(3, nil), 3, -1, NewSparseArray(4, nil, arrayHash{ 3: &arrayElement{ data: -1 } }))
	ConfirmSet(NewSparseArray(3, nil), 4, -1, NewSparseArray(5, nil, arrayHash{ 4: &arrayElement{ data: -1 } }))

	ConfirmSet(NewSparseArray(3, nil), 3, nil, NewSparseArray(4, nil))
	ConfirmSet(NewSparseArray(3, nil), 4, nil, NewSparseArray(5, nil))
}

func TestSparseArrayEach(t *testing.T) {
	s := NewSparseArray(10, nil, denseArrayHash(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	count := 0

	ConfirmEach := func(c *SparseArray, f interface{}) {
		count = 0
		c.Each(f)
		if l := c.Len(); l != count {
			t.Fatalf("%v.Each() should have iterated %v times not %v times", c, l, count)
		}
	}

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

	s = &SparseArray{}
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

func TestSparseArrayInsert(t *testing.T) {
}

func TestSparseArrayDelete(t *testing.T) {
}

func TestSparseArrayCopy(t *testing.T) {
}

func TestSparseArrayCommit(t *testing.T) {
}

func TestSparseArrayRevert(t *testing.T) {
}