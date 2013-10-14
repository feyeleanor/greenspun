package greenspun

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
*/

func Cons(items... interface{}) (c *Cell) {
	switch len(items) {
	case 0:
	case 1:
		c = &Cell{ Head: items[0] }
	case 2:
		c = &Cell{ Head: items[0], Tail: items[1] }
	default:
		c = &Cell{ Head: items[0], Tail: Cons(items[1:]...) }
	}
	return
}