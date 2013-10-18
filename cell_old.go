package greenspun
/*
import (
	"fmt"
)

type cell struct {
	Head		interface{}
	Tail		*cell
}

func (c *cell) End() (r *cell) {
	if c != nil {
		for r = c; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (c cell) Offset(i int) (l *cell) {
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

func (c *cell) Each(f interface{}) {
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

func (c *cell) While(f interface{}) (i int, k *cell) {
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

func (c *cell) Until(f interface{}) (i int, k *cell) {
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