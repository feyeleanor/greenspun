package greenspun

import (
	"fmt"
)

/*
	A Cell is a traditional Lisp dotted pair, storing a data item in the Head, and either a data item or
	a pointer to a Cell in Tail.

	A number of operations are defined on a Cell which treat a chain of cells connected 
*/
type Cell struct {
	Head		interface{}
	Tail		interface{}
}

func Cons(head, tail interface{}) (c *Cell) {
	return &Cell{ head, tail }
}

func List(items... interface{}) (c *Cell) {
	switch len(items) {
	case 0:
	case 1:
		c = Cons(items[0], nil)
	case 2:
		c = Cons(items[0], items[1])
	default:
		c = Cons(items[0], List(items[1:]...))
	}
	return
}

func (c *Cell) String() (r string) {
	if (c == nil) || (c.Head == nil && c.Tail == nil) {
		r = "()"
	} else {
		if t, ok := c.Tail.(*Cell); ok {
			r = fmt.Sprintf("(%v %v)", c.Head, t)
		} else {
			r = fmt.Sprintf("(%v . %v)", c.Head, c.Tail)
		}
	}
	return
}

func (c *Cell) Len() (i int) {
	if (c != nil) && (c.Head != nil || c.Tail != nil) {
		ok := true
		for n := c; ok; n, ok = n.Tail.(*Cell) {
			i++
			c = n
		}
		//	if c.Tail is not a *Cell then it's a value and the length of the chain of values should be incremented
		if c != nil && c.Tail != nil {
			i++
		}
	}
	return
}

func (c Cell) equal(o Cell) (r bool) {
	defer func() {
		if x := recover(); x != nil {
			r = false
		}
	}()
	if v, ok := c.Head.(Equatable); ok {
		r = v.Equal(o.Head)
	} else {
		r = c.Head == o.Head
	}
	if v, ok := c.Tail.(Equatable); ok {
		r = r && v.Equal(o.Tail)
	} else {
		r = r && (c.Tail == o.Tail)
	}
	return
}

func (c *Cell) Equal(o interface{}) (r bool) {
	if c == nil {
		return o == nil
	}
	switch o := o.(type) {
	case *Cell:
		r = o != nil && c.equal(*o)
	default:
		r = c.equal(Cell{ Head: o })
	}
	return
}

func (c *Cell) Car() (r interface{}) {
	if c != nil {
		r = c.Head
	}
	return
}

func (c *Cell) Cdr() (r interface{}) {
	if c != nil {
		r = c.Tail
	}
	return
}

func (c *Cell) Rplaca(i interface{}) {
	if c == nil {
		*c = Cell{ Head: i }
	} else {
		c.Head = i
	}
}

func (c *Cell) Rplacd(i interface{}) {
	if c == nil {
		*c = Cell{ Tail: i }
	} else {
		c.Tail = i
	}
}