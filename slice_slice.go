package greenspun

type sliceSlice	[]*versionedValue

func denseSlice(values ...interface{}) (r sliceSlice) {
	r = make(sliceSlice, len(values))
	for i, v := range values {
		r[i] = &versionedValue{ data: v }
	}
	return
}

func (s sliceSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case sliceSlice:
		if len(s) == len(o) {
			for i, v := range o {
				if r = s[i].Equal(v); !r {
					break
				}
			}
		}
	case nil:
		r = s == nil
	}
	return
}