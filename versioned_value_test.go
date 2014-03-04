package greenspun

import "testing"

func TestArrayElementEqual(t *testing.T) {
	ConfirmEqual := func(a *versionedValue, v interface{}, r bool) {
		if x := a.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
		}
		if v, ok := v.(*versionedValue); ok {
			if x := v.Equal(a); x != r {
				t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
			}
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(&versionedValue{ data: nil }, nil, true)
	ConfirmEqual(&versionedValue{ data: nil }, &versionedValue{ data: nil }, true)
	ConfirmEqual(&versionedValue{}, nil, true)
	ConfirmEqual(&versionedValue{}, &versionedValue{}, true)

	ConfirmEqual(&versionedValue{ data: 1 }, nil, false)
	ConfirmEqual(&versionedValue{ data: 1 }, &versionedValue{}, false)
	ConfirmEqual(&versionedValue{ data: 1 }, 1, true)
	ConfirmEqual(&versionedValue{ data: 1 }, &versionedValue{ data: 1 }, true)

	ConfirmEqual(&versionedValue{ data: stack(0, 1) }, nil, false)
	ConfirmEqual(&versionedValue{ data: stack(0, 1) }, &versionedValue{}, false)
	ConfirmEqual(&versionedValue{ data: stack(0, 1) }, &versionedValue{ data: stack(0) }, false)
	ConfirmEqual(&versionedValue{ data: stack(0, 1) }, &versionedValue{ data: stack(0, 1) }, true)
}

func TestArrayElementCommit(t *testing.T) {
	ConfirmCommit := func(a, r *versionedValue) {
		if a == nil {
			if x := a.Commit(); x != nil {
				t.Fatalf("%v.Commit() should be nil but is %v", a, x)
			}
		} else {
			switch x := a.Commit(); {
			case !x.Equal(r):
				t.Fatalf("%v.Commit() should be %v but is %v", a, r, x)
			case x.versionedValue != nil:
				t.Fatalf("%v.Commit().versionedValue should be nil but is %v", a, x.versionedValue)
			case x.version != 0:
				t.Fatalf("%v.Commit().versionedValue should be 0 but is %v", a, x.version)
			}
		}
	}

	ConfirmCommit(nil, nil)
	ConfirmCommit(&versionedValue{}, &versionedValue{})
	ConfirmCommit(&versionedValue{ data: 0 }, &versionedValue{ data: 0 })
	ConfirmCommit(&versionedValue{ data: 0, version: 1 }, &versionedValue{ data: 0 })
	ConfirmCommit(&versionedValue{ data: 1, version: 1, versionedValue: &versionedValue{ data: 0 } }, &versionedValue{ data: 1 })
}

func TestArrayElementAtVersion(t *testing.T) {
	ConfirmAtVersion := func(a *versionedValue, v int, r *versionedValue) {
		if x := a.AtVersion(v); !x.Equal(r) {
			t.Fatalf("%v.AtVersion(%v) should be %v but is %v", a, v, r, x)
		}
	}

	ConfirmAtVersion(nil, 0, nil)
	ConfirmAtVersion(nil, 1, nil)

	ConfirmAtVersion(&versionedValue{}, 0, nil)
	ConfirmAtVersion(&versionedValue{}, 0, &versionedValue{})

	ConfirmAtVersion(&versionedValue{}, 1, nil)
	ConfirmAtVersion(&versionedValue{}, 1, &versionedValue{})

	ConfirmAtVersion(&versionedValue{ data: 0, version: 0 }, 0, &versionedValue{ data: 0, version: 0 })
	ConfirmAtVersion(&versionedValue{ data: 0, version: 0 }, 1, &versionedValue{ data: 0, version: 0 })

	ConfirmAtVersion(&versionedValue{ data: 0, version: 1 }, 1, &versionedValue{ data: 0, version: 1 })
	ConfirmAtVersion(&versionedValue{ data: 0, version: 1 }, 0, nil)

	ConfirmAtVersion(&versionedValue{ data: 1, version: 1, versionedValue: &versionedValue{ data: 0, version: 0 } }, 0, &versionedValue{ data: 0, version: 0 })
	ConfirmAtVersion(&versionedValue{ data: 1, version: 1, versionedValue: &versionedValue{ data: 0, version: 0 } }, 1, &versionedValue{ data: 1, version: 1, versionedValue: &versionedValue{ data: 0, version: 0 } })
}

func TestArrayElementUndo(t *testing.T) {
	ConfirmUndo := func(a, r *versionedValue) {
		switch x := a.Undo(); {
		case a == nil:
			if x != nil {
				t.Fatalf("%v.Undo() should be nil but is %v", a, x)
			}
		case !x.Equal(r):
			t.Fatalf("%v.Undo() should be %v but is %v", a, r, x)
		}
	}

	ConfirmUndo(nil, nil)
	ConfirmUndo(&versionedValue{ data: 0 }, &versionedValue{ data: 0 })
	ConfirmUndo(&versionedValue{ data: 1, versionedValue: &versionedValue{ data: 0 } }, &versionedValue{ data: 0 })
}

func TestArrayElementRollback(t *testing.T) {
	ConfirmRollback := func(a, r *versionedValue) {
		switch x := a.Rollback(); {
		case a == nil:
			if x != nil {
				t.Fatalf("%v.Undo() should be nil but is %v", a, x)
			}
		case !x.Equal(r):
			t.Fatalf("%v.Undo() should be %v but is %v", a, r, x)
		}
	}

	ConfirmRollback(nil, nil)
	ConfirmRollback(&versionedValue{ data: 0 }, &versionedValue{ data: 0 })
	ConfirmRollback(&versionedValue{ data: 0, version: 1 }, nil)
	ConfirmRollback(&versionedValue{ data: 1, version: 1, versionedValue: &versionedValue{ data: 0 } }, &versionedValue{ data: 0 })
	ConfirmRollback(&versionedValue{ data: 1, version: 2, versionedValue: &versionedValue{ data: 0, version: 1 } }, nil)
}