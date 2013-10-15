package greenspun
/*
import (
	"fmt"
)

type Cell struct {
	Head		interface{}
	Tail		*Cell
}

/*
	List() uses the Cell type to construct a classic Lisp-style Cons list data structure, a chain of value-link pairs.
*/
/*
func List(items... interface{}) (c *Cell) {
	var n *Cell
	for i, v := range items {
		if i == 0 {
			c = &Cell{ Head: v }
			n = c
		} else {
			n.Tail = &Cell{ Head: v }
			n = n.Tail
		}
	}
	return
}

func (c *Cell) End() (r *Cell) {
	if c != nil {
		for r = c; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (c Cell) Offset(i int) (l *Cell) {
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

func (c *Cell) Append(l interface{}) (r *Cell) {
	if c == nil {
		switch l := l.(type) {
		case Cell:
			r = &Cell{ l.Head, l.Tail }
		case *Cell:
			r = l
		default:
			r = &Cell{ Head: l }
		}
	} else {
		r = c
		switch l := l.(type) {
		case Cell:
			r.Tail = &Cell{ l.Head, l.Tail }
		case *Cell:
			r.Tail = l
		default:
			r.Tail = &Cell{ Head: l }
		}
	}
	return
}

func (c *Cell) Prepend(l interface{}) (r *Cell) {
	switch l := l.(type) {
	case Cell:
		r = &l
		r.Tail = c
	case *Cell:
		r = l
		r.Tail = c
	default:
		r = &Cell{ Head: l, Tail: c }
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
	return
}

func (c *Cell) Equal(o interface{}) (r bool) {
	if c != nil {
		switch o := o.(type) {
		case *Cell:
			r = o != nil && c.equal(*o)
		case Cell:
			r = c.equal(o)
		default:
			r = c.equal(Cell{ Head: o })
		}
	} else {
		if o, ok := o.(*Cell); ok {
			r = o == nil
		}
	}
	return
}

func (c *Cell) Each(f interface{}) {
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

func (c *Cell) While(f interface{}) (i int, k *Cell) {
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

func (c *Cell) Until(f interface{}) (i int, k *Cell) {
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