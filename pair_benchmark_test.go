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

func BenchmarkPairAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ListInterface.(*Pair)
	}
}

func BenchmarkPairString0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.String()
	}
}

func BenchmarkPairString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.String()
	}
}

func BenchmarkPairString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.String()
	}
}

func BenchmarkPairString100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.String()
	}
}

func BenchmarkPairString1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.String()
	}
}
func BenchmarkPairLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Len()
	}
}

func BenchmarkPairLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Len()
	}
}

func BenchmarkPairLen10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Len()
	}
}

func BenchmarkPairLen100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Len()
	}
}

func BenchmarkPairLen1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Len()
	}
}

func BenchmarkPairIsNil0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.IsNil()
	}
}

func BenchmarkPairIsNil1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.IsNil()
	}
}

func BenchmarkPairEqual0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Equal(List0)
	}
}

func BenchmarkPairEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Equal(List1)
	}
}

func BenchmarkPairEqual10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Equal(List10)
	}
}

func BenchmarkPairEqual100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Equal(List100)
	}
}

func BenchmarkPairEqual1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Equal(List1000)
	}
}

func BenchmarkPairPush0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Push(nil)
	}
}

func BenchmarkPairPush1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Push(nil)
	}
}

func BenchmarkPairPop0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Pop()
	}
}

func BenchmarkPairPop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Pop()
	}
}

func BenchmarkPairIntPair(b *testing.B) {
	b.StopTimer()
		p := Cons(0, 0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.IntPair()
	}
}

func BenchmarkPairPairPair(b *testing.B) {
	b.StopTimer()
		p := Cons(List(), List())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.PairPair()
	}
}

func BenchmarkPairNext0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Next()
	}
}

func BenchmarkPairNext1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Next()
	}
}

func BenchmarkPairCar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Car()
	}
}

func BenchmarkPairCar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Car()
	}
}

func BenchmarkPairCdr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cdr()
	}
}

func BenchmarkPairCdr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cdr()
	}
}

func BenchmarkPairCaar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Caar()
	}
}

func BenchmarkPairCaar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Caar()
	}
}

func BenchmarkPairCaar10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Caar()
	}
}

func BenchmarkPairCaar100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Caar()
	}
}

func BenchmarkPairCaar1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Caar()
	}
}

func BenchmarkPairCadr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cadr()
	}
}

func BenchmarkPairCadr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cadr()
	}
}

func BenchmarkPairCadr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cadr()
	}
}

func BenchmarkPairCadr100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cadr()
	}
}

func BenchmarkPairCadr1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cadr()
	}
}

func BenchmarkPairCdar0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cdar()
	}
}

func BenchmarkPairCdar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cdar()
	}
}

func BenchmarkPairCdar10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cdar()
	}
}

func BenchmarkPairCdar100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cdar()
	}
}

func BenchmarkPairCdar1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cdar()
	}
}

func BenchmarkPairCddr0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Cddr()
	}
}

func BenchmarkPairCddr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cddr()
	}
}

func BenchmarkPairCddr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Cddr()
	}
}

func BenchmarkPairCddr100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Cddr()
	}
}

func BenchmarkPairCddr1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Cddr()
	}
}

func BenchmarkPairRplaca0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Rplaca(nil)
	}
}

func BenchmarkPairRplaca1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplaca(nil)
	}
}

func BenchmarkPairRplacd0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Rplacd(nil)
	}
}

func BenchmarkPairRplacd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplacd(nil)
	}
}

func BenchmarkPairOffset0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Offset(0)
	}
}

func BenchmarkPairOffset1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Offset(0)
	}
}

func BenchmarkPairOffset10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Offset(9)
	}
}

func BenchmarkPairOffset100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Offset(99)
	}
}

func BenchmarkPairOffset1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Offset(999)
	}
}

func BenchmarkPairEnd0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.End()
	}
}

func BenchmarkPairEnd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.End()
	}
}

func BenchmarkPairEnd10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.End()
	}
}

func BenchmarkPairEnd100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.End()
	}
}

func BenchmarkPairEnd1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.End()
	}
}

func BenchmarkPairEach0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Each(func(v interface{}) {})
	}
}

func BenchmarkPairEach1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Each(func(v interface{}) {})
	}
}

func BenchmarkPairEach10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Each(func(v interface{}) {})
	}
}

func BenchmarkPairEach100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Each(func(v interface{}) {})
	}
}

func BenchmarkPairStep0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkPairStep1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkPairStep10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkPairStep100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkPairStep1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Step(0, 1, func(v interface{}) {})
	}
}

func BenchmarkPairMap0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkPairMap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkPairMap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkPairMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkPairMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Map(func(interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReduce0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReduce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReduce10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReduce100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReduce1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Reduce(nil, func(interface{}, interface{}) interface{} { return nil })
	}
}

func BenchmarkPairReverse0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Reverse()
	}
}

func BenchmarkPairReverse1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Reverse()
	}
}

func BenchmarkPairReverse10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Reverse()
	}
}

func BenchmarkPairReverse100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Reverse()
	}
}

func BenchmarkPairReverse1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Reverse()
	}
}

func BenchmarkPairCopy0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Copy()
	}
}

func BenchmarkPairCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Copy()
	}
}

func BenchmarkPairCopy10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Copy()
	}
}

func BenchmarkPairCopy100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Copy()
	}
}

func BenchmarkPairCopy1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Copy()
	}
}

func BenchmarkPairRepeat0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(0)
	}
}

func BenchmarkPairRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(1)
	}
}

func BenchmarkPairRepeat10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(10)
	}
}

func BenchmarkPairRepeat100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(100)
	}
}

func BenchmarkPairRepeat1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Repeat(1000)
	}
}

func BenchmarkPairAppend0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List0.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkPairAppend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List1.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkPairAppend10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List10.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkPairAppend100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List100.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkPairAppend1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
			l := List1000.Copy()
		b.StartTimer()
		l.Append(1)
	}
}

func BenchmarkPairWhile0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkPairWhile1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkPairWhile10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkPairWhile100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkPairWhile1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.While(true, func(i interface{}) bool {
			return true
		})
	}
}

func BenchmarkPairPartition0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Partition(0)
	}
}

func BenchmarkPairPartition1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Partition(0)
	}
}

func BenchmarkPairPartition10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Partition(9)
	}
}

func BenchmarkPairPartition100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Partition(99)
	}
}

func BenchmarkPairPartition1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Partition(999)
	}
}

func BenchmarkPairZip0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List0.Zip(List0)
	}
}

func BenchmarkPairZip1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Zip(List1)
	}
}

func BenchmarkPairZip10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List10.Zip(List10)
	}
}

func BenchmarkPairZip100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List100.Zip(List100)
	}
}

func BenchmarkPairZip1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1000.Zip(List1000)
	}
}