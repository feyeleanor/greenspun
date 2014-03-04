package greenspun

type sliceHash	map[int] *versionedValue

func denseSliceHash(values ...interface{}) (r sliceHash) {
	r = make(sliceHash)
	for i, v := range values {
		r[i] = &versionedValue{ data: v }
	}
	return
}

func (s sliceHash) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case sliceHash:
		if len(s) == len(o) {
			for k, vo := range o {
				if va, ok := s[k]; ok {
					r = va.Equal(vo)
				}
				if !r {
					break
				}
			}
		}
	case nil:
		r = s == nil
	}
	return
}