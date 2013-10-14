package greenspun

type LispPair interface {
	Car() interface{}
	Cdr() interface{}
	Rplaca(v interface{})
	Rplacd(v interface{})
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