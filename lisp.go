package greenspun

const _ITERATION = iota

func ThrowIteration() {
	panic(_ITERATION)
}

func CatchIteration() {
	if x := recover(); x != _ITERATION {
		panic(x)
	}
}

func Len(l LispPair) (i int) {
	if l != nil {
		if m, ok := l.(Measureable); ok {
			i = m.Len()
		} else {
			i = While(l, func(v interface{}) bool {
				return true
			})
		}
	}
	return
}

func Equal(l, o LispPair) (r bool) {
	if l == nil {
		r = o == nil
	} else {
		defer CatchIteration()
		not_equal := func() {
			ThrowIteration()
		}
		Each(l, func(v interface{}) {
			car := o.Car()
			o, _ = o.Cdr().(LispPair)
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
			not_equal()
		})
		r = Len(o) == 0
	}
	return
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

func End(l LispPair) (r LispPair) {
	if l != nil {
		if cdr, ok := l.Cdr().(LispPair); ok {
			r = End(cdr)
		} else {
			r = l
		}
	}
	return
}

func Each(l, f interface{}) {
	switch l := l.(type) {
	case Iterable:
		l.Each(f)
	case LispPair:
		if l != nil {
			var i int
			switch f := f.(type) {
			case func(interface{}):
				for ok := true; ok; l, ok = l.Cdr().(LispPair) {
					f(l.Car())
				}
				if cdr := Cdr(l); cdr != nil {
					f(cdr)
				}
			case func(int, interface{}):
				for ok := true; ok; l, ok = l.Cdr().(LispPair) {
					f(i, l.Car())
					i++
				}
				if cdr := l.Cdr(); cdr != nil {
					f(i, cdr)
				}
			case func(interface{}, interface{}):
				for ok := true; ok; l, ok = l.Cdr().(LispPair) {
					f(i, l.Car())
					i++
				}
				if cdr := l.Cdr(); cdr != nil {
					f(i, cdr)
				}
			}
		}
	}
}

func While(l, f interface{}) (i int) {
	switch l := l.(type) {
	case Iterable:
		defer CatchIteration()
		switch f := f.(type) {
		case func(interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(v) {
					ThrowIteration()
				}
				i++
			})
		case func(int, interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(i, v) {
					ThrowIteration()
				}
				i++
			})
		case func(interface{}, interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(i, v) {
					ThrowIteration()
				}
				i++
			})
		case Equatable:
			l.Each(func(v interface{}) {
				if !f.Equal(v) {
					ThrowIteration()
				}
				i++
			})
		case interface{}:
			l.Each(func(v interface{}) {
				if f != v {
					ThrowIteration()
				}
				i++
			})
		}
	case LispPair:
		if l != nil {
			switch f := f.(type) {
			case func(interface{}) bool:
				for r, ok := l, true; ok && f(Car(r)); r, ok = r.Cdr().(LispPair) {
					i++
				}
			case func(int, interface{}) bool:
				for r, ok := l, true; ok && f(i, r.Car()); r, ok = r.Cdr().(LispPair) {
					i++
				}
			case func(interface{}, interface{}) bool:
				for r, ok := l, true; ok && f(i, r.Car()); r, ok = r.Cdr().(LispPair) {
					i++
				}
			case Equatable:
				for r, ok := l, true; ok && f.Equal(r.Car()); r, ok = r.Cdr().(LispPair) {
					i++
				}
			case interface{}:
				for r, ok := l, true; ok && f == r.Car(); r, ok = r.Cdr().(LispPair) {
					i++
				}
			}
		}
	}
	return
}

func Until(l LispPair, f interface{}) (i int) {
	switch l := l.(type) {
	case Iterable:
		defer CatchIteration()
		switch f := f.(type) {
		case func(interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(v) {
					ThrowIteration()
				}
				i++
			})
		case func(int, interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(i, v) {
					ThrowIteration()
				}
				i++
			})
		case func(interface{}, interface{}) bool:
			l.Each(func(v interface{}) {
				if !f(i, v) {
					ThrowIteration()
				}
				i++
			})
		case Equatable:
			l.Each(func(v interface{}) {
				if !f.Equal(v) {
					ThrowIteration()
				}
				i++
			})
		case interface{}:
			l.Each(func(v interface{}) {
				if f != v {
					ThrowIteration()
				}
				i++
			})
		}
	case LispPair:
		if l != nil {
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
	}
	return
}