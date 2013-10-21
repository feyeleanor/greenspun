package greenspun

func List(items... interface{}) (c *cell) {
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

func (c *cell) Each(f interface{}) {
	if ok := c != nil; ok {
		var i		int
		var l		LispPair

		switch f := f.(type) {
		case func(interface{}):
			for l = c; ok; l, ok = l.Cdr().(LispPair) {
				f(l.Car())
			}
		case func(int, interface{}):
			for l = c; ok; l, ok = l.Cdr().(LispPair) {
				f(i, l.Car())
				i++
			}
		case func(interface{}, interface{}):
			for l = c; ok; l, ok = l.Cdr().(LispPair) {
				f(i, l.Car())
				i++
			}
		}
	}
}