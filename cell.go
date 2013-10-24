package greenspun

import (
	"fmt"
	"strings"
)

/*
	A cell is a traditional Lisp dotted pair, storing a data item in the head, and either a data item or
	a pointer to another dotted pair in the tail.
*/

type cell struct {
	head		interface{}
	tail		interface{}
}

func Cons(head, tail interface{}) (c *cell) {
	return &cell{ head, tail }
}

func (c *cell) String() (r string) {
	if c != nil {
		if _, ok := c.tail.(LispPair); ok {
			terms := make([]string, 0, 4)		
			Each(c, func(v interface{}) {
				terms = append(terms, fmt.Sprintf("%v", v))
			})
			r = strings.Join(terms, " ")
		} else {
			if c.tail == nil {
				r = fmt.Sprintf("%v", c.head)
			} else {
				r = fmt.Sprintf("%v . %v", c.head, c.tail)
			}
		}
	}
	return "(" + r + ")"
}

func (c *cell) IsNil() (r bool) {
	return c == nil
}

func (c *cell) equal(o *cell) (r bool) {
	if r = c.IsNil() && o.IsNil(); !r {
		defer IgnorePanic()
		if v, ok := c.head.(Equatable); ok {
			r = v.Equal(o.head)
		} else {
			r = c.head == o.head
		}

		if r {
			if v, ok := c.tail.(Equatable); ok {
				r = v.Equal(o.tail)
			} else {
				r = c.tail == o.tail
			}
		}
	}
	return
}

func (c *cell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *cell:
		r = o != nil && c.equal(o)
	case LispPair:
		r = c.equal(Cons(o.Car(), o.Cdr()))
	}
	return
}

func (c *cell) Car() interface{} {
	if !c.IsNil() {
		return c.head
	}
	return nil
}

func (c *cell) Cdr() interface{} {
	if !c.IsNil() {
		return c.tail
	}
	return nil
}

func (c *cell) Rplaca(i interface{}) LispPair {
	if !c.IsNil() {
		c.head = i
		return c
	}
	return nil
}

func (c *cell) Rplacd(i interface{}) LispPair {
	if !c.IsNil() {
		c.tail = i
		return c
	}
	return nil
}