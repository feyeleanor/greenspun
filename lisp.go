package greenspun

import "fmt"

const _ITERATION = "iteration"

func ThrowIteration() {
	panic(_ITERATION)
}

func CatchIteration() {
	if x := recover(); x != _ITERATION {
		panic(x)
	}
}

/*
	A convenience wrapper for calculating the length of a chain of LispPairs.
*/
func Len(l LispPair) (i int) {
	Each(l, func(v interface{}) {
		i++
	})
	return
}

/*
	We define an atom as any type which doesn't implement the LispPair interface or
	which does implement a LispPair and has a Cdr() which links to another LispPair.
	This means go's slices, hashes and channels are by default atoms.
*/

func IsList(v interface{}) (r bool) {
	if v, ok := v.(LispPair); ok {
		if v, r = v.Cdr().(LispPair); !r {
			r = IsNil(v)
		} 
	}
	return
}

func IsAtom(v interface{}) bool {
	return !IsList(v)
}

func IsNil(v interface{}) (r bool) {
	if v == nil {
		r = true
	} else {
		if v, ok := v.(Nilable); ok {
			r = v.IsNil()
		}
	}
	return
}

func Equal(l, r interface{}) bool {
	switch {
	case IsNil(l):
		return IsNil(r)
	case IsNil(r):
		return false
	default:
		if l, ok := l.(Equatable); ok {
			return l.Equal(r)
		}
		if r, ok := r.(Equatable); ok {
			return r.Equal(l)
		}
		if l, ok := l.(LispPair); ok {
			if r, ok := r.(LispPair); ok {
				defer CatchIteration()
				Each(l, func(v interface{}) {
					car := r.Car()
					r, _ = r.Cdr().(LispPair)
					if car != nil {
						if v, ok := v.(Equatable); ok && v.Equal(car) {
							return
						}
						if car, ok := car.(Equatable); ok && car.Equal(v) {
							return
						}
						if v == car {
							return
						}
					}
					ThrowIteration()
				})
				return Len(r) == 0
			}
		}
	}
	return l == r
}

func Car(l LispPair) (r interface{}) {
	if l != nil {
		r = l.Car()
	}
	return
}

func Cdr(l LispPair) (r interface{}) {
	if l != nil {
		r = l.Cdr()
	}
	return
}

func Caar(l LispPair) (r interface{}) {
	r = Car(l)
	if h, ok := r.(LispPair); ok {
		r = Car(h)
	} else {
		r = nil
	}
	return
}

func Cadr(l LispPair) (r interface{}) {
	r = Car(l)
	if h, ok := r.(LispPair); ok {
		r = Cdr(h)
	} else {
		r = nil
	}
	return
}

func Cdar(l LispPair) (r interface{}) {
	r = Cdr(l)
	if h, ok := r.(LispPair); ok {
		r = Car(h)
	} else {
		r = nil
	}
	return
}

func Cddr(l LispPair) (r interface{}) {
	r = Cdr(l)
	if h, ok := r.(LispPair); ok {
		r = Cdr(h)
	} else {
		r = nil
	}
	return
}

func Offset(l LispPair, i int) (r LispPair) {
	switch {
	case i < 0:
		r = nil
	case i == 0:
		r = l
	default:
		n := l
		for ok := true; i > 0 && ok && n != nil; i-- {
			n, ok = n.Cdr().(LispPair)
		}
		r = n
	}
	return
}

func End(l LispPair) (r LispPair) {
	if do := IsList(l); !do {
		r = l
	} else {
		r = l
		n, do := l.Cdr().(LispPair)
		for ; do && n != l; n, do = n.Cdr().(LispPair) {
			r = n
		}
	}
	return
}

func valueAppend(l LispPair, v interface{}) (r LispPair) {
	if x, ok := v.(LispPair); ok {
		l.Rplacd(x)
		r = End(x)
	} else {
		l.Rplacd(Cons(v, nil))
		r = l.Cdr().(LispPair)				
	}
	return
}

func Append(l LispPair, v... interface{}) (r LispPair) {
	var head LispPair

	if len(v) > 0 {
		if x, ok := v[0].(LispPair); ok {
			head = x
		} else {
			head = Cons(v[0], nil)
		}
		r = End(head)
		for _, v := range v[1:] {
			r = valueAppend(r, v)
		}
	}

	if !IsNil(l) {
		End(l).Rplacd(head)
		r = l
	} else {
		r = head
	}
	return
}

func Each(l LispPair, f interface{}) {
	if ok := !IsNil(l); ok {
		var i int
		switch f := f.(type) {
		case func(interface{}):
			for ; ok; l, ok = l.Cdr().(LispPair) {
				f(l.Car())
			}
		case func(int, interface{}):
			for ; ok; l, ok = l.Cdr().(LispPair) {
				f(i, l.Car())
				i++
			}
		case func(interface{}, interface{}):
			for ; ok; l, ok = l.Cdr().(LispPair) {
				f(i, l.Car())
				i++
			}
		}
	}
}

func Map(l LispPair, f interface{}) (r LispPair) {
	if !IsNil(l) {
		var c *cell
		head := &cell{}
		r = head
		switch f := f.(type) {
		case func(interface{}) interface{}:
			Each(l, func(v interface{}) {
				c = Cons(f(v), nil)
				r.Rplacd(c)
				r = c
			})
		case func(int, interface{}) interface{}:
			Each(l, func(i int, v interface{}) {
				c = Cons(f(i, v), nil)
				r.Rplacd(c)
				r = c
			})
		case func(interface{}, interface{}) interface{}:
			Each(l, func(k, v interface{}) {
				c = Cons(f(k, v), nil)
				r.Rplacd(c)
				r = c
			})
		}
		if head.Cdr() != nil {
			r = head.Cdr().(LispPair)
		} else {
			r = nil
		}
	}
	return
}

func Reduce(l LispPair, seed, f interface{}) (r interface{}) {
	if !IsNil(l) {
		r = seed
		switch f := f.(type) {
		case func(seed, value interface{}) interface{}:
			Each(l, func(v interface{}) {
				r = f(r, v)
			})
		case func(index int, seed, value interface{}) interface{}:
			Each(l, func(i int, v interface{}) {
				r = f(i, r, v)
			})
		case func(key, seed, value interface{}) interface{}:
			Each(l, func(k, v interface{}) {
				r = f(k, r, v)
			})
		}
	}
	return
}

func While(l LispPair, f interface{}) (i int) {
	if !IsNil(l) {
		var r LispPair
		var ok bool
		switch f := f.(type) {
		case func(interface{}) bool:
			for r, ok = l, true; ok && f(Car(r)); i++ {
				r, ok = r.Cdr().(LispPair)
			}
		case func(int, interface{}) bool:
			for r, ok = l, true; ok && f(i, r.Car()); i++ {
				r, ok = r.Cdr().(LispPair)
			}
		case func(interface{}, interface{}) bool:
			for r, ok = l, true; ok && f(i, r.Car()); i++ {
				r, ok = r.Cdr().(LispPair)
			}
		case Equatable:
			for r, ok = l, true; ok && f.Equal(r.Car()); i++ {
				r, ok = r.Cdr().(LispPair)
			}
		case interface{}:
			for r, ok = l, true; ok && f == r.Car(); i++ {
				r, ok = r.Cdr().(LispPair)
			}
		}
	}
	return
}

func Until(l LispPair, f interface{}) (i int) {
	if !IsNil(l) {
		switch f := f.(type) {
		case func(interface{}) bool:
			for r, ok := l, true; ok && !f(r.Car()); r, ok = r.Cdr().(LispPair) {
				i++
			}
		case func(int, interface{}) bool:
			for r, ok := l, true; ok && !f(i, r.Car()); r, ok = r.Cdr().(LispPair) {
				i++
			}
		case func(interface{}, interface{}) bool:
			for r, ok := l, true; ok && !f(i, r.Car()); r, ok = r.Cdr().(LispPair) {
				i++
			}
		case Equatable:
			for r, ok := l, true; ok && !f.Equal(r.Car()); r, ok = r.Cdr().(LispPair) {
				i++
			}
		case interface{}:
			for r, ok := l, true; ok && f != r.Car(); r, ok = r.Cdr().(LispPair) {
				i++
			}
		}
	}
	return
}

func Reverse(l LispPair) (r LispPair) {
	for !IsNil(l) {
		n, _ := l.Cdr().(LispPair)
		l.Rplacd(r)
		r = l
		l = n
	}
	return
}

func Copy(l LispPair) (r LispPair) {
	if !IsNil(l) {
		var c *cell
		head := &cell{}
		r = head

		Each(l, func(v interface{}) {
			c = Cons(v, nil)
			r.Rplacd(c)
			r = c
		})

		if head.Cdr() != nil {
			r = head.Cdr().(LispPair)
		} else {
			r = nil
		}
	}
fmt.Printf("Copy(%v) => %v\n", l, r)
	return
}

func Repeat(l LispPair, count int) (r LispPair) {
	if !IsNil(l) {
		var c *cell
		head := &cell{}
		r = head

		for i := count; i > 0; i-- {
			Each(l, func(v interface{}) {
				c = Cons(v, nil)
				r.Rplacd(c)
				r = c
			})
		}

		if head.Cdr() != nil {
			r = head.Cdr().(LispPair)
		} else {
			r = nil
		}
	}
	return
}