package greenspun

import "testing"

var (
	List0 = List()
	List1 = List(0)
	List10 = List(0).Repeat(10)
	List100 = List(0).Repeat(100)
	List1000 = List(0).Repeat(1000)
	ListInterface interface{} = List0
)

func BenchmarkCellAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListInterface.(*Cell)
	}
}

func BenchmarkCellString0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.String()
	}
}

func BenchmarkCellString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.String()
	}
}

func BenchmarkCellString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.String()
	}
}

func BenchmarkCellString100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.String()
	}
}

func BenchmarkCellString1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.String()
	}
}
func BenchmarkCellLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Len()
	}
}

func BenchmarkCellLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Len()
	}
}

func BenchmarkCellLen10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Len()
	}
}

func BenchmarkCellLen100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Len()
	}
}

func BenchmarkCellLen1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Len()
	}
}

func BenchmarkCellIsNil0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.IsNil()
	}
}

func BenchmarkCellIsNil1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.IsNil()
	}
}

func BenchmarkCellEqual0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Equal(List0)
	}
}

func BenchmarkCellEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Equal(List1)
	}
}

func BenchmarkCellEqual10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Equal(List10)
	}
}

func BenchmarkCellEqual100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Equal(List100)
	}
}

func BenchmarkCellEqual1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Equal(List1000)
	}
}

func BenchmarkCellCar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Car()
	}
}

func BenchmarkCellCar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Car()
	}
}

func BenchmarkCellCdr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cdr()
	}
}

func BenchmarkCellCdr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cdr()
	}
}

func BenchmarkCellCaar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Caar()
	}
}

func BenchmarkCellCaar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Caar()
	}
}

func BenchmarkCellCaar10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Caar()
	}
}

func BenchmarkCellCaar100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Caar()
	}
}

func BenchmarkCellCaar1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Caar()
	}
}

func BenchmarkCellCadr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cadr()
	}
}

func BenchmarkCellCadr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cadr()
	}
}

func BenchmarkCellCadr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cadr()
	}
}

func BenchmarkCellCadr100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cadr()
	}
}

func BenchmarkCellCadr1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cadr()
	}
}

func BenchmarkCellCdar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cdar()
	}
}

func BenchmarkCellCdar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cdar()
	}
}

func BenchmarkCellCdar10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cdar()
	}
}

func BenchmarkCellCdar100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cdar()
	}
}

func BenchmarkCellCdar1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cdar()
	}
}

func BenchmarkCellCddr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cddr()
	}
}

func BenchmarkCellCddr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cddr()
	}
}

func BenchmarkCellCddr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cddr()
	}
}

func BenchmarkCellCddr100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cddr()
	}
}

func BenchmarkCellCddr1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cddr()
	}
}

func BenchmarkCellRplaca0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Rplaca(nil)
	}
}

func BenchmarkCellRplaca1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplaca(nil)
	}
}

func BenchmarkCellRplacd0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Rplacd(nil)
	}
}

func BenchmarkCellRplacd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplacd(nil)
	}
}

func BenchmarkCellOffset0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Offset(0)
	}
}

func BenchmarkCellOffset1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Offset(0)
	}
}

func BenchmarkCellOffset10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Offset(9)
	}
}

func BenchmarkCellOffset100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Offset(99)
	}
}

func BenchmarkCellOffset1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Offset(999)
	}
}

func BenchmarkCellEnd0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.End()
	}
}

func BenchmarkCellEnd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.End()
	}
}

func BenchmarkCellEnd10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.End()
	}
}

func BenchmarkCellEnd100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.End()
	}
}

func BenchmarkCellEnd1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.End()
	}
}

func BenchmarkCellEach0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Each(func(v interface{}) {})
	}
}

func BenchmarkCellEach1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Each(func(v interface{}) {})
	}
}

func BenchmarkCellEach10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Each(func(v interface{}) {})
	}
}

func BenchmarkCellEach100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Each(func(v interface{}) {})
	}
}

func BenchmarkCellStep0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkCellStep1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkCellStep10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkCellStep100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkCellStep1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkCellMap0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkCellMap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkCellMap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkCellMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkCellMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReduce0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReduce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReduce10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReduce100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReduce1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkCellReverse0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Reverse()
	}
}

func BenchmarkCellReverse1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Reverse()
	}
}

func BenchmarkCellReverse10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Reverse()
	}
}

func BenchmarkCellReverse100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Reverse()
	}
}

func BenchmarkCellReverse1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Reverse()
	}
}

func BenchmarkCellCopy0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Copy()
	}
}

func BenchmarkCellCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Copy()
	}
}

func BenchmarkCellCopy10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Copy()
	}
}

func BenchmarkCellCopy100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Copy()
	}
}

func BenchmarkCellCopy1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Copy()
	}
}

func BenchmarkCellRepeat0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(0)
	}
}

func BenchmarkCellRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(1)
	}
}

func BenchmarkCellRepeat10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(10)
	}
}

func BenchmarkCellRepeat100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(100)
	}
}

func BenchmarkCellRepeat1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(1000)
	}
}

func BenchmarkCellAppend0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List0.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkCellAppend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List1.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkCellAppend10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List10.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkCellAppend100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List100.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkCellAppend1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List1000.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkCellWhile0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkCellWhile1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkCellWhile10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkCellWhile100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkCellWhile1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkCellPartition0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Partition(0)
	}
}

func BenchmarkCellPartition1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Partition(0)
	}
}

func BenchmarkCellPartition10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Partition(9)
	}
}

func BenchmarkCellPartition100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Partition(99)
	}
}

func BenchmarkCellPartition1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Partition(999)
	}
}

func BenchmarkCellZip0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Zip(List0)
	}
}

func BenchmarkCellZip1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Zip(List1)
	}
}

func BenchmarkCellZip10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Zip(List10)
	}
}

func BenchmarkCellZip100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Zip(List100)
	}
}

func BenchmarkCellZip1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Zip(List1000)
	}
}