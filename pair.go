package greenspun

import (
	"fmt"
	"strings"
)

/*
	This is an implementation of the traditional Lisp dotted pair, storing a data item in the head, and either
	a data item or a pointer to another dotted pair in the tail.

			cf:			http://en.wikipedia.org/wiki/Cons
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

func (c *Pair) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Pair:
		r = c.Equal(&o)
	case *Pair:
		switch {
		case c == nil:
			r = o == nil
		case o != nil:
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

func (c *Pair) Len() (i int) {
	c.Each(func(v interface{}) {
		i++
	})
	return
}

func (c *Pair) IsNil() (r bool) {
	return c == nil
}

func (c *Pair) Push(v interface{}) (r *Pair) {
	if c == nil {
		r = &Pair{ head: v, tail: nil }
	} else {
		r = &Pair{ head: v, tail: c }
	}
	return
}

/*
	Return the data item stored in the current pair, or a nil if the stack is empty.
	If the current cell is nil then panic.
*/
func (c *Pair) Peek() interface{} {
	if c == nil {
		panic(PAIR_EMPTY)
	}
	return c.head
}

/*
	Return the data item stored in the current Pair, along with a reference to the succeeding cell in the stack.

	If the current cell is nil then panic.
*/
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
	if c == nil {
		panic(PAIR_EMPTY)
	}
	return c.head
}

func (c *Pair) Cdr() interface{} {
	if c == nil {
		panic(PAIR_EMPTY)
	}
	return c.tail
}

func (c *Pair) Caar() (r interface{}) {
	if h, ok := c.Car().(*Pair); ok {
		r = h.Car()
	} else {
		panic(PAIR_REQUIRED)
	}
	return
}

func (c *Pair) Cadr() (r interface{}) {
	if h, ok := c.Car().(*Pair); ok {
		r = h.Cdr()
	} else {
		panic(PAIR_REQUIRED)
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
	if c == nil {
		panic(PAIR_EMPTY)
	}
	c.head = i
	return c
}

func (c *Pair) Rplacd(i interface{}) *Pair {
	if c == nil {
		panic(PAIR_EMPTY)
	}
	c.tail = i
	return c
}

func (c *Pair) Move(n int) (r *Pair) {
	switch {
	case c == nil:
		panic(PAIR_EMPTY)
	case n < 0:
		panic(ARGUMENT_NEGATIVE_INDEX)
	default:
		for r = c; n > 0 && r.tail != nil; r = r.Next() {
			n--
		}

		if n > 0 || r == nil {
			panic(PAIR_LIST_TOO_SHALLOW)
		}
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

	if c == nil {
		r = head
	} else {
		c.End().Rplacd(head)
		r = c
	}
	return
}

func (c *Pair) Each(f interface{}) {
	var i		int

	switch f := f.(type) {
	case func():
		for ; c != nil; c = c.Next() {
			f()
		}
	case func(interface{}):
		for ; c != nil; c = c.Next() {
			f(c.Car())
		}
	case func(int, interface{}):
		for ; c != nil; c = c.Next() {
			f(i, c.Car())
			i++
		}
	case func(interface{}, interface{}):
		for ; c != nil; c = c.Next() {
			f(i, c.Car())
			i++
		}
	case func(*Pair):
		for ; c != nil; c = c.Next() {
			f(c)
		}
	case func(int, *Pair):
		for ; c != nil; c = c.Next() {
			f(i, c)
			i++
		}
	case func(interface{}, *Pair):
		for ; c != nil; c = c.Next() {
			f(i, c)
			i++
		}
	}
}

func (c *Pair) Step(start, n int, f interface{}) {
	defer func() {
		switch x := recover(); x {
		case PAIR_EMPTY, PAIR_LIST_TOO_SHALLOW:
			//	An empty or shallow pair indicates that the list has terminated before the next step increment can be reached. In this particular case this is not an error so recover cleanly.
		default:
			//	Any other panic occurring during step iteration should be propagated to the caller.
			panic(x)
		}
	}()

	if start > 0 {
		c = c.Move(start)
	}

	switch f := f.(type) {
	case func():
		for ; c != nil; c = c.Move(n) {
			f()
		}
	case func(interface{}):
		for ; c != nil; c = c.Move(n) {
			f(c.Car())
		}
	case func(int, interface{}):
		for i := 0; c != nil; c = c.Move(n) {
			f(i, c.Car())
			i++
		}
	case func(interface{}, interface{}):
		for i := 0; c != nil; c = c.Move(n) {
			f(i, c.Car())
			i++
		}
	case func(*Pair):
		for ; c != nil; c = c.Move(n) {
			f(c)
		}
	case func(int, *Pair):
		for i := 0; c != nil; c = c.Move(n) {
			f(i, c)
			i++
		}
	case func(interface{}, *Pair):
		for i := 0; c != nil; c = c.Move(n) {
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
		for r := c; r != nil && f(r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(int, interface{}) bool:
		for r := c; r != nil && f(i, r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(interface{}, interface{}) bool:
		for r := c; r != nil && f(i, r.Car()) == condition; r = r.Next() {
			i++
		}
	case func(*Pair) bool:
		for r := c; r != nil && f(r) == condition; r = r.Next() {
			i++
		}
	case func(int, *Pair) bool:
		for r := c; r != nil && f(i, r) == condition; r = r.Next() {
			i++
		}
	case func(interface{}, *Pair) bool:
		for r := c; r != nil && f(i, r) == condition; r = r.Next() {
			i++
		}
	case Equatable:
		for r := c; r != nil && f.Equal(r.Car()) == condition; r = r.Next() {
			i++
		}
	case interface{}:
		for r := c; r != nil && (f == r.Car()) == condition; r = r.Next() {
			i++
		}
	}
	return
}

func (c *Pair) Partition(offset int) (x, y *Pair) {
	defer func() {
		switch x := recover(); x {
		case PAIR_EMPTY, PAIR_LIST_TOO_SHALLOW:
			//	An empty or shallow pair indicates that the list has terminated before the partition has been reached. In this particular case this is not an error so recover cleanly.
			y = nil
		default:
			//	Any other panic occurring during step iteration should be propagated to the caller.
			panic(x)
		}
	}()

	if offset < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	if y = c.Move(offset); y != nil {
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
		for ; c != nil || n != nil; c, n = c.Next(), n.Next() {
			cursor = cursor.append(Cons(c.Car(), n.Car()))
		}
	})
}