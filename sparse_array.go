package greenspun

//import "fmt"

/*
	The SparseArray is a sparse, persistent integer-indexed data store. Internally elements are
	stored in a hash table which provides uniform access, and each access is represented by a
	stack of versioned values.

	Because SparseArray is sparse we allow a default value to be set which will be returned when
	a value within bounds of 0 and length - 1 is queried. Queries out of bounds will panic as
	with a conventional slice.
*/
type SparseArray struct {
	elements		arrayHash
	length			int
	version			int
	Default			interface{}
}

func NewSparseArray(n int, d interface{}, items ...arrayHash) (r *SparseArray) {
	r = &SparseArray{ elements: make(arrayHash), length: n, Default: d }
	for _, cells := range items {
		for k, v := range cells {
			if r.length <= k {
				r.length = k + 1
			}
			r.elements[k] = v
		}
	}
	return
}

/*
func (s *SparseArray) String() (r string) {
	return
}
*/

func (s *SparseArray) Len() (r int) {
	if s != nil {
		r = s.length
	}
	return
}

func (s *SparseArray) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case SparseArray:
		r = s.Equal(&o)
	case *SparseArray:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil, s == nil && o != nil:
			r = false
		case s.length != o.length:
			r = false
		case s.length == 0:
			r = true
		default:
			concrete_keys := make(map[int] bool)

			if s.length == o.length && len(s.elements) != s.length && len(o.elements) != o.length {
				if r = Equal(s.Default, o.Default); !r {
					return
				}
			}

			for k, vo := range o.elements {
				if vs, ok := s.elements[k]; ok {
					concrete_keys[k] = true
					r = vo.Equal(vs)
				} else {
					r = Equal(vo, s.Default)
				}
				if !r {
					return
				}
			}

			for k, vs := range s.elements {
				if _, ok := concrete_keys[k]; !ok {
					if vo, ok := o.elements[k]; ok {
						concrete_keys[k] = true
						r = vs.Equal(vo)
					} else {
						r = Equal(vs, o.Default)
					}
				} else {
					delete(concrete_keys, k)
				}
				if !r {
					return
				}
			}

			if len(concrete_keys) > 0 {
				r = Equal(s.Default, o.Default)
			}
		}
	case nil:
		r = s == nil
	}
	return
}

func (s *SparseArray) At(i int) (r interface{}) {
	if s == nil || i < 0 || i >= s.length {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	if e, ok := s.elements[i]; ok && e != nil {
		r = e.data
	} else {
		r = s.Default
	}
	return
}

func (s *SparseArray) Set(i int, v interface{}) *SparseArray {
	switch {
	case i < 0:
		panic(ARGUMENT_OUT_OF_BOUNDS)
	case s == nil:
		a := NewSparseArray(0, nil)
		*s = *a
	default:
		s.version++
	}

	switch e, ok := s.elements[i]; {
	case ok:
		s.elements[i] = &arrayElement{ data: v, version: s.version, arrayElement: e }
	case v != s.Default:
		s.elements[i] = &arrayElement{ data: v, version: s.version }
	}
	if i >= s.length {
		s.length = i + 1
	}
	return s
}

func rescueOutOfBounds() {
	switch x := recover(); x {
	case nil, ARGUMENT_OUT_OF_BOUNDS:
		return
	default:
		panic(x)
	}
}

/*
	Iterate through all cells in order, applying the supplied closure.
*/
func (s *SparseArray) Each(f interface{}) {
	defer rescueOutOfBounds()
	var i	int

	switch f := f.(type) {
	case func():
		for {
			f()
			i++
		}
	case func(interface{}):
		for {
			f(s.At(i))
			i++
		}
	case func(int, interface{}):
		for {
			f(i, s.At(i))
			i++
		}
	case func(interface{}, interface{}):
		for {
			f(i, s.At(i))
			i++
		}
	}
}

func (s *SparseArray) Move(x, y, n int) (r *SparseArray) {
	if s != nil {
		r = &SparseArray{ length: s.length, version: s.version, Default: s.Default, elements: s.elements }
		if y + n >= r.length {
			r.length = y + n
		}
		for n > 0 {
			n--
			r.Set(y + n, r.At(x + n))
			delete(r.elements, x + n)
		}
	}
	return
}

func (s *SparseArray) Insert(i int, items... interface{}) (r *SparseArray) {
	//	Inserting elements means creating a new SparseArray header and copying the
	//	current elements across with those from the insertion point onwards shifted
	//	to their new index

	//	TO DO:	allow insertion before start of array?

	if i < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	n := len(items)

	if s == nil {
		r = NewSparseArray(i + n, nil)		
	} else {
		if i < s.length {
			r = NewSparseArray(s.length + n, s.Default)
			r.version = s.version + 1
		} else {
			r = NewSparseArray(i + n, s.Default)
			r.version = s.version + 1
		}

		for k, v := range s.elements {
			if k >= i {
				r.elements[k + n] = v
			} else {
				r.elements[k] = v
			}
		}
	}

	for k, v := range items {
		r.elements[i + k] = &arrayElement{ data: v, version: r.version }
	}
	return
}

func (s *SparseArray) Delete(i int, params ...int) (r *SparseArray) {
	if i < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	if s != nil {
		r = s
		var n	int
		if len(params) > 0 {
			n = params[0]
		} else {
			n = 1
		}
		switch {
		case n < 0:
			n = 0
		case i + n > s.length:
			n = s.length - i
		}
		if n > 0 {
			r = &SparseArray{ elements: make(arrayHash), version: s.version + 1, Default: s.Default, length: s.length - n }
			last := i + n - 1
			for k, v := range s.elements {
				switch {
				case k < i:
					r.elements[k] = v
				case k > last:
					r.elements[k - n] = v
				default:
				}
			}
		}
	}
	return
}

func (s *SparseArray) Copy() (r *SparseArray) {
	if s != nil {
		r = &SparseArray{ elements: make(arrayHash), length: s.length, Default: s.Default }
		for k, v := range s.elements {
			r.elements[k] = &arrayElement{ data: v.data }
		}
	}
	return
}

func (s *SparseArray) Commit() (r *SparseArray) {
	//	Create a new header which treats the current state of the SparseArray as a base
	//	state for future operations
	if s != nil {
		r = NewSparseArray(s.length, s.Default)
		for k, v := range s.elements {
			r.elements[k] = &arrayElement{ data: v }
		}
	}
	return
}

func (s *SparseArray) Rollback(version int) (r *SparseArray) {
	if version < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}
	if s != nil {
		r = &SparseArray{ elements: make(arrayHash), Default: s.Default, version: version }
		for k, v := range s.elements {
			for ; v != nil && v.version > r.version; v = v.arrayElement {}
			if v != nil {
				r.elements[k] = v
			}
			if r.length <= k {
				r.length = k + 1
			}
		}
	}
	return r
}