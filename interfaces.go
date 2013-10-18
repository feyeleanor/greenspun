package greenspun

type LispPair interface {
	Car() interface{}
	Cdr() interface{}
	Rplaca(v interface{}) LispPair
	Rplacd(v interface{}) LispPair
}

type Equatable interface {
	Equal(f interface{}) bool
}

type Iterable interface {
	Each(f interface{})
}

type Measureable interface {
	Len() int
}