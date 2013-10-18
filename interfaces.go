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

/*
	An iterable produces each of its elements in an order defined by its underlying type.
	Each element is combined with the value contained in f according to the rules defined
	by the underlying type of the iterable.
	Usually f will be a function which takes appropriate parameters or a channel but the
	details are up to the specific type.
*/
type Iterable interface {
	Each(f interface{})
}

type Measureable interface {
	Len() int
}