package greenspun

type Nilable interface {
	IsNil() bool
}

type LispPair interface {
	Car() interface{}
	Cdr() interface{}
	Rplaca(v interface{}) LispPair
	Rplacd(v interface{}) LispPair
	IsNil() bool
}

type Equatable interface {
	Equal(f interface{}) bool
}