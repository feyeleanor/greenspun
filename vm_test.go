package greenspun

import (
	"fmt"
	"testing"
)

func TestVMLocate(t *testing.T) {
	vm := NewVM(
					List(List(0, 1, 2),
						List(10, 11, 12),
						List(20, 21, 22),
						),
					nil,
	)

	ConfirmLocate := func(v *VM, e, s int, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		if x := vm.Locate(e, s); x != r {
			t.Fatalf("%v.Locate(%v, %v) should be %v but is %v", vs, e, s, r, x)
		}
	}

	ConfirmLocate(vm, 0, 0, 0)
	ConfirmLocate(vm, 0, 1, 1)
	ConfirmLocate(vm, 0, 2, 2)

	ConfirmLocate(vm, 1, 0, 10)
	ConfirmLocate(vm, 1, 1, 11)
	ConfirmLocate(vm, 1, 2, 12)

	ConfirmLocate(vm, 2, 0, 20)
	ConfirmLocate(vm, 2, 1, 21)
	ConfirmLocate(vm, 2, 2, 22)
}

func TestVMAdvance(t *testing.T) {
	vm := NewVM(nil, List(0, 1, 2, 3, 4, 5))

	ConfirmAdvance := func(v *VM, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		if vm.Advance(); vm.C.head != r {
			t.Fatalf("%v.Advance() should be %v but is %v", vs, r, vm.C.head)
		}
	}

	ConfirmAdvance(vm, 1)
	ConfirmAdvance(vm, 2)
	ConfirmAdvance(vm, 3)
	ConfirmAdvance(vm, 4)
	ConfirmAdvance(vm, 5)
}

func TestVMRun(t *testing.T) {
	t.Logf("Implement tests for VM::Run()")
}

func TestVMNil(t *testing.T) {
	vm := NewVM(nil, List(NIL, NIL, NIL, NIL))
	
	ConfirmNil := func(v *VM, l int) {
		vs := fmt.Sprintf("%v", v)
		switch v.Nil(); {
		case v.S.Peek() != nil:
			t.Fatalf("%v.Nil() the head should be nil but is %v", vs, v.S.Peek())
		case v.S.Len() != l:
			t.Fatalf("%v.Nil() should result in a stack of %v elements but has %v elements", vs, l, v.S.Len())
		}
	}

	ConfirmNil(vm, 1)
	ConfirmNil(vm, 2)
	ConfirmNil(vm, 3)
	ConfirmNil(vm, 4)
}

func TestVMLdc(t *testing.T) {
	vm := NewVM(
					nil,
					List(
						Cons(LDC, 0),
						Cons(LDC, 1),
						Cons(LDC, 2),
						Cons(LDC, 3),
					),
	)

	ConfirmLdc := func(v *VM, l int, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		switch v.Ldc(); {
		case v.S.Peek() != r:
			t.Fatalf("%v.Ldc() the head should be %v but is %v", vs, r, v.S.Peek())
		case v.S.Len() != l:
			t.Fatalf("%v.Ldc() should result in a stack of %v elements but has %v elements", vs, l, v.S.Len())
		}
	}

	ConfirmLdc(vm, 1, 0)
	ConfirmLdc(vm, 2, 1)
	ConfirmLdc(vm, 3, 2)
	ConfirmLdc(vm, 4, 3)
}

func TestVMLd(t *testing.T) {
	ConfirmLd := func(v *VM, l int, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		switch v.Ld(); {
		case v.S.Peek() != r:
			t.Fatalf("%v.Ld() the head should be %v but is %v", vs, r, v.S.Peek())
		case v.S.Len() != l:
			t.Fatalf("%v.Ld() should result in a stack of %v elements but has %v elements", vs, l, v.S.Len())
		}
	}

	env := List(List(0, 1, 2),
							List(10, 11, 12),
							List(20, 21, 22),
							)
	
	ConfirmLd(&VM{ S: Stack(), E: env, C: List(LD, Cons(0, 0)) }, 1, 0)

	vm := NewVM(
					env,
					List(	LD, Cons(0, 0),
								LD, Cons(0, 1),
								LD, Cons(0, 2),
								LD, Cons(1, 0),
								LD, Cons(1, 1),
								LD, Cons(1, 2),
								LD, Cons(2, 0),
								LD, Cons(2, 1),
								LD, Cons(2, 2),
					),
	)
	ConfirmLd(vm, 1, 0)
	ConfirmLd(vm, 2, 1)
	ConfirmLd(vm, 3, 2)
	ConfirmLd(vm, 4, 10)
	ConfirmLd(vm, 5, 11)
	ConfirmLd(vm, 6, 12)
	ConfirmLd(vm, 7, 20)
	ConfirmLd(vm, 8, 21)
	ConfirmLd(vm, 9, 22)
}

func TestVMLdf(t *testing.T) {
	t.Logf("Implement tests for VM::Ldf()")
}

func TestVMSel(t *testing.T) {
	t.Logf("Implement tests for VM::Sel()")
}

func TestVMJoin(t *testing.T) {
	t.Logf("Implement tests for VM::Join()")
}

func TestVMAp(t *testing.T) {
	t.Logf("Implement tests for VM::Ap()")
}

func TestVMRet(t *testing.T) {
	t.Logf("Implement tests for VM::Ret()")
}

func TestVMDum(t *testing.T) {
	t.Logf("Implement tests for VM::Dum()")
}

func TestVMRap(t *testing.T) {
	t.Logf("Implement tests for VM::Rap()")
}

func TestVMScar(t *testing.T) {
	RefuteScar := func(v *VM) {
		vs := fmt.Sprintf("%v", v)
		defer ConfirmPanic(t, "%v.Scar() should panic", vs)()
		v.Scar()
	}

//	ConfirmScar := func(v *VM, r interface{}) {
//		vs := fmt.Sprintf("%v", v)
//		if v.Scar(); v.S.Peek() != r {
//			t.Fatalf("%v.Scar() should be %v but is %v", vs, r, v.S.Peek())
//		}
//	}

//	RefuteScar(&VM{})
//	RefuteScar(&VM{ S: Stack() })
//	RefuteScar(&VM{ S: Stack(0) })
	RefuteScar(&VM{ S: Stack(List()) })

//	ConfirmScar(&VM{ S: Stack(List(0)) }, 0)
//	ConfirmScar(&VM{ S: Stack(List(0, 1)) }, 0)
//	ConfirmScar(&VM{ S: Stack(List(0, 1, 2)) }, 0)
}

func TestVMScdr(t *testing.T) {
	ConfirmScdr := func(v *VM, r *Lifo) {
		vs := fmt.Sprintf("%v", v)
		if v.Scdr(); !r.Equal(v.S) {
			t.Fatalf("%v.Scdr() should be %v but is %v", vs, r, v.S)
		}
	}

	ConfirmScdr(&VM{ S: Stack(Cons(0, 1), 2), C: List(SCDR) }, Stack(1, 2))
	ConfirmScdr(&VM{ S: Stack(Cons(1, 2), 3), C: List(SCDR) }, Stack(2, 3))
	ConfirmScdr(&VM{ S: Stack(Cons(2, 3), 4), C: List(SCDR) }, Stack(3, 4))
}

func TestVMScons(t *testing.T) {
	ConfirmScons := func(v *VM, r *Lifo) {
		vs := fmt.Sprintf("%v", v)
		if v.Scons(); !r.Equal(v.S) {
			t.Fatalf("%v.Scons() should be %v but is %v", vs, r, v.S)
		}
	}

	ConfirmScons(&VM{ S: Stack(0, 1, 2, 3), C: List(SCONS) }, Stack(Cons(0, 1), 2, 3))
	ConfirmScons(&VM{ S: Stack(Cons(0, 1), 2, 3), C: List(SCONS) }, Stack(Cons(Cons(0, 1), 2), 3))
	ConfirmScons(&VM{ S: Stack(Cons(Cons(0, 1), 2), 3), C: List(SCONS) }, Stack(Cons(Cons(Cons(0, 1), 2), 3)))
}

func TestVMSeq(t *testing.T) {
	ConfirmSeq := func(v *VM, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		if v.Seq(); v.S.Peek() != r {
			t.Fatalf("%v.Seq() should be %v but is %v", vs, r, v.S.Peek())
		}
	}

	ConfirmSeq(&VM{ S: Stack(0, 0), C: List(SEQ) }, TRUE)
	ConfirmSeq(&VM{ S: Stack(0, 1), C: List(SEQ) }, nil)
}