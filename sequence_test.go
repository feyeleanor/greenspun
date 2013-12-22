package greenspun

import "testing"

func TestMatchValue(t *testing.T) {
	ConfirmMatchValue := func(s, o Sequence, r bool) {
		if x := MatchValue(s, o); x != r {
			t.Fatalf("MatchValue(%v, %v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmMatchValue(nil, nil, true)
	ConfirmMatchValue(nil, stack(0), false)
	ConfirmMatchValue(stack(0), nil, false)
	ConfirmMatchValue(stack(0), stack(0), true)
	ConfirmMatchValue(stack(0, 1), stack(0, -1), true)

	ConfirmMatchValue(nil, List(0), false)
	ConfirmMatchValue(List(0), nil, false)
	ConfirmMatchValue(List(0), List(0), true)
	ConfirmMatchValue(List(0, 1), List(0, -1), true)
}