package greenspun

import "fmt"

/*
	A cell is a traditional Lisp dotted pair, storing a data item in the Head, and either a data item or
	a pointer to another dotted pair in the Tail.
*/

type cell struct {
	Head		interface{}
	Tail		interface{}
}

func Cons(head, tail interface{}) (c *cell) {
	return &cell{ head, tail }
}

func (c *cell) String() (r string) {
	if c == nil {
		r = "()"
	} else {
		if t, ok := c.Tail.(LispPair); ok {
			r = fmt.Sprintf("(%v %v)", c.Head, t)
		} else {
			r = fmt.Sprintf("(%v . %v)", c.Head, c.Tail)
		}
	}
	return
}

func (c cell) equal(o cell) (r bool) {
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

	if r {
		if v, ok := c.Tail.(Equatable); ok {
			r = v.Equal(o.Tail)
		} else {
			r = c.Tail == o.Tail
		}
	}
	return
}

func (c cell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case nil:
		r = false
	case cell:
		r = c.equal(o)
	case *cell:
		r = o != nil && c.equal(*o)
	case LispPair:
		r = c.equal(cell{ Head: o.Car(), Tail: o.Cdr() })
	}
	return
}

func (c *cell) Car() (r interface{}) {
	if c != nil {
		r = c.Head
	}
	return
}

func (c *cell) Cdr() (r interface{}) {
	if c != nil {
		r = c.Tail
	}
	return
}

func (c *cell) Rplaca(i interface{}) (r LispPair) {
	if c != nil {
		c.Head = i
	}
	return
}

func (c *cell) Rplacd(i interface{}) (r LispPair) {
	if c != nil {
		c.Tail = i
	}
	return
}