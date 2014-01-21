package greenspun

import "testing"

func TestEqual(t *testing.T) {
	ConfirmEqual := func(lhs, rhs interface{}, r bool) {
		if x := Equal(lhs, rhs); x != r {
			t.Fatalf("Equal(%v, %v) should be %v but is %v", lhs, rhs, r, x)
		}
	}

	ConfirmEqual(0, 0, true)
	ConfirmEqual(0, 1, false)
	ConfirmEqual(1, 0, false)
	ConfirmEqual(1, 1, true)
}