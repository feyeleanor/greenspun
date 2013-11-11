package greenspun

import (
	"fmt"
	"strings"
)

/*
	A Pair is a traditional Lisp dotted pair, storing a data item in the head, and either a data item or
	a pointer to another dotted pair in the tail.
*/

type Pair struct {
	head		interface{}
	tail		interface{}
}

func Cons(head, tail interface{}) (c *Pair) {
	return &Pair{ head, tail }
}

func List(items... interface{}) (c *Pair) {
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

func (c *Pair) String() string {
	terms := make([]string, 0)
	if c != nil {
		var head	string
		c.Each(func(i int, v *Pair) {
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

func (c *Pair) Len() (i int) {
	c.Each(func(v interface{}) {
		i++
	})
	return
}

func (c *Pair) IsNil() (r bool) {
	return c == nil
}

func (c *Pair) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Pair:
		r = c.Equal(&o)
	case *Pair:
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

func (c *Pair) Push(v interface{}) (r *Pair) {
	if !c.IsNil() {
		r = Cons(v, c)
	} else {
		r = Cons(v, nil)
	}
	return 
}

func (c *Pair) Pop() (interface{}, *Pair) {
	return c.Car(), c.Next()
}

//	Combine the first two items at the front of the list into a Cons Pair and make this the front of the list

func (c *Pair) Cons() *Pair {
	return Cons(Cons(c.Car(), c.Cdar()), c.Cddr())
}

//	These are convenience wrappers for two common storage situations: a pair of integers, and a pair of pairs

func (c *Pair) IntPair() (l, r int) {
	return c.Car().(int), c.Cdr().(int)
}

func (c *Pair) PairPair() (l, r *Pair) {
	return c.Car().(*Pair), c.Cdr().(*Pair)
}

func (c *Pair) Next() (r *Pair) {
	r, _ = c.Cdr().(*Pair)
	return
}

func (c *Pair) Car() interface{} {
	if !c.IsNil() {
		return c.head
	}
	return nil
}

func (c *Pair) Cdr() interface{} {
	if !c.IsNil() {
		return c.tail
	}
	return nil
}

func (c *Pair) Caar() (r interface{}) {
	if h, ok := c.Car().(*Pair); ok {
		r = h.Car()
	}
	return
}

func (c *Pair) Cadr() (r interface{}) {
	if h, ok := c.Car().(*Pair); ok {
		r = h.Cdr()
	}
	return
}

func (c *Pair) Cdar() interface{} {
	return c.Next().Car()
}

func (c *Pair) Cddr() interface{} {
	return c.Next().Cdr()
}

func (c *Pair) Rplaca(i interface{}) *Pair {
	if !c.IsNil() {
		c.head = i
		return c
	}
	return nil
}

func (c *Pair) Rplacd(i interface{}) *Pair {
	if !c.IsNil() {
		c.tail = i
		return c
	}
	return nil
}

func (c *Pair) Offset(i int) (r *Pair) {
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

func (c *Pair) End() (r *Pair) {
	r = c
	for n := c.Next(); !n.IsNil() && n != c; n = n.Next() {
		r = n
	}
	return
}

func (c *Pair) valueAppend(v interface{}) (r *Pair) {
	if x, ok := v.(*Pair); ok {
		c.Rplacd(x)
		r = x.End()
	} else {
		c.Rplacd(Cons(v, nil))
		r = c.Cdr().(*Pair)				
	}
	return
}

func (c *Pair) Append(v... interface{}) (r *Pair) {
	var head *Pair

	if len(v) > 0 {
		if x, ok := v[0].(*Pair); ok {
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

func (c *Pair) Each(f interface{}) {
	c.Step(0, 1, f)
}

func (c *Pair) Step(start, n int, f interface{}) {
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
	case func(*Pair):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(c)
		}
	case func(int, *Pair):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c)
			i++
		}
	case func(interface{}, *Pair):
		for ; !c.IsNil(); c = c.Offset(n) {
			f(i, c)
			i++
		}
	}
}

func (c *Pair) append(v interface{}) (r *Pair) {
	r = Cons(v, nil)
	c.Rplacd(r)
	return
}

func (c *Pair) constructList(f func(anchor *Pair)) *Pair {
	anchor := &Pair{}
	f(anchor)
	return anchor.Next()
}

func (c *Pair) Map(f interface{}) (r *Pair) {
	return c.constructList(func(cursor *Pair) {
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
		case func(*Pair) interface{}:
			c.Each(func(v *Pair) {
				cursor = cursor.append(f(v))
			})
		case func(int, *Pair) interface{}:
			c.Each(func(i int, v *Pair) {
				cursor = cursor.append(f(i, v))
			})
		case func(interface{}, *Pair) interface{}:
			c.Each(func(k interface{}, v *Pair) {
				cursor = cursor.append(f(k, v))
			})
		}
	})
}

func (c *Pair) Reduce(seed, f interface{}) (r interface{}) {
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
	case func(seed interface{}, value *Pair) interface{}:
		c.Each(func(v *Pair) {
			r = f(r, v)
		})
	case func(index int, seed interface{}, value *Pair) interface{}:
		c.Each(func(i int, v *Pair) {
			r = f(i, r, v)
		})
	case func(key, seed interface{}, value *Pair) interface{}:
		c.Each(func(k interface{}, v *Pair) {
			r = f(k, r, v)
		})
	}
	return
}

func (c *Pair) While(condition bool, f interface{}) (i int) {
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
	case func(*Pair) bool:
		for r := c; !r.IsNil() && f(r) == condition; r = r.Next() {
			i++
		}
	case func(int, *Pair) bool:
		for r := c; !r.IsNil() && f(i, r) == condition; r = r.Next() {
			i++
		}
	case func(interface{}, *Pair) bool:
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

func (c *Pair) Partition(offset int) (x, y *Pair) {
	if y = c.Offset(offset); !y.IsNil() {
		r := y.Next()
		y.Rplacd(nil)
		y = r
	}
	return c, y
}

func (c *Pair) Reverse() (r *Pair) {
	c.Each(func(v interface{}) {
		r = r.Push(v)
	})
	return
}

func (c *Pair) Copy() (r *Pair) {
	return c.constructList(func(cursor *Pair) {
		c.Each(func(v interface{}) {
			cursor = cursor.append(v)
		})
	})
}

func (c *Pair) Repeat(count int) (r *Pair) {
	return c.constructList(func(cursor *Pair) {
		for i := count; i > 0; i-- {
			c.Each(func(v interface{}) {
				cursor = cursor.append(v)
			})
		}
	})
}

func (c *Pair) Zip(n *Pair) (r *Pair) {
	return c.constructList(func(cursor *Pair) {
		for ; !c.IsNil() || !n.IsNil(); c, n = c.Next(), n.Next() {
			cursor = cursor.append(Cons(c.Car(), n.Car()))
		}
	})
}