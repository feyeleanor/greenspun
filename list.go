package greenspun

func List(items... interface{}) (c *cell) {
	switch len(items) {
	case 0:
		c = nil
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
	ok := true
	for n := c; ok && n != nil; n, ok = n.Tail.(*cell) {
		i++
		c = n
	}
	//	if c.Tail is not a *cell then it's a value and the length of the chain should be incremented
	if c != nil && c.Tail != nil {
		i++
	}
	return
}