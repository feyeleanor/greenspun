package greenspun

import (
	"fmt"
	"testing"
)

func TestVMLocate(t *testing.T) {
	vm := &VM{
		E: List(List(0, 1, 2),
						List(10, 11, 12),
						List(20, 21, 22),
						),
	}

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
	vm := &VM{ C: List(0, 1, 2, 3, 4, 5) }

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
	vm := &VM{ S: List() }
	
	ConfirmNil := func(v *VM, l int) {
		vs := fmt.Sprintf("%v", v)
		switch v.Nil(); {
		case v.S.Car() != nil:
			t.Fatalf("%v.Nil() the head should be nil but is %v", vs, v.S.Car())
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
	vm := &VM{ C: List(
									Cons(LDC, 0),
									Cons(LDC, 1),
									Cons(LDC, 2),
									Cons(LDC, 3),
								),
	}

	ConfirmLdc := func(v *VM, l int, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		switch v.Ldc(); {
		case v.S.Car() != r:
			t.Fatalf("%v.Ldc() the head should be %v but is %v", vs, r, v.S.Car())
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
	vm := &VM{
							E: List(List(0, 1, 2),
											List(10, 11, 12),
											List(20, 21, 22),
											),
							C: List(LD, Cons(0, 0),
											LD, Cons(0, 1),
											LD, Cons(0, 2),
											LD, Cons(1, 0),
											LD, Cons(1, 1),
											LD, Cons(1, 2),
											LD, Cons(2, 0),
											LD, Cons(2, 1),
											LD, Cons(2, 2),
								),
	}

	ConfirmLd := func(v *VM, l int, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		switch v.Ld(); {
		case v.S.Car() != r:
			t.Fatalf("%v.Ld() the head should be %v but is %v", vs, r, v.S.Car())
		case v.S.Len() != l:
			t.Fatalf("%v.Ld() should result in a stack of %v elements but has %v elements", vs, l, v.S.Len())
		}
	}

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

func TestVMCar(t *testing.T) {
	ConfirmCar := func(v *VM, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		if v.Car(); v.S.head != r {
			t.Fatalf("%v.Car() should be %v but is %v", vs, r, v.S.head)
		}
	}

	ConfirmCar(&VM{ S: List(List(0)) }, 0)
	ConfirmCar(&VM{ S: List(List(1)) }, 1)
	ConfirmCar(&VM{ S: List(List(2)) }, 2)
}

func TestVMCdr(t *testing.T) {
	ConfirmCdr := func(v *VM, r *Cell) {
		vs := fmt.Sprintf("%v", v)
		if v.Cdr(); !r.Equal(v.S) {
			t.Fatalf("%v.Cdr() should be %v but is %v", vs, r, v.S)
		}
	}

	ConfirmCdr(&VM{ S: List(Cons(0, 1), 2) }, List(1, 2))
	ConfirmCdr(&VM{ S: List(Cons(1, 2), 3) }, List(2, 3))
	ConfirmCdr(&VM{ S: List(Cons(2, 3), 4) }, List(3, 4))
}

func TestVMCons(t *testing.T) {
	vm := &VM{ S: List(0, 1, 2, 3) }
	ConfirmCons := func(v *VM, r *Cell) {
		vs := fmt.Sprintf("%v", v)
		if v.Cons(); !r.Equal(v.S) {
			t.Fatalf("%v.Cons() should be %v but is %v", vs, r, v.S)
		}
	}

	ConfirmCons(vm, List(Cons(0, 1), 2, 3))
	ConfirmCons(vm, List(Cons(Cons(0, 1), 2), 3))
	ConfirmCons(vm, List(Cons(Cons(Cons(0, 1), 2), 3)))
}

func TestVMEq(t *testing.T) {
	ConfirmCons := func(v *VM, r interface{}) {
		vs := fmt.Sprintf("%v", v)
		if v.Eq(); v.S.head != r {
			t.Fatalf("%v.Eq() should be %v but is %v", vs, r, v.S.head)
		}
	}

	ConfirmCons(&VM{ S: List(0, 0) }, "T")
	ConfirmCons(&VM{ S: List(0, 1) }, nil)
}