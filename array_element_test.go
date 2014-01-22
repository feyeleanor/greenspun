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