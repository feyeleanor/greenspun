package greenspun

type arrayHash	map[int] *arrayElement

func denseArrayHash(values ...interface{}) (r arrayHash) {
	r = make(arrayHash)
	for i, v := range values {
		r[i] = &arrayElement{ data: v }
	}
	return
}

func (s arrayHash) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case arrayHash:
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