package greenspun

type Lambda	struct {
	Formals		*Cell
	Body			*Cell
}

func (l Lambda) Eval(e Environment) interface{} {
	return nil
}

type DespatchTable	map[string] func(*Cell) *Cell

var SpecialForms DespatchTable = DespatchTable{
	"lambda":		func(l *Cell) *Cell {
		largs := l.Cadr().(*Cell)
		lsexp := largs.Cdr().(*Cell)
		return Cons(&Lambda{ largs, lsexp }, nil)
	},
}


func Eval(v interface{}, e Environment) interface{} {
	switch v := v.(type) {
	case *Cell:
		switch {
		case v.IsNil():
			return nil
		case v.Car() == "lambda":
//		largs := Cadr(v).(*Cell)
//		lsexp := Cdr(largs).(*Cell)
//		return lambda(largs, lsexp)
			panic("no support for lambda at this time")
		default:
			return v.Map(func(x interface{}) interface{} {
				return(Eval(x, e))
			})
		}
	default:
		if val := e.Find(v); val != nil {
			return val
		}
	}
	return v
}