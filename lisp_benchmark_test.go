package greenspun

import "testing"

var (
	List0 = List()
	List1 = List(0)
	List10 = Repeat(List(0), 10)
	List100 = Repeat(List(0), 100)
	List1000 = Repeat(List(0), 1000)
)

func BenchmarkLispLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(List0)
	}
}

func BenchmarkLispLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(List1)
	}
}

func BenchmarkLispLen10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(List10)
	}
}

func BenchmarkLispLen100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(List100)
	}
}

func BenchmarkLispLen1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(List1000)
	}
}

func BenchmarkLispIsListValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsList(false)
	}
}

func BenchmarkLispIsListList(b *testing.B) {
	b.StopTimer()
		l := List(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsList(l)
	}
}

func BenchmarkLispIsAtomList(b *testing.B) {
	b.StopTimer()
		l := List(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsAtom(l)
	}
}

func BenchmarkLispIsAtomValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAtom(false)
	}
}

func BenchmarkLispIsNilNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsNil(nil)
	}
}

func BenchmarkLispIsNilList(b *testing.B) {
	b.StopTimer()
		l := List(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsNil(l)
	}
}

func BenchmarkLispIsNilValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsNil(false)
	}
}

func BenchmarkLispEqual0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(List0, List0)
	}
}

func BenchmarkLispEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(List1, List1)
	}
}

func BenchmarkLispEqual10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(List10, List10)
	}
}

func BenchmarkLispEqual100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(List100, List100)
	}
}

func BenchmarkLispEqual1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(List1000, List1000)
	}
}

func BenchmarkLispCarNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Car(nil)
	}
}

func BenchmarkLispCarCell(b *testing.B) {
	b.StopTimer()
		c := Cons(nil, nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Car(c)
	}
}

func BenchmarkLispCdrNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cdr(nil)
	}
}

func BenchmarkLispCdrCell(b *testing.B) {
	b.StopTimer()
		c := Cons(nil, nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Cdr(c)
	}
}

func BenchmarkLispCaarNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Caar(nil)
	}
}

func BenchmarkLispCaarCell(b *testing.B) {
	b.StopTimer()
		c := Cons(Cons(nil, nil), nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Caar(c)
	}
}

func BenchmarkLispCadrNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cadr(nil)
	}
}

func BenchmarkLispCadrCell(b *testing.B) {
	b.StopTimer()
		c := Cons(Cons(nil, nil), nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Cadr(c)
	}
}

func BenchmarkLispCdarNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cdar(nil)
	}
}

func BenchmarkLispCdarCell(b *testing.B) {
	b.StopTimer()
		c := Cons(nil, Cons(nil, nil))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Cdar(c)
	}
}

func BenchmarkLispCddrNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cddr(nil)
	}
}

func BenchmarkLispCddrCell(b *testing.B) {
	b.StopTimer()
		c := Cons(nil, Cons(nil, nil))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Cddr(c)
	}
}

func BenchmarkLispOffset0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Offset(List0, 0)
	}
}

func BenchmarkLispOffset1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Offset(List1, 0)
	}
}

func BenchmarkLispOffset10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Offset(List10, 9)
	}
}

func BenchmarkLispOffset100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Offset(List100, 99)
	}
}

func BenchmarkLispOffset1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Offset(List1000, 999)
	}
}

func BenchmarkLispEnd0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		End(List0)
	}
}

func BenchmarkLispEnd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		End(List1)
	}
}

func BenchmarkLispEnd10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		End(List10)
	}
}

func BenchmarkLispEnd100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		End(List100)
	}
}

func BenchmarkLispEnd1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		End(List1000)
	}
}

func BenchmarkLispEach0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(List0, func(interface{}) {})
	}
}

func BenchmarkLispEach1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(List1, func(interface{}) {})
	}
}

func BenchmarkLispEach10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(List10, func(interface{}) {})
	}
}

func BenchmarkLispEach100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(List100, func(interface{}) {})
	}
}

func BenchmarkLispEach1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(List1000, func(interface{}) {})
	}
}

func BenchmarkLispMap0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(List0, func(interface{}) interface{} { return nil })
	}
}

func BenchmarkLispMap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(List1, func(interface{}) interface{} { return nil })
	}
}

func BenchmarkLispMap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(List10, func(interface{}) interface{} { return nil })
	}
}

func BenchmarkLispMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(List100, func(interface{}) interface{} { return nil })
	}
}

func BenchmarkLispMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(List1000, func(interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReduce0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(List0, nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReduce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(List1, nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReduce10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(List10, nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReduce100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(List100, nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReduce1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(List1000, nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkLispReverse0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(List0)
	}
}

func BenchmarkLispReverse1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(List1)
	}
}

func BenchmarkLispReverse10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(List10)
	}
}

func BenchmarkLispReverse100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(List100)
	}
}

func BenchmarkLispReverse1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(List1000)
	}
}

func BenchmarkLispCopy0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(List0)
	}
}

func BenchmarkLispCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(List1)
	}
}

func BenchmarkLispCopy10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(List10)
	}
}

func BenchmarkLispCopy100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(List100)
	}
}

func BenchmarkLispCopy1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(List1000)
	}
}

func BenchmarkLispRepeat0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(List1, 0)
	}
}

func BenchmarkLispRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(List1, 1)
	}
}

func BenchmarkLispRepeat10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(List1, 10)
	}
}

func BenchmarkLispRepeat100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(List1, 100)
	}
}

func BenchmarkLispRepeat1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(List1, 1000)
	}
}


/*
func TestAppend(t *testing.T) {
	ConfirmAppend := func(c LispPair, v interface{}, r interface{}) {
		cs := fmt.Sprintf("%v", c)
		if x := Append(c, v); !Equal(x, r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", cs, v, r, x)
		}
	}

	ConfirmAppend(List(), 1, List(1))
	ConfirmAppend(List(), List(1), List(1))
	ConfirmAppend(List(), List(1, 2), List(1, 2))
	ConfirmAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmAppend(List(1), 2, List(1, 2))
	ConfirmAppend(List(1), List(2), List(1, 2))
	ConfirmAppend(List(1), List(2, 3), List(1, 2, 3))

	ConfirmMultipleAppend := func(c LispPair, r interface{}, v... interface{}) {
		call := fmt.Sprintf("%v.Append(%v)", c, v)
		if x := Append(c, v...); !Equal(x, r) {
			t.Fatalf("%v should be %v but is %v", call, r, x)
		}
	}

	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1), List(2, 3))
	ConfirmMultipleAppend(List(), List(1, 2, 3, 4), List(1, 2), 3, List(4))
	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmMultipleAppend(List(1), List(1, 2), 2)
	ConfirmMultipleAppend(List(1), List(1, 2, 3), 2, 3)
	ConfirmMultipleAppend(List(1), List(1, 2), List(2))
	ConfirmMultipleAppend(List(1), List(1, 2, 3), List(2, 3))
}

func TestWhile(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c LispPair, l int, f interface{}) {
		if count := While(c, f); count != l {
			t.Fatalf("While(%v, %v) should have iterated %v times not %v times", c, l, l, count)
		}
	}

	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		return i != limit
	})

	limit = 6
	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	limit = 7
	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key != limit
	})
}

func TestUntil(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c LispPair, l int, f interface{}) {
		if count := Until(c, f); count != l {
			t.Fatalf("Until(%v, %v) should have iterated %v times not %v times", c, l, l, count)
		}
	}

	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		return i == limit
	})

	limit = 6
	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	limit = 7
	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key == limit
	})
}
*/