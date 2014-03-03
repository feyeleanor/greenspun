package greenspun

/*
	The SparseArray is a sparse, persistent integer-indexed data store. Internally elements are
	stored in a hash table which provides uniform access, and each access is represented by a
	stack of versioned values.

	Because SparseArray is sparse we allow a default value to be set which will be returned when
	a value within bounds of 0 and length - 1 is queried. Queries out of bounds will panic as
	with a conventional slice.

	The SparseArray contains a creationVersion field which is initialised when a new header is
	created, and a currentVersion which is incremented whenever an element is updated.
*/
type SparseArray struct {
	elements					arrayHash
	length						int
	creationVersion		int
	currentVersion		int
	defaultValue			*arrayElement
	*SparseArray
}

/*
	NewSparseArray returns a SparseArray initialized with a default value and a minimum number
	of elements. If optional arrayHash parameters are provided these will result in values being
	assigned to cells and if necessary the length of the SparseArray adjusted to reflect this.
*/
func NewSparseArray(n int, d interface{}, items ...arrayHash) (r *SparseArray) {
	r = &SparseArray{ elements: make(arrayHash), length: n, defaultValue: &arrayElement{ data: d } }
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

func (s *SparseArray) copyHeader() (r *SparseArray) {
	if s != nil {
		r = &SparseArray{ s.elements, s.length, s.creationVersion, s.currentVersion, s.defaultValue, s.SparseArray }
	}
	return
}

func (s *SparseArray) newHeader() (r *SparseArray) {
	if s != nil {
		creationVersion := s.currentVersion + 1
		r = &SparseArray{ s.elements, s.length, creationVersion, creationVersion, s.defaultValue, s }
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
				if r = Equal(s.Default(), o.Default()); !r {
					return
				}
			}

			for k, vo := range o.elements {
				if vs, ok := s.elements[k]; ok {
					concrete_keys[k] = true
					r = vo.Equal(vs)
				} else {
					r = Equal(vo, s.Default())
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
						r = Equal(vs, o.Default())
					}
				} else {
					delete(concrete_keys, k)
				}
				if !r {
					return
				}
			}

			if len(concrete_keys) > 0 {
				r = Equal(s.Default(), o.Default())
			}
		}
	case nil:
		r = s == nil
	}
	return
}

func (s *SparseArray) Default() (r interface{}) {
	if s != nil {
		r = s.defaultValue.data
	}
	return
}

func (s *SparseArray) SetDefault(v interface{}) (r *SparseArray) {
	if s != nil {
		r = s.newHeader()
		r.defaultValue = &arrayElement{ data: v, version: r.currentVersion, arrayElement: s.defaultValue }
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
		r = s.Default()
	}
	return
}

/*
	For Set() operations we preserve the array header so long as the length of the array doesn't change.
*/
func (s *SparseArray) Set(i int, v interface{}) (r *SparseArray) {
	switch {
	case i < 0:
		panic(ARGUMENT_OUT_OF_BOUNDS)
	case s == nil:
		r = NewSparseArray(0, nil)
	default:
		switch {
		case i >= s.length:
			r = s.newHeader()
			r.length = i + 1
		default:
			r = s
			r.currentVersion++
		}
	}

	if e, ok := r.elements[i]; ok {
		r.elements[i] = &arrayElement{ data: v, version: r.currentVersion, arrayElement: e }
	} else {
		r.elements[i] = &arrayElement{ data: v, version: r.currentVersion }
	}
	if i >= r.length {
		r.length = i + 1
	}
	return
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
		for i = s.length; i > 0; i-- {
			f()
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
		r = s.copyHeader()
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
	/*
		Inserting elements means creating a new SparseArray header and copying the
		current elements across with those from the insertion point onwards shifted
		to their new index

		TO DO:	allow insertion before start of array?
	*/

	if i < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	n := len(items)

	if s == nil {
		r = NewSparseArray(i + n, nil)		
	} else {
		if i < s.length {
			r = NewSparseArray(s.length + n, s.Default())
		} else {
			r = NewSparseArray(i + n, s.Default())
		}
		r.creationVersion = s.currentVersion + 1
		r.currentVersion = r.creationVersion

		for k, v := range s.elements {
			if k >= i {
				r.elements[k + n] = v
			} else {
				r.elements[k] = v
			}
		}
	}

	for k, v := range items {
		r.elements[i + k] = &arrayElement{ data: v, version: r.currentVersion }
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
			r = &SparseArray{ elements: make(arrayHash), currentVersion: s.currentVersion + 1, defaultValue: s.defaultValue, length: s.length - n }
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
		r = NewSparseArray(s.length, s.defaultValue)	// { elements: make(arrayHash), length: s.length, defaultValue: s.defaultValue }
		for k, v := range s.elements {
			r.elements[k] = &arrayElement{ data: v.data }
		}
	}
	return
}

/*
	Commit creates a new SparseArray which is identical to the current state of a give SparseArray
*/
func (s *SparseArray) Commit() (r *SparseArray) {
	if s != nil {
		r = NewSparseArray(s.length, s.Default())
		for k, v := range s.elements {
			r.elements[k] = v.Commit()
		}
	}
	return
}

/*
	Undo returns the previous state of a given SparseArray
*/
func (s *SparseArray) Undo() (r *SparseArray) {
	return
}

/*
	Return a valid header for the state of the SparseArray at a given version point.
*/
func (s *SparseArray) Rollback(version int) (r *SparseArray) {
	if version < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	for ; s != nil && s.creationVersion > version; s = s.SparseArray {}

	if s != nil {
		r = &SparseArray{ elements: make(arrayHash), defaultValue: s.defaultValue, creationVersion: version, currentVersion: version }
		for k, v := range s.elements {
			if x := v.AtVersion(version); x != nil {
				if x.data != s.defaultValue {
					r.elements[k] = x
				}
				if end := k + 1; r.length < end {
					r.length = end
				}
			}
		}

		if l := len(s.elements); l < s.length {
			r.length = s.length
		}
	}
	return
}