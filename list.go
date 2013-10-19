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

/*
	Calculate the length of a list of cells in terms of the contained items.
	We allow two values to be packed into cell, or a single value and a link to another cell,
	therefore the length may need to be adjusted accordingly.
*/
func (c *cell) Len() (i int) {
	if ok := c != nil; ok {
		for n := c; ok && n != nil; n, ok = n.tail.(*cell) {
			i++
			c = n
		}
		//	if c.tail is not a *cell then it's a value and the length of the chain should be incremented
		if c != nil && c.tail != nil {
			i++
		}		
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