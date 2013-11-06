package greenspun

import (
	"fmt"
	"strings"
)

/*
	A Cell is a traditional Lisp dotted pair, storing a data item in the head, and either a data item or
	a pointer to another dotted pair in the tail.
*/

type Cell struct {
	head		interface{}
	tail		interface{}
}

func Cons(head, tail interface{}) (c *Cell) {
	return &Cell{ head, tail }
}

func List(items... interface{}) (c *Cell) {
	switch len(items) {
	case 0:
		return nil
	case 1:
		c = Cons(items[0], nil)
	default:
		c = Cons(items[0], List(items[1:]...))
	}
	return
}

func (c *Cell) String() string {
	terms := make([]string, 0)
	if c != nil {
		var head	string
		c.Each(func(i int, v *Cell) {
			switch v.head {
			case nil, false:
				head = "nil"
			case true:
				head = "t"
			default:
				head = fmt.Sprintf("%v", v.head)
			}
			if v.tail != nil && v.Next() == nil {
				terms = append(terms, head, ".", fmt.Sprintf("%v", v.tail))
				return
			}
			terms = append(terms, head)
		})
	}
	return "(" + strings.Join(terms, " ") + ")"
}

func (c *Cell) Len() (i int) {
	c.Each(func(v interface{}) {
		i++
	})
	return
}

func (c *Cell) IsNil() (r bool) {
	return c == nil
}

func (c *Cell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Cell:
		r = c.Equal(&o)
	case *Cell:
		if c.IsNil() {
			r = o.IsNil()
		} else if !o.IsNil() {
			if v, ok := c.head.(Equatable); ok {
				r = v.Equal(o.head)
			} else if v, ok = o.head.(Equatable); ok {
				r = v.Equal(c.head)
			} else {
				r = c.head == o.head
			}

			if r {
				if v, ok := c.tail.(Equatable); ok {
					r = v.Equal(o.tail)
				} else if v, ok = o.tail.(Equatable); ok {
					r = v.Equal(c.tail)
				} else {
					r = c.tail == o.tail
				}
			}
		}
	}
	return
}

func (c *Cell) Push(v interface{}) (r *Cell) {
	if !c.IsNil() {
		r = Cons(v, c)
	} else {
		r = Cons(v, nil)
	}
	return 
}

func (c *Cell) Pop() (interface{}, *Cell) {
	return c.Car(), c.Next()
}

//	Combine the first two items at the front of the list into a Cons Cell and make this the front of the list

func (c *Cell) Cons() *Cell {
	return Cons(Cons(c.Car(), c.Cdar()), c.Cdr())
}

//	These are convenience wrappers for two common storage situations: a pair of integers, and a pair of pairs

func (c *Cell) IntPair() (l, r int) {
	return c.Car().(int), c.Cdr().(int)
}

func (c *Cell) CellPair() (l, r *Cell) {
	return c.Car().(*Cell), c.Cdr().(*Cell)
}

func (c *Cell) Next() (r *Cell) {
	r, _ = c.Cdr().(*Cell)
	return
}

func (c *Cell) Car() interface{} {
	if !c.IsNil() {
		return c.head
	}
	return nil
}

func (c *Cell) Cdr() interface{} {
	if !c.IsNil() {
		return c.tail
	}
	return nil
}

func (c *Cell) Caar() (r interface{}) {
	if h, ok := c.Car().(*Cell); ok {
		r = h.Car()
	}
	return
}

func (c *Cell) Cadr() (r interface{}) {
	if h, ok := c.Car().(*Cell); ok {
		r = h.Cdr()
	}
	return
}

func (c *Cell) Cdar() interface{} {
	return c.Next().Car()
}

func (c *Cell) Cddr() interface{} {
	return c.Next().Cdr()
}

func (c *Cell) Rplaca(i interface{}) *Cell {
	if !c.IsNil() {
		c.head = i
		return c
	}
	return nil
}

func (c *Cell) Rplacd(i interface{}) *Cell {
	if !c.IsNil() {
		c.tail = i
		return c
	}
	return nil
}

func (c *Cell) Offset(i int) (r *Cell) {
	switch {
	case i < 0:
		r = nil
	case i == 0:
		r = c
	default:
		n := c
		for ; i > 0 && !n.IsNil(); i-- {
			n = n.Next()
		}
		r = n
	}
	return
}

func (c *Cell) End() (r *Cell) {
	r = c
	for n := c.Next(); !n.IsNil() && n != c; n = n.Next() {
		r = n
	}
	return
}

func (c *Cell) valueAppend(v interface{}) (r *Cell) {
	if x, ok := v.(*Cell); ok {
		c.Rplacd(x)
		r = x.End()
	} else {
		c.Rplacd(Cons(v, nil))
		r = c.Cdr().(*Cell)				
	}
	return
}

func (c *Cell) Append(v... interface{}) (r *Cell) {
	var head *Cell

	if len(v) > 0 {
		if x, ok := v[0].(*Cell); ok {
			head = x
		} else {
			head = Cons(v[0], nil)
		}
		r = head.End()
		for _, v := range v[1:] {
			r = r.valueAppend(v)
		}
	}

	if !c.IsNil() {
		c.End().Rplacd(head)
		r = c
	} else {
		r = head
	}
	return
}

func (c *Cell) Each(f interface{}) {
	c.Step(0, 1, f)
}

func (c *Cell) Step(start, n int, f interface{}) {
	var i		int

	c = c.Offset(start)
	switch f := f.(type) {
	case func():
		for ; !c.IsNil(); c = c.Offset(n) {
			f()
		}
	case func(interface{}):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(c.Car())
		}
	case func(int, interface{}):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c.Car())
			i++
		}
	case func(interface{}, interface{}):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c.Car())
			i++
		}
	case func(*Cell):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(c)
		}
	case func(int, *Cell):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c)
			i++
		}
	case func(interface{}, *Cell):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c)
			i++
		}
	}
}

func (c *Cell) append(v interface{}) (r *Cell) {
	r = Cons(v, nil)
	c.Rplacd(r)
	return
}

func (c *Cell) constructList(f func(anchor *Cell)) *Cell {
	anchor := &Cell{}
	f(anchor)
	return anchor.Next()
}

func (c *Cell) Map(f interface{}) (r *Cell) {
	return c.constructList(func(cursor *Cell) {
		switch f := f.(type) {
		case func(interface{}) interface{}:
			c.Each(func(v interface{}) {
				cursor = cursor.append(f(v))
			})
		case func(int, interface{}) interface{}:
			c.Each(func(i int, v interface{}) {
				cursor = cursor.append(f(i, v))
			})
		case func(interface{}, interface{}) interface{}:
			c.Each(func(k, v interface{}) {
				cursor = cursor.append(f(k, v))
			})
		case func(*Cell) interface{}:
			c.Each(func(v *Cell) {
				cursor = cursor.append(f(v))
			})
		case func(int, *Cell) interface{}:
			c.Each(func(i int, v *Cell) {
				cursor = cursor.append(f(i, v))
			})
		case func(interface{}, *Cell) interface{}:
			c.Each(func(k interface{}, v *Cell) {
				cursor = cursor.append(f(k, v))
			})
		}
	})
}

func (c *Cell) Reduce(seed, f interface{}) (r interface{}) {
	r = seed
	switch f := f.(type) {
	case func(seed, value interface{}) interface{}:
		c.Each(func(v interface{}) {
			r = f(r, v)
		})
	case func(index int, seed, value interface{}) interface{}:
		c.Each(func(i int, v interface{}) {
			r = f(i, r, v)
		})
	case func(key, seed, value interface{}) interface{}:
		c.Each(func(k, v interface{}) {
			r = f(k, r, v)
		})
	case func(seed interface{}, value *Cell) interface{}:
		c.Each(func(v *Cell) {
			r = f(r, v)
		})
	case func(index int, seed interface{}, value *Cell) interface{}:
		c.Each(func(i int, v *Cell) {
			r = f(i, r, v)
		})
	case func(key, seed interface{}, value *Cell) interface{}:
		c.Each(func(k interface{}, v *Cell) {
			r = f(k, r, v)
		})
	}
	return
}

func (c *Cell) While(condition bool, f interface{}) (i int) {
	switch f := f.(type) {
	case func(interface{}) bool:
		for r := c; !r.IsNil() && f(r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(int, interface{}) bool:
		for r := c; !r.IsNil() && f(i, r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(interface{}, interface{}) bool:
		for r := c; !r.IsNil() && f(i, r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(*Cell) bool:
		for r := c; !r.IsNil() && f(r) == condition; r = r.Next() {
			i++
		}
	case func(int, *Cell) bool:
		for r := c; !r.IsNil() && f(i, r) == condition; r = r.Next() {
			i++
		}
	case func(interface{}, *Cell) bool:
		for r := c; !r.IsNil() && f(i, r) == condition; r = r.Next() {
			i++
		}
	case Equatable:
		for r := c; !r.IsNil() && f.Equal(r.Car()) == condition; r = r.Next() {
			i++
		}
	case interface{}:
		for r := c; !r.IsNil() && (f == r.Car()) == condition; r = r.Next() {
			i++
		}
	}
	return
}

func (c *Cell) Partition(offset int) (x, y *Cell) {
	if y = c.Offset(offset); !y.IsNil() {
		r := y.Next()
		y.Rplacd(nil)
		y = r
	}
	return c, y
}

func (c *Cell) Reverse() (r *Cell) {
	c.Each(func(v interface{}) {
		r = r.Push(v)
	})
	return
}

func (c *Cell) Copy() (r *Cell) {
	return c.constructList(func(cursor *Cell) {
		c.Each(func(v interface{}) {
			cursor = cursor.append(v)
		})
	})
}

func (c *Cell) Repeat(count int) (r *Cell) {
	return c.constructList(func(cursor *Cell) {
		for i := count; i > 0; i-- {
			c.Each(func(v interface{}) {
				cursor = cursor.append(v)
			})
		}
	})
}

func (c *Cell) Zip(n *Cell) (r *Cell) {
	return c.constructList(func(cursor *Cell) {
		for ; !c.IsNil() || !n.IsNil(); c, n = c.Next(), n.Next() {
			cursor = cursor.append(Cons(c.Car(), n.Car()))
		}
	})
}