package greenspun

import "testing"

func TestArrayElementEqual(t *testing.T) {
	ConfirmEqual := func(a *arrayElement, v interface{}, r bool) {
		if x := a.Equal(v); x != r {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
		}
		if v, ok := v.(*arrayElement); ok {
			if x := v.Equal(a); x != r {
				t.Fatalf("%v.Equal(%v) should be %v but is %v", a, v, r, x)
			}
		}
	}

	ConfirmEqual(nil, nil, true)
	ConfirmEqual(&arrayElement{ data: nil }, nil, true)
	ConfirmEqual(&arrayElement{ data: nil }, &arrayElement{ data: nil }, true)
	ConfirmEqual(&arrayElement{}, nil, true)
	ConfirmEqual(&arrayElement{}, &arrayElement{}, true)

	ConfirmEqual(&arrayElement{ data: 1 }, nil, false)
	ConfirmEqual(&arrayElement{ data: 1 }, &arrayElement{}, false)
	ConfirmEqual(&arrayElement{ data: 1 }, 1, true)
	ConfirmEqual(&arrayElement{ data: 1 }, &arrayElement{ data: 1 }, true)

	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, nil, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{}, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{ data: stack(0) }, false)
	ConfirmEqual(&arrayElement{ data: stack(0, 1) }, &arrayElement{ data: stack(0, 1) }, true)
}

func TestArrayElementCommit(t *testing.T) {
	ConfirmCommit := func(a, r *arrayElement) {
		if a == nil {
			if x := a.Commit(); x != nil {
				t.Fatalf("%v.Commit() should be nil but is %v", a, x)
			}
		} else {
			switch x := a.Commit(); {
			case !x.Equal(r):
				t.Fatalf("%v.Commit() should be %v but is %v", a, r, x)
			case x.arrayElement != nil:
				t.Fatalf("%v.Commit().arrayElement should be nil but is %v", a, x.arrayElement)
			case x.version != 0:
				t.Fatalf("%v.Commit().arrayElement should be 0 but is %v", a, x.version)
			}
		}
	}

	ConfirmCommit(nil, nil)
	ConfirmCommit(&arrayElement{}, &arrayElement{})
	ConfirmCommit(&arrayElement{ data: 0 }, &arrayElement{ data: 0 })
	ConfirmCommit(&arrayElement{ data: 0, version: 1 }, &arrayElement{ data: 0 })
	ConfirmCommit(&arrayElement{ data: 1, version: 1, arrayElement: &arrayElement{ data: 0 } }, &arrayElement{ data: 1 })
}

func TestArrayElementAtVersion(t *testing.T) {
	ConfirmAtVersion := func(a *arrayElement, v int, r *arrayElement) {
		if x := a.AtVersion(v); !x.Equal(r) {
			t.Fatalf("%v.AtVersion(%v) should be %v but is %v", a, v, r, x)
		}
	}

	ConfirmAtVersion(nil, 0, nil)
	ConfirmAtVersion(nil, 1, nil)

	ConfirmAtVersion(&arrayElement{}, 0, nil)
	ConfirmAtVersion(&arrayElement{}, 0, &arrayElement{})

	ConfirmAtVersion(&arrayElement{}, 1, nil)
	ConfirmAtVersion(&arrayElement{}, 1, &arrayElement{})

	ConfirmAtVersion(&arrayElement{ data: 0, version: 0 }, 0, &arrayElement{ data: 0, version: 0 })
	ConfirmAtVersion(&arrayElement{ data: 0, version: 0 }, 1, &arrayElement{ data: 0, version: 0 })

	ConfirmAtVersion(&arrayElement{ data: 0, version: 1 }, 1, &arrayElement{ data: 0, version: 1 })
	ConfirmAtVersion(&arrayElement{ data: 0, version: 1 }, 0, nil)

	ConfirmAtVersion(&arrayElement{ data: 1, version: 1, arrayElement: &arrayElement{ data: 0, version: 0 } }, 0, &arrayElement{ data: 0, version: 0 })
	ConfirmAtVersion(&arrayElement{ data: 1, version: 1, arrayElement: &arrayElement{ data: 0, version: 0 } }, 1, &arrayElement{ data: 1, version: 1, arrayElement: &arrayElement{ data: 0, version: 0 } })
}

func TestArrayElementUndo(t *testing.T) {
	ConfirmUndo := func(a, r *arrayElement) {
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
	ConfirmUndo(&arrayElement{ data: 0 }, &arrayElement{ data: 0 })
	ConfirmUndo(&arrayElement{ data: 1, arrayElement: &arrayElement{ data: 0 } }, &arrayElement{ data: 0 })
}

func TestArrayElementRollback(t *testing.T) {
	ConfirmRollback := func(a, r *arrayElement) {
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
	ConfirmRollback(&arrayElement{ data: 0 }, &arrayElement{ data: 0 })
	ConfirmRollback(&arrayElement{ data: 0, version: 1 }, nil)
	ConfirmRollback(&arrayElement{ data: 1, version: 1, arrayElement: &arrayElement{ data: 0 } }, &arrayElement{ data: 0 })
	ConfirmRollback(&arrayElement{ data: 1, version: 2, arrayElement: &arrayElement{ data: 0, version: 1 } }, nil)
}