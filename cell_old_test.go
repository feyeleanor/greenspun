package greenspun
/*
import (
	"testing"
	"reflect"
)

func TestcellEnd(t *testing.T) {
	ConfirmEnd := func(c *cell, r interface{}) {
		x := c.End()
		switch {
		case x == nil:
			t.Fatalf("%v.End() returned nil", c)
		case x.Head != r:
			t.Fatalf("%v.End() should be '%v' but is '%v'", c, r, x.Head)
		}
	}
	RefuteEnd := func(c *cell) {
		if x := c.End(); x != nil {
			t.Fatalf("%v.End() should be nil but is '%v'", c, x.Head)
		}
	}
	RefuteEnd(List())
	ConfirmEnd(List(0), 0)
	ConfirmEnd(List(0, 1), 1)
	ConfirmEnd(List(0, 1, 2), 2)
}

func TestcellOffset(t *testing.T) {
	ConfirmOffset := func(c *cell, i int, r interface{}) {
		if x := c.Offset(i); !x.Equal(r) {
			t.Fatalf("%v.Offset(%v) should be '%v' but is '%v'", c, i, r, x.Head)
		}
	}
	RefuteOffset := func(c *cell, i int) {
		if x := c.Offset(i); x != nil {
			t.Fatalf("%v.Offset(%v) should be nil but is %v of type %v", c, i, x, reflect.TypeOf(x))
		}
	}
	c := List(0, 1, 2, 3, 4)
	RefuteOffset(c, -1)
	ConfirmOffset(c, 0, 0)
	ConfirmOffset(c, 1, 1)
	ConfirmOffset(c, 2, 2)
	ConfirmOffset(c, 3, 3)
	ConfirmOffset(c, 4, 4)
	RefuteOffset(c, 5)
}

func TestcellEach(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0
	list.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	list.Each(func(index int, i interface{}) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	list.Each(func(key, i interface{}) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestcellWhile(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c *cell, l int, f interface{}) {
		if count, _ := c.While(f); count != l {
			t.Fatalf("%v.While() should have iterated %v times not %v times", c, l, count)
		}
	}

	count := 0
	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		if count == limit {
			return false
		}
		count++
		return true
	})

	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key.(int) != limit
	})
}

func TestcellUntil(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c *cell, l int, f interface{}) {
		if count, _ := c.Until(f); count != l {
			t.Fatalf("%v.Until() should have iterated %v times not %v times", c, l, count)
		}
	}

	count := 0
	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		if count == limit {
			return true
		}
		count++
		return false
	})

	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key.(int) == limit
	})
}

func TestcellLen(t *testing.T) {
	ConfirmLen := func(c *cell, l int) {
		if r := c.Len(); r != l {
			t.Fatalf("%v.Len() should be %v but is %v", c, l, r)
		}
	}
	ConfirmLen(List(), 0)
	ConfirmLen(List(0), 1)
	ConfirmLen(List(0, 1), 2)
	ConfirmLen(List(List(0, 1), 1), 2)
}
*/