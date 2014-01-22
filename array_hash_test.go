package greenspun

import "testing"

func TestArrayHashEqual(t *testing.T) {
	ConfirmEqual := func(a arrayHash, v interface{}, r bool) {
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
	ConfirmEqual(arrayHash{ 0: nil }, nil, false)
	ConfirmEqual(arrayHash{ 0: nil }, arrayHash{ 0: nil }, true)
	ConfirmEqual(arrayHash{ 0: &arrayElement{ data: 1 } }, arrayHash{ 0: nil }, false)
	ConfirmEqual(arrayHash{ 0: &arrayElement{ data: nil } }, arrayHash{ 0: nil }, true)
	ConfirmEqual(arrayHash{ 0: &arrayElement{ data: 1 } }, arrayHash{ 0: &arrayElement{ data: 1 } }, true)
	ConfirmEqual(arrayHash{ 0: &arrayElement{ data: 1 } }, arrayHash{ 1: &arrayElement{ data: 1 } }, false)
}