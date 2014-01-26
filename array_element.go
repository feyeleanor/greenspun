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