package greenspun

/*
	The DenseSlice is a persistent integer-indexed data store. Internally elements are slice
	which provides uniform access, and each access is represented by a stack of versioned values.

	The DenseSlice contains a creationVersion field which is initialised when a new header is
	created, and a currentVersion which is incremented whenever an element is modified.

	Queries out of bounds will panic as with a conventional slice.
*/
type DenseSlice struct {
	elements					sliceSlice			"versioned elements"
	creationVersion		int							"the version number for which this header was created"
	currentVersion		int							"the version number of the most recent change"
	*DenseSlice												"the DenseSlice from which the current header is derived"
}

/*
	NewDenseSlice returns a DenseSlice. If optional parameters are provided these will result in values being
	assigned to the initial cells and if necessary the length of the DenseSlice adjusted to reflect this.
*/
func NewDenseSlice(n int, items ...interface{}) (r *DenseSlice) {
	if n < len(items) {
		n = len(items)
	}
	r = &DenseSlice{ elements: make(sliceSlice, n) }
	for i, v := range items {
		r.elements[i] = &versionedValue{ data: v }
	}
	return
}

func (s *DenseSlice) copyHeader() (r *DenseSlice) {
	if s != nil {
		r = &DenseSlice{ s.elements, s.creationVersion, s.currentVersion, s.DenseSlice }
	}
	return
}

func (s *DenseSlice) newHeader() (r *DenseSlice) {
	if s != nil {
		creationVersion := s.currentVersion + 1
		r = &DenseSlice{ s.elements, creationVersion, creationVersion, s }
	}
	return
}

/*
func (s *DenseSlice) String() (r string) {
	return
}
*/

/*
	Return the current length of the slice.
*/
func (s *DenseSlice) Len() (r int) {
	if s != nil {
		r = len(s.elements)
	}
	return
}

/*
	Determine if a passed value is equivalent to the current slice.
*/
func (s *DenseSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case DenseSlice:
		r = s.Equal(&o)
	case *DenseSlice:
		switch {
		case s == nil && o == nil:
			r = true
		case s != nil && o == nil, s == nil && o != nil:
			r = false
		case len(s.elements) != len(o.elements):
			r = false
		case len(s.elements) == 0:
			r = true
		default:
			for i, v := range o.elements {
				if r = v.Equal(s.elements[i]); !r {
					break
				}
			}
		}
	case nil:
		r = s == nil
	}
	return
}

/*
	At returns the value stored at a particular location in the slice, which may be the default value or
	else a value which has been explicitly set. In the event that the location is not within the bounds
	specified by the slice a panic is raised with ARGUMENT_OUT_OF_BOUNDS.
*/
func (s *DenseSlice) At(i int) (r interface{}) {
	if s == nil || i < 0 || i >= len(s.elements) {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	if e := s.elements[i]; e != nil {
		r = e.data
	}
	return
}

/*
	For Set() operations we preserve the slice header so long as the length of the slice doesn't change.
*/
func (s *DenseSlice) Set(i int, v interface{}) (r *DenseSlice) {
	switch {
	case i < 0:
		panic(ARGUMENT_OUT_OF_BOUNDS)
	case s == nil:
		r = NewDenseSlice(0, nil)
	default:
		switch {
		case i >= len(s.elements):
			r = s.newHeader()
			e := make(sliceSlice, i)
			copy(e, r.elements)
			r.elements = e
		default:
			r = s
			r.currentVersion++
		}
	}

	r.elements[i] = &versionedValue{ data: v, version: r.currentVersion, versionedValue: r.elements[i] }
	return
}

/*
	Iterate through all cells in order, applying the supplied closure to the current cell value.
*/
func (s *DenseSlice) Each(f interface{}) {
	defer RescueOutOfBounds()
	var i	int

	switch f := f.(type) {
	case func():
		for i = len(s.elements); i > 0; i-- {
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

/*
	Return a new header in which the specified range of cells has been moved to the specified
	target locations.
*/
func (s *DenseSlice) Move(x, y, n int) (r *DenseSlice) {
	if s != nil {
		r = s.copyHeader()
		for n > 0 {
			n--
			r.Set(y + n, r.At(x + n))
			r.elements[x + n] = nil
		}
	}
	return
}

/*
	Create a new DenseSlice header referencing the cells in the existing DenseSlice but with
	the specified items inserted.

	If an existing cell currently contains the default value, we don't bother to reference this
	in the new DenseSlice. Likewise if one of the values to be inserted is the same as the
	default value then we don't bother creating a DenseSlice entry.
*/
func (s *DenseSlice) Insert(i int, items... interface{}) (r *DenseSlice) {
	if i < 0 {
		panic(ARGUMENT_NEGATIVE_INDEX)
	}

	n := len(items)

	if s == nil {
		r = NewDenseSlice(i + n, nil)		
	} else {
		if i < len(s.elements) {
			r = NewDenseSlice(len(s.elements) + n)
		} else {
			r = NewDenseSlice(i + n)
		}
		r.creationVersion = s.currentVersion + 1
		r.currentVersion = r.creationVersion

		copy(r.elements[:i], s.elements[:i])
		copy(r.elements[i + n:], s.elements[i:])
	}

	for k, v := range items {
		r.elements[i + k] = &versionedValue{ data: v, version: r.currentVersion }
	}
	return
}

/*
	Return a new DenseSlice header referencing the cells in the existing DenseSlice but with
	one or more cells removed.

	If an existing cell currently contains the default value, we don't bother to reference this
	in the new DenseSlice.
*/
func (s *DenseSlice) Delete(i int, params ...int) (r *DenseSlice) {
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
		case i + n > len(s.elements):
			n = len(s.elements) - i
		}
		if n > 0 {
			r = &DenseSlice{ elements: make(sliceSlice, len(s.elements) - n), currentVersion: s.currentVersion + 1 }
			copy(r.elements, s.elements[:i])
			copy(r.elements[i:], s.elements[i + n:])
		}
	}
	return
}

/*
	Return a new DenseSlice header referencing the cells in the existing DenseSlice.

	If an existing cell currently contains the default value, we don't bother to reference this
	in the new DenseSlice.
*/
func (s *DenseSlice) Copy() (r *DenseSlice) {
	if s != nil {
		r = NewDenseSlice(len(s.elements))
		copy(r.elements, s.elements)
	}
	return
}

/*
	Commit creates a new DenseSlice which is identical to the current state of a given DenseSlice.
*/
func (s *DenseSlice) Commit() (r *DenseSlice) {
	if s != nil {
		r = NewDenseSlice(len(s.elements))
		for k, v := range s.elements {
			r.elements[k] = v.Commit()
		}
	}
	return
}

/*
	Return a new DenseSlice header for the previous state of a given DenseSlice relative to its
	current state.

	If an existing cell currently contains the default value, we don't bother to reference this
	in the new DenseSlice.
*/
func (s *DenseSlice) Undo(steps int) (r *DenseSlice) {
	if steps < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	if s != nil {
		r = s.Rollback(s.currentVersion - steps)
	}
	return
}

/*
	Return a new DenseSlice header for the state of the DenseSlice at a given version point.

	If an existing cell currently contains the default value, we don't bother to reference this
	in the new DenseSlice.
*/
func (s *DenseSlice) Rollback(version int) (r *DenseSlice) {
	if version < 0 {
		panic(ARGUMENT_OUT_OF_BOUNDS)
	}

	for ; s != nil && s.creationVersion > version; s = s.DenseSlice {}

	if s != nil {
		r = &DenseSlice{ elements: make(sliceSlice, 0), creationVersion: version, currentVersion: version }
		for _, v := range s.elements {
			if x := v.AtVersion(version); x != nil {
				r.elements = append(r.elements, x)
			}
		}
	}
	return
}