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