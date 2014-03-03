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

//	Commit returns a new arrayElement containing the data in the current array element.
func (a *arrayElement) Commit() (r *arrayElement) {
	if a != nil {
		r = &arrayElement{ data: a.data }
	}
	return
}

//	AtVersion returns the value of the element when the given version number was current.
func (a *arrayElement) AtVersion(v int) (r *arrayElement) {
	for ; a != nil; a = a.arrayElement {
		if a.version <= v {
			return a
		}
	}
	return
}

//	Undo returns the previous value for the current array element.
//	When the end of the chain is reached it returns nil if the version is greater than
//	zero, otherwise it returns the cell which has version zero.
func (a *arrayElement) Undo() (r *arrayElement) {
	if a != nil {
		switch {
		case a.arrayElement != nil:
			r = a.arrayElement
		case a.version == 0:
			r = a
		}
	}
	return
}

//	Rollback returns the original value for the array element.
//	If the last element of the chain has version > 0 a nil is returned, otherwise the
//	last element is returned.
func (a *arrayElement) Rollback() (r *arrayElement) {
	for r = a ; r != nil && r.version > 0; r = r.arrayElement {}
	return
}