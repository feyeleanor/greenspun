package greenspun

import "testing"

func TestSliceHashEqual(t *testing.T) {
	ConfirmEqual := func(a sliceHash, v interface{}, r bool) {
		if x := a.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
		}
		if v, ok := v.(*versionedValue); ok {
			if x := v.Equal(a); x != r {
				t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
			}
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(sliceHash{ 0: nil }, nil, false)
	ConfirmEqual(sliceHash{ 0: nil }, sliceHash{ 0: nil }, true)
	ConfirmEqual(sliceHash{ 0: &versionedValue{ data: 1 } }, sliceHash{ 0: nil }, false)
	ConfirmEqual(sliceHash{ 0: &versionedValue{ data: nil } }, sliceHash{ 0: nil }, true)
	ConfirmEqual(sliceHash{ 0: &versionedValue{ data: 1 } }, sliceHash{ 0: &versionedValue{ data: 1 } }, true)
	ConfirmEqual(sliceHash{ 0: &versionedValue{ data: 1 } }, sliceHash{ 1: &versionedValue{ data: 1 } }, false)
}