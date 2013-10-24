package greenspun

import "testing"

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

func BenchmarkCellequal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.equal(List1)
	}
}

func BenchmarkCellEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Equal(List1)
	}
}

func BenchmarkCellCar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Car()
	}
}

func BenchmarkCellCdr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Cdr()
	}
}

func BenchmarkCellRplaca(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplaca(nil)
	}
}

func BenchmarkCellRplacd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List1.Rplacd(nil)
	}
}