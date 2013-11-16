package greenspun

//import "fmt"

//	This type embodies the core of the SECD virtual machine for implementing functional languages
//
type VM struct {
	S		*StackList	"stack"
	E		*Pair				"environment"
	C		*Pair				"control"
	D		StackList		"dump"
}

func NewVM(env, code *Pair) *VM {
	return &VM{ S: Stack(), E: env, C: code }
}

//	Initialise the virtual machine with a global environment and code for execution
//
func (vm *VM) Initialize(env, code *Pair) {
//	vm.S = nil
	vm.E = env
	vm.C = code
//	vm.D = nil
}


//	Load a value from a slot in the environment
//
func (vm *VM) Locate(env, slot int) interface{} {
	return vm.E.Offset(env).Car().(*Pair).Offset(slot).Car()
}


//	Advance the control register to the next instruction
func (vm *VM) Advance() {
	vm.C = vm.C.Next()
}


//	Execute the program
//	If a panic occurs in the running program we catch it here and signal it with a boolean false.
func (vm *VM) Run() bool {
	defer func() {
		recover()
	}()
	for ; !vm.C.IsNil(); {
		switch vm.C.head {
		case NIL:
			vm.Nil()
		case LDC:
			vm.Ldc()
		case LD:
			vm.Ld()
		case LDF:
			vm.Ldf()
		case SEL:
			vm.Sel()
		case JOIN:
			vm.Join()
		case AP:
			vm.Ap()
		case RET:
			vm.Ret()
		case DUM:
			vm.Dum()
		case RAP:
			vm.Rap()
		case SCAR:
			vm.SCar()
		case SCDR:
			vm.SCdr()
		case SCONS:
			vm.SCons()
		case SEQ:
			vm.SEq()
		case HALT:
			break
		default:
			panic("unknown instruction")
		}
	}
	return true
}


//	NIL - push a nil pointer onto the stack
//		s						-> (nil . s)
//		(NIL . c)		-> c
func (vm *VM) Nil() {
	vm.S.Push(nil)
	vm.Advance()
}


//	LDC - push a constant value onto the stack
//		s						-> (x . s)
//		(LDC x . c)	-> c
func (vm *VM) Ldc() {
	vm.S.Push(vm.C.Cadr())
	vm.Advance()
}


//	LD - push the value of a variable onto the stack. The variable is indicated by the argument, a pair.
//	The pair's car specifies the level, the cdr the position. So "(1 . 3)" gives the current function's (level 1) third parameter.
//		s									-> (locate((i . j), e) . s)
//		(LD (i . j) . c)	-> c
func (vm *VM) Ld() {
	vm.Advance()
	env, slot := vm.C.Car().(*Pair).IntPair()
	vm.S.Push(vm.Locate(env, slot))
	vm.Advance()
}


//	LDF - takes one list argument representing a function. It constructs a closure (a pair containing the function and the current environment)
//	and pushes that onto the stack.
//		s									-> ((f . e) . s)
//		(LDF f . c)				-> c
func (vm *VM) Ldf() {
	vm.Advance()
	vm.S.Push(Cons(vm.C.Car(), vm.E))
	vm.Advance()
}


//	SEL - expects two list arguments, and pops a value from the stack. The first list is executed if the popped value was non-nil, the
//	second list otherwise. Before one of these list pointers is made the new C, a pointer to the instruction following sel is saved on the dump.
//		(x . s)						-> s
//		(SEL ct cf . c)		-> ct if x is T, or cf if x is F
//		d									-> (c . d)
func (vm *VM) Sel() {
	vm.D.Push(vm.C.Cddr().(*Pair).Cdr())
	if vm.S.Pop() == TRUE {
		vm.C = vm.C.Cadr().(*Pair)
	} else {
		vm.C = vm.C.Cadr().(*Pair).Cdr().(*Pair)
	}
}


//	JOIN - pop a list reference from the dump and make this the new value of C. This instruction occurs at the end of both alternatives of a sel.
//		(JOIN . c)				-> cr
//		(cr . d)					-> d
func (vm *VM) Join() {
	vm.C = vm.D.Pop().(*Pair)
}


//	AP - pop a closure and a list of parameter values from the stack. The closure is applied to the parameters by installing its environment as
//	the current one, pushing the parameter list in front of that, clearing the stack, and setting C to the closure's function pointer. The
//	previous values of S, E, and the next value of C are saved on the dump.
//		((f . e') v . s)	-> NIL
//		e									-> (v . e')
//		(AP . c)					-> f
//		d									-> (s e c . d)
func (vm *VM) Ap() {
	closure := vm.S.Pop().(*Pair)
	vm.E = Cons(closure.Cdr(), vm.S.Pop().(*Pair))
	vm.C = closure.Car().(*Pair)

	vm.D.Push(vm.C.Cdr())
	vm.D.Push(vm.E)
	vm.D.Push(vm.S)
	vm.S = Stack()
}


//	Ret - Pop one return value from the stack, restore S, E, and C from the dump, and push the return value onto the now-current stack.
//		(x . z)						-> (x . s)
//		e'								-> e
//		(RTN . q)					-> c
//		(s e c . d)				-> d
func (vm *VM) Ret() {
	x := vm.S.Top()
	vm.S = vm.D.Top().(*StackList)
	vm.S.Push(x)
	vm.E = vm.D.Pop().(*Pair)
	vm.C = vm.D.Pop().(*Pair)
}


//	Dum - push a "dummy", an empty list, in front of the environment list.
//
//		e (DUM.c) d			-> (W . e)
//		           					where W has been called PENDING earlier
//		(DUM . c)				-> c
//
func (vm *VM) Dum() {
	vm.E = Cons(List(), vm.E)
	vm.Advance()
}


//		Works like ap, only that it replaces an occurrence of a dummy environment with the current one, thus making recursive functions possible.
//		Pop a closure and a list of parameter values from the stack. The closure is applied to the parameters by installing its environment as
//		the current one, pushing the parameter list in front of that, clearing the stack, and setting C to the closure's function pointer. The
//		previous values of S, E, and the next value of C are saved on the dump.
//
//		((f . (W . e)) v . s)	-> nil
//		(W . e)								-> rplaca((W . e), v)
//		(RAP . c)							-> f
//		d											-> (s e c . d)
//
func (vm *VM) Rap() {
	closure := vm.S.Pop().(*Pair)
	vm.E = vm.S.Pop().(*Pair)
	vm.E.Rplaca(closure.Cdr())
	vm.C = closure.Car().(*Pair)

	vm.D.Push(vm.C.Cdr())
	vm.D.Push(vm.E.Cdr())
	vm.D.Push(vm.S)
	vm.S = Stack()
}


//		Replace TOS with Car(TOS)
//
//		((a . b) . s)		-> (a . s)
//		(CAR . c) 			-> c
//
func (vm *VM) SCar() {
	switch t := vm.S.Top().(type) {
	case nil:
	case *Pair:
		vm.S.Replace(t.Car())
	default:
		panic(vm)
	}
	vm.Advance()
}


//		Replace TOS with Cdr(TOS)
//
//		((a . b) . s) 	-> (b . s)
//		(CDR . c)				-> c
//
func (vm *VM) SCdr() {
	switch t := vm.S.Top().(type) {
	case nil:
	case *Pair:
		vm.S.Replace(t.Cdr())
	default:
		panic(vm)
	}
	vm.Advance()
}


//		Pop top two items of stack, combine them into a Pair and push onto stack
//
//		(a b . s)				-> ((a . b) . s)
//		(CONS . c)			-> c
//
func (vm *VM) SCons() {
	vm.S.Replace(Cons(vm.S.Pop(), vm.S.Top()))
	vm.Advance()
}


//		Push EQ of TOS and Cdr(TOS) onto the stack
//
//		(a a . s)				-> (#t . s)
//		(a b . s)				-> (#f . s)
//		(EQ . c) 				-> c
//
func (vm *VM) SEq() {
	if vm.S.Pop() == vm.S.Top() {
		vm.S.Replace(TRUE)
	} else {
		vm.S.Replace(nil)
	}
	vm.Advance()
}