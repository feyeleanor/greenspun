package greenspun
/*
import (
	"fmt"
)

type cell struct {
	Head		interface{}
	Tail		*cell
}

/*
	List() uses the cell type to construct a classic Lisp-style Cons list data structure, a chain of value-link pairs.
*/
/*
func List(items... interface{}) (c *cell) {
	var n *cell
	for i, v := range items {
		if i == 0 {
			c = &cell{ Head: v }
			n = c
		} else {
			n.Tail = &cell{ Head: v }
			n = n.Tail
		}
	}
	return
}

func (c *cell) End() (r *cell) {
	if c != nil {
		for r = c; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (c cell) Offset(i int) (l *cell) {
	switch {
	case i < 0:
		break
	case i == 0:
		l = &c
	default:
		n := &c
		for ; i > 0 && n != nil; i-- {
			n = n.Tail
		}
		if n != nil {
			l = n
		}
	}
	return
}

func (c *cell) Append(l interface{}) (r *cell) {
	if c == nil {
		switch l := l.(type) {
		case cell:
			r = &cell{ l.Head, l.Tail }
		case *cell:
			r = l
		default:
			r = &cell{ Head: l }
		}
	} else {
		r = c
		switch l := l.(type) {
		case cell:
			r.Tail = &cell{ l.Head, l.Tail }
		case *cell:
			r.Tail = l
		default:
			r.Tail = &cell{ Head: l }
		}
	}
	return
}

func (c *cell) Prepend(l interface{}) (r *cell) {
	switch l := l.(type) {
	case cell:
		r = &l
		r.Tail = c
	case *cell:
		r = l
		r.Tail = c
	default:
		r = &cell{ Head: l, Tail: c }
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
	return
}

func (c *cell) Equal(o interface{}) (r bool) {
	if c != nil {
		switch o := o.(type) {
		case *cell:
			r = o != nil && c.equal(*o)
		case cell:
			r = c.equal(o)
		default:
			r = c.equal(cell{ Head: o })
		}
	} else {
		if o, ok := o.(*cell); ok {
			r = o == nil
		}
	}
	return
}

func (c *cell) Each(f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		for k := c; k != nil; k = k.Tail {
			f(k.Head)
		}
	case func(int, interface{}):
		for i, k := 0, c; k != nil; k = k.Tail {
			f(i, k.Head)
			i++
		}
	case func(interface{}, interface{}):
		for i, k := 0, c; k != nil; k = k.Tail {
			f(i, k.Head)
			i++
		}
	}
}

func (c *cell) While(f interface{}) (i int, k *cell) {
	switch f := f.(type) {
	case func(interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if !f(k.Head) {
				break
			}
			i++
		}
	case func(int, interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if !f(i, k.Head) {
				break
			}
			i++
		}
	case func(interface{}, interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if !f(i, k.Head) {
				break
			}
			i++
		}
	case Equatable:
		for k = c; k != nil; k = k.Tail {
			if !f.Equal(k.Head) {
				break
			}
			i++
		}
	case interface{}:
		for k = c; k != nil; k = k.Tail {
			if f != k.Head {
				break
			}
			i++
		}
	}
	return
}

func (c *cell) Until(f interface{}) (i int, k *cell) {
	switch f := f.(type) {
	case func(interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if f(k.Head) {
				break
			}
			i++
		}
	case func(int, interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if f(i, k.Head) {
				break
			}
			i++
		}
	case func(interface{}, interface{}) bool:
		for k = c; k != nil; k = k.Tail {
			if f(i, k.Head) {
				break
			}
			i++
		}
	case Equatable:
		for k = c; k != nil; k = k.Tail {
			if f.Equal(k.Head) {
				break
			}
			i++
		}
	case interface{}:
		for k = c; k != nil; k = k.Tail {
			if f == k.Head {
				break
			}
			i++
		}
	}
	return
}
*/