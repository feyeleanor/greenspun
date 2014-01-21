package greenspun

type arrayElement struct {
	data					interface{}
	version				int
	*arrayElement
}

func (a *arrayElement) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *arrayElement:
		if a == nil {
			r = o == nil || o.data == nil
		} else if o == nil {
			r = a.data == nil
		} else if v, ok := o.data.(Equatable); ok {
			r = v.Equal(a.data)
		} else if v, ok = a.data.(Equatable); ok {
			r = v.Equal(o.data)
		} else {
			r = a.data == o.data
		}
	case nil:
		r = a == nil || a.data == nil
	default:
		if v, ok := o.(Equatable); ok {
			r = v.Equal(a.data)
		} else if v, ok = a.data.(Equatable); ok {
			r = v.Equal(o)
		} else {
			r = a.data == o
		}
	}
	return
}

type arrayCells	map[int] *arrayElement

func (a arrayCells) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case arrayCells:
		if len(a) == len(o) {
			for k, vo := range o {
				if va, ok := (a)[k]; ok {
					r = va.Equal(vo)
				}
				if !r {
					break
				}
			}
		}
	case nil:
		r = a == nil
	}
	return
}

/*
	The SparseArray is a sparse, persistent integer-indexed data store. Internally elements are
	stored in a hash table which provides uniform access, and each access is represented by a
	stack of versioned values.

	Because SparseArray is sparse we allow a default value to be set which will be returned when
	a value within bounds of 0 and length - 1 is queried. Queries out of bounds will panic as
	with a conventional slice.
*/
type SparseArray struct {
	elements		arrayCells
	length			int
	version			int
	Default			interface{}
}

func NewSparseArray(n int, d interface{}, items ...arrayCells) (r *SparseArray) {
	r = &SparseArray{ elements: make(arrayCells), length: n, Default: d }
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

func (s *SparseArray) Len() int {
	return s.length
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
		default:
			if len(s.elements) != len(o.elements) {
				if r = Equal(s.Default, o.Default); !r {
					break
				}
			}

			r = true
			tested_keys := make(map[int] bool)

			for k, vo := range o.elements {
				if vs, ok := s.elements[k]; ok {
					tested_keys[k] = true
					r = vo.Equal(vs)
				} else {
					r = Equal(vo, s.Default)
				}
				if !r {
					return
				}
			}

			for k, vs := range s.elements {
				if _, ok := tested_keys[k]; !ok {
					if vo, ok := o.elements[k]; ok {
						r = vs.Equal(vo)
					} else {
						r = Equal(vs, o.Default)
					}
				} else {
					delete(tested_keys, k)
				}
				if !r {
					return
				}
			}
		}
	case nil:
		r = s == nil
	}
	return
}

func (s *SparseArray) Set(i int, v interface{}) *SparseArray {
	e := s.elements[i]
	s.elements[i] = &arrayElement{ data: v, version: s.version, arrayElement: e }
	if i > s.length {
		s.length++
	}
	s.version++
	return s
}

func (s *SparseArray) At(i int) (r interface{}) {
	if i < 0 || i >= s.length {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	if e, ok := s.elements[i]; ok && e != nil {
		r = e
	} else {
		r = s.Default
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
	if i < s.length {
		r = NewSparseArray(s.length + n, s.Default)
	} else {
		r = NewSparseArray(i + n, s.Default)
	}
	r.version = s.version + 1
	for k, v := range items {
		r.elements[n + k] = &arrayElement{ data: v, version: r.version }
	}
	for k, v := range s.elements {
		if k >= i {
			r.elements[k + n] = v
		} else {
			r.elements[k] = v
		}
	}
	return
}

func (s *SparseArray) Delete(i, n int) (r *SparseArray) {
	r = &SparseArray{ elements: make(arrayCells), version: s.version + 1, Default: s.Default }
	if n < s.length {
		r.length = s.length - n
		last := i + n - 1
		for k, v := range s.elements {
			switch {
			case k < i:
				r.elements[k] = v
			case k > last:
				r.elements[k - n] = v
			}
		}
	}
	return
}

func (s *SparseArray) Copy() (r *SparseArray) {
	//	A copy can be made by inserting zero elements into an existing array, creating
	//	a new header with an incremented version number
	return s.Insert(0)
}

func (s *SparseArray) Commit() (r *SparseArray) {
	//	Create a new header which treats the current state of the SparseArray as a base
	//	state for future operations
	r = NewSparseArray(s.length, s.Default)
	for k, v := range s.elements {
		r.elements[k] = &arrayElement{ data: v }
	}
	return
}

func (s *SparseArray) Revert(version int) (r *SparseArray) {
	if version < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}
	r = &SparseArray{ elements: make(arrayCells), Default: s.Default, version: version }
	for k, v := range s.elements {
		for ; v != nil && v.version > r.version; v = v.arrayElement {}
		if v != nil {
			r.elements[k] = v
		}
		if r.length <= k {
			r.length = k + 1
		}
	}
	return r
}