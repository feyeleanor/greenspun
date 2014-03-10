package greenspun

import "fmt"

//	The SparseSlice is a sparse, persistent integer-indexed data store. Internally elements are
//	stored in a hash table which provides uniform access, and each access is represented by a
//	stack of versioned values.
//
//	The SparseSlice contains a creationVersion field which is initialised when a new header is
//	created, and a currentVersion which is incremented whenever an element is modified.
//
//	Because SparseSlice is sparse we allow a default value to be set which will be returned when
//	a value within bounds of 0 and length - 1 is queried. Queries out of bounds will panic as
//	with a conventional slice.
//
type SparseSlice struct {
	elements					sliceHash				"versioned elements"
	length						int							"total number of elements in the slice"
	creationVersion		int							"the version number for which this header was created"
	currentVersion		int							"the version number of the most recent change"
	defaultValue			*versionedValue	"the default value for unset elements in the slice"
	*SparseSlice											"the SparseSlice from which the current header is derived"
}

//	NewSparseSlice returns a SparseSlice initialized with a default value and a minimum number
//	of elements. If optional sliceHash parameters are provided these will result in values being
//	assigned to cells and if necessary the length of the SparseSlice adjusted to reflect this.
//
func NewSparseSlice(n int, d interface{}, items ...sliceHash) (r *SparseSlice) {
	r = &SparseSlice{ elements: make(sliceHash), length: n, defaultValue: &versionedValue{ data: d } }
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

//	Make a duplicate of the current SparseSlice header.
//
func (s *SparseSlice) copyHeader() (r *SparseSlice) {
	if s != nil {
		r = &SparseSlice{ s.elements, s.length, s.creationVersion, s.currentVersion, s.defaultValue, s.SparseSlice }
	}
	return
}

//	Returns a new header referencing the same elements etc. as the current header but with the version
//	incremented and a pointer to the current header as its parent.
//
func (s *SparseSlice) newVersion() (r *SparseSlice) {
	if s != nil {
		creationVersion := s.currentVersion + 1
		r = &SparseSlice{ s.elements, s.length, creationVersion, creationVersion, s.defaultValue, s }
	}
	return
}

func (s *SparseSlice) String() (r string) {
	if s != nil {
		r = "["
		for i := 0; i < s.length; i++ {
			r = fmt.Sprintf("%v%v ", r, s.At(i))
		}
		if l := len(r); l > 1 {
			r = r[:l - 1]
		}
		r += "]"
	} else {
		r = fmt.Sprintf("%v", nil)
	}
	return
}

//	Return the current length of the slice.
//
func (s *SparseSlice) Len() (r int) {
	if s != nil {
		r = s.length
	}
	return
}

//	Determine if a passed value is equivalent to the current slice.
//
func (s *SparseSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case SparseSlice:
		r = s.Equal(&o)
	case *SparseSlice:
		switch {
		case s == nil && o == nil:
			r = true
		case s == nil, o == nil:
			r = false
		case s.length != o.length:
			r = false
		case s.length == 0:
			r = true
		default:
			instantiated_keys := make(map[int] bool)

			if len(s.elements) != s.length && len(o.elements) != o.length {
				if r = Equal(s.Default(), o.Default()); !r {
					return
				}
			}

			for k, vo := range o.elements {
				if vs, ok := s.elements[k]; ok {
					instantiated_keys[k] = true
					r = vo.Equal(vs)
				} else {
					r = Equal(vo, s.Default())
				}
				if !r {
					return
				}
			}

			for k, vs := range s.elements {
				if _, ok := instantiated_keys[k]; !ok {
					if vo, ok := o.elements[k]; ok {
						r = vs.Equal(vo)
					} else {
						r = Equal(vs, o.Default())
					}
				}
				if !r {
					return
				}
			}
			r = true
		}
	case nil:
		r = s == nil
	}
	return
}

//	Return the default value for the slice.
//
func (s *SparseSlice) Default() (r interface{}) {
	if s != nil {
		r = s.defaultValue.data
	}
	return
}

//	Assign a new default value for the slice.
//
func (s *SparseSlice) SetDefault(v interface{}) (r *SparseSlice) {
	if s != nil {
		r.currentVersion++
		r.defaultValue = &versionedValue{ data: v, version: r.currentVersion, versionedValue: s.defaultValue }
	}
	return
}

//	At returns the value stored at a particular location in the slice, which may be the default value or
//	else a value which has been explicitly set. In the event that the location is not within the bounds
//	specified by the slice a panic is raised with ARGUMENT_OUT_OF_BOUNDS.
//
func (s *SparseSlice) At(i int) (r interface{}) {
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

//	For Set() operations we preserve the slice header so long as the length of the slice doesn't change.
//
func (s *SparseSlice) Set(i int, v interface{}) (r *SparseSlice) {
	switch {
	case i < 0:
		panic(ARGUMENT_OUT_OF_BOUNDS)
	case s == nil:
		r = NewSparseSlice(0, nil)
	default:
		switch {
		case i >= s.length:
			r = s.newVersion()
			r.length = i + 1
		default:
			r = s
			r.currentVersion++
		}
	}

	if e, ok := r.elements[i]; ok {
		r.elements[i] = &versionedValue{ data: v, version: r.currentVersion, versionedValue: e }
	} else {
		r.elements[i] = &versionedValue{ data: v, version: r.currentVersion }
	}
	if i >= r.length {
		r.length = i + 1
	}
	return
}

//	Iterate through all cells in order, applying the supplied closure to the current cell value.
//
func (s *SparseSlice) Each(f interface{}) {
	defer RescueOutOfBounds()
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

//	Return a new header in which the specified range of cells has been moved to the specified
//	target locations.
//
func (s *SparseSlice) Move(x, y, n int) (r *SparseSlice) {
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

//	Create a new SparseSlice header referencing the cells in the existing SparseSlice but with
//	the specified items inserted.
//
//	If an existing cell currently contains the default value, we don't bother to reference this
//	in the new SparseSlice. Likewise if one of the values to be inserted is the same as the
//	default value then we don't bother creating a SparseSlice entry.
//
func (s *SparseSlice) Insert(i int, items... interface{}) (r *SparseSlice) {
	if i < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	n := len(items)

	if s == nil {
		r = NewSparseSlice(i + n, nil)		
	} else {
		if i < s.length {
			r = NewSparseSlice(s.length + n, s.Default())
		} else {
			r = NewSparseSlice(i + n, s.Default())
		}
		r.creationVersion = s.currentVersion + 1
		r.currentVersion = r.creationVersion

		for k, v := range s.elements {
			switch {
			case v.data == s.defaultValue:
				continue
			case k >= i:
				r.elements[k + n] = v
			default:
				r.elements[k] = v
			}
		}
	}

	for k, v := range items {
		if v != r.defaultValue {
			r.elements[i + k] = &versionedValue{ data: v, version: r.currentVersion }
		}
	}
	return
}

//	Return a new SparseSlice header referencing the cells in the existing SparseSlice but with
//	one or more cells removed.
//
//	If an existing cell currently contains the default value, we don't bother to reference this
//	in the new SparseSlice.
//
func (s *SparseSlice) Delete(i int, params ...int) (r *SparseSlice) {
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
			r = &SparseSlice{ elements: make(sliceHash), currentVersion: s.currentVersion + 1, defaultValue: s.defaultValue, length: s.length - n }
			last := i + n - 1
			for k, v := range s.elements {
				switch {
				case v == r.defaultValue:
					continue
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

//	Return a new SparseSlice header referencing the cells in the existing SparseSlice.
//
//	If an existing cell currently contains the default value, we don't bother to reference this
//	in the new SparseSlice.
//
func (s *SparseSlice) Copy() (r *SparseSlice) {
	if s != nil {
		r = NewSparseSlice(s.length, s.defaultValue)
		for k, v := range s.elements {
			if v.data != r.defaultValue {
				r.elements[k] = &versionedValue{ data: v.data }
			}
		}
	}
	return
}

//	Commit creates a new SparseSlice which is identical to the current state of a given SparseSlice.
//
func (s *SparseSlice) Commit() (r *SparseSlice) {
	if s != nil {
		r = NewSparseSlice(s.length, s.Default())
		for k, v := range s.elements {
			if v.data != r.defaultValue {
				r.elements[k] = v.Commit()
			}
		}
	}
	return
}

//	Return a new SparseSlice header for the previous state of a given SparseSlice relative to its
//	current state.
//
//	If an existing cell currently contains the default value, we don't bother to reference this
//	in the new SparseSlice.
//
func (s *SparseSlice) Undo(steps int) (r *SparseSlice) {
	if steps < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	if s != nil {
		r = s.Rollback(s.currentVersion - steps)
	}
	return
}

//	Return a new SparseSlice header for the state of the SparseSlice at a given version point.
//
//	If an existing cell currently contains the default value, we don't bother to reference this
//	in the new SparseSlice.
//
func (s *SparseSlice) Rollback(version int) (r *SparseSlice) {
	if version < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	for ; s != nil && s.creationVersion > version; s = s.SparseSlice {}

	if s != nil {
		r = &SparseSlice{ elements: make(sliceHash), defaultValue: s.defaultValue, creationVersion: version, currentVersion: version }
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