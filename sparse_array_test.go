package greenspun

import "testing"

func TestArrayElementEqual(t *testing.T) {
	ConfirmEqual := func(a *arrayElement, v interface{}, r bool) {
		if x := a.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
		}
		if v, ok := v.(*arrayElement); ok {
			if x := v.Equal(a); x != r {
				t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
			}
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(&arrayElement{ data: nil }, nil, true)
	ConfirmEqual(&arrayElement{ data: nil }, &arrayElement{ data: nil }, true)
	ConfirmEqual(&arrayElement{}, nil, true)
	ConfirmEqual(&arrayElement{}, &arrayElement{}, true)

	ConfirmEqual(&arrayElement{ data: 1 }, nil, false)
	ConfirmEqual(&arrayElement{ data: 1 }, &arrayElement{}, false)
	ConfirmEqual(&arrayElement{ data: 1 }, 1, true)
	ConfirmEqual(&arrayElement{ data: 1 }, &arrayElement{ data: 1 }, true)

	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, nil, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{}, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{ data: stack(0) }, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{ data: stack(0, 1) }, true)
}

func TestArrayCellsEqual(t *testing.T) {
	ConfirmEqual := func(a arrayCells, v interface{}, r bool) {
		if x := a.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
		}
		if v, ok := v.(*arrayElement); ok {
			if x := v.Equal(a); x != r {
				t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
			}
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(arrayCells{ 0: nil }, nil, false)
	ConfirmEqual(arrayCells{ 0: nil }, arrayCells{ 0: nil }, true)
	ConfirmEqual(arrayCells{ 0: &arrayElement{ data: 1 } }, arrayCells{ 0: nil }, false)
	ConfirmEqual(arrayCells{ 0: &arrayElement{ data: nil } }, arrayCells{ 0: nil }, true)
	ConfirmEqual(arrayCells{ 0: &arrayElement{ data: 1 } }, arrayCells{ 0: &arrayElement{ data: 1 } }, true)
	ConfirmEqual(arrayCells{ 0: &arrayElement{ data: 1 } }, arrayCells{ 1: &arrayElement{ data: 1 } }, false)
}

func TestNewSparseArray(t *testing.T) {
	ConfirmNewArray := func(n int, d interface{}, v []arrayCells, r *SparseArray) {
		if x := NewSparseArray(n, d, v...); !x.Equal(r) {
			t.Fatalf("NewSparseArray(%v, %v, %v) should be %v but is %v", n, d, v, r, x)
		}
	}

	ConfirmNewArray(0, nil, []arrayCells{}, &SparseArray{ elements: make(arrayCells), length: 0, version: 0, Default: nil })
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

	ConfirmEqual(	&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 1, elements:	arrayCells{ 1: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 1, elements:	arrayCells{ 1: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								&SparseArray{	length: 0 },
								false )

	ConfirmEqual(	&SparseArray{	length: 0 },
								&SparseArray{	length: 1, elements:	arrayCells{ 0: &arrayElement{ data: 0 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								true )

	ConfirmEqual(	&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 3: &arrayElement{ data: 1 } } },
								&SparseArray{	length: 2, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 } } },
								false )

	ConfirmEqual(	&SparseArray{	length: 3, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								&SparseArray{	length: 3, elements:	arrayCells{ 0: &arrayElement{ data: 0 }, 1: &arrayElement{ data: 1 }, 2: &arrayElement{ data: 2 } } },
								true )
}

func TestSparseArraySet(t *testing.T) {
}

func TestSparseArrayAt(t *testing.T) {
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