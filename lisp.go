package greenspun

type Lambda	struct {
	Formals		*Pair		"a list of formal parameters to be bound by the lambda"
	Body			*Pair		"the operations to be executed by the lambda"
}

func (l Lambda) Eval(e Environment) interface{} {
	return nil
}

type DespatchTable	map[string] func(*Pair) *Pair

var SpecialForms DespatchTable = DespatchTable{
	"lambda":		func(l *Pair) *Pair {
		largs := l.Cadr().(*Pair)
		lsexp := largs.Cdr().(*Pair)
		return Cons(&Lambda{ largs, lsexp }, nil)
	},
}


func Eval(v interface{}, e Environment) interface{} {
	switch v := v.(type) {
	case *Pair:
		switch {
		case v.IsNil():
			return nil
		case v.Car() == "lambda":
//		largs := Cadr(v).(*Pair)
//		lsexp := Cdr(largs).(*Pair)
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