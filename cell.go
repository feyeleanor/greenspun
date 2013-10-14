package greenspun

import (
	"fmt"
)

type Cell struct {
	Head		interface{}
	Tail		interface{}
}

func (c *Cell) String() (r string) {
	if c == nil {
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
	if c != nil {
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