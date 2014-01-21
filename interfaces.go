package greenspun

type Nilable interface {
	IsNil() bool
}

type Equatable interface {
	Equal(f interface{}) bool
}

type HasLength interface {
	Len() int
}

type Environment interface {
	Bind(k, v interface{})
	Find(k interface{}) interface{}
}

type Executable interface {
	Eval(e Environment) interface{}
}

func Equal(lhs, rhs interface{}) (r bool) {
	if d, ok := lhs.(Equatable); ok {
		r = d.Equal(rhs)
	} else if d, ok = rhs.(Equatable); ok {
		r = d.Equal(lhs)
	} else {
		r = lhs == rhs
	}
	return
}