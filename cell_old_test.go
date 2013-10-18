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

func TestcellEqual(t *testing.T) {
	ConfirmEqual := func(c, o *cell) {
		if !c.Equal(o) {
			t.Fatalf("%v.Equal(%v) should be true", c, o)
		}
	}
	RefuteEqual := func(c, o *cell) {
		if c.Equal(o) {
			t.Fatalf("%v.Equal(%v) should be false", c, o)
		}
	}

	ConfirmEqual(List(), List())
	ConfirmEqual(List(1), List(1))
	ConfirmEqual(List(List(2, 3), 1), List(List(2, 3), 1))

	RefuteEqual(List(), List(1))
	RefuteEqual(List(), List(List(1)))
	RefuteEqual(List(1), List(2))
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

func TestcellAppend(t *testing.T) {
	ConfirmAppend := func(c *cell, v interface{}, r interface{}) {
		c = c.Append(v)
		if end := c.End(); !end.Equal(r) {
			t.Fatalf("%v.Append(%v) should have tail %v but has %v", c, v, r, end)
		}
	}
	ConfirmAppend(nil, 1, 1)
	ConfirmAppend(nil, 1, List(1))
	ConfirmAppend(nil, List(1), List(1))
	ConfirmAppend(List(), 1, 1)
	ConfirmAppend(List(), 1, List(1))
	ConfirmAppend(List(), List(1), List(1))
	ConfirmAppend(List(1), 2, 2)
	ConfirmAppend(List(1), 2, List(2))
	ConfirmAppend(List(1), List(2), List(2))
}

func TestcellPrepend(t *testing.T) {
	ConfirmPrepend := func(c *cell, v interface{}, r interface{}) {
		c = c.Prepend(v)
		if !c.Equal(r) {
			t.Fatalf("%v.Prepend(%v) should have head %v but has %v", c, v, r, c)
		}
	}
	ConfirmPrepend(nil, 1, 1)
	ConfirmPrepend(nil, 1, List(1))
	ConfirmPrepend(nil, List(1), List(1))
	ConfirmPrepend(List(), 1, 1)
	ConfirmPrepend(List(), 1, List(1))
	ConfirmPrepend(List(), List(1), List(1))
	ConfirmPrepend(List(1), 2, 2)
	ConfirmPrepend(List(1), 2, List(2))
	ConfirmPrepend(List(1), List(2), List(2))
}

func TestcellString(t *testing.T) {
	ConfirmString := func(c *cell, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v but is %v", c, r, s)
		}
	}

	ConfirmString(List(0, 1, 2), "(0 1 2)")
	ConfirmString(List(0, List(1, List(2))), "(0 (1 (2)))")
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