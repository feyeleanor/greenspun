package greenspun

import "testing"

func TestSliceSliceEqual(t *testing.T) {
	ConfirmEqual := func(a sliceSlice, v interface{}, r bool) {
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
	ConfirmEqual(sliceSlice{ 0: nil }, nil, false)
	ConfirmEqual(sliceSlice{ 0: nil }, sliceSlice{ 0: nil }, true)
	ConfirmEqual(sliceSlice{ 0: &versionedValue{ data: 1 } }, sliceSlice{ 0: nil }, false)
	ConfirmEqual(sliceSlice{ 0: &versionedValue{ data: nil } }, sliceSlice{ 0: nil }, true)
	ConfirmEqual(sliceSlice{ 0: &versionedValue{ data: 1 } }, sliceSlice{ 0: &versionedValue{ data: 1 } }, true)
	ConfirmEqual(sliceSlice{ 0: &versionedValue{ data: 1 } }, sliceSlice{ 1: &versionedValue{ data: 1 } }, false)
}