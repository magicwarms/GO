package benchmark

import (
	"belajar-unit-test/helper"
	"testing"
)

// cara running benchmark
// go test -v -bench=. --> running semua benchmark file dan unit test
// go test -v -run=NotMathUnitTest -bench=. --> running semua benchmark tanpa unit test
// go test -v -run=NotMathUnitTest -bench=NamaFunctionBenchmark --> running benchmark tertentu tanpa unit test

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloWorld("Andhana")
	}
}

func BenchmarkHelloWorldYuna(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloWorld("Yuna")
	}
}

// sub test benchmark ini memungkinkan kita membuat function benchmark di dalam benchmark
// untuk menjalankan sub test benchmark ini, bisa menggunakan command seperti dibawah ini:
// go test -v -bench=NamaFunctionBenchmark/NamaSub
// example: go test -v -bench=BenchmarkSub/Mahrani
func BenchmarkSub(b *testing.B) {
	b.Run("Mahrani", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.HelloWorld("Mahrani")
		}
	})
	b.Run("Yuna", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helper.HelloWorld("Yuna")
		}
	})
}

// cobain konsep table benchmark
// konsep ini memudahkan kita untuk mengetest benchmark beberapa functions dengan cara iterasi
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarksTable := []struct {
		name string
	}{
		{
			name: "Bulan",
		},
		{
			name: "Matahari",
		},
		{
			name: "Dunias",
		},
	}

	for _, benchmark := range benchmarksTable {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helper.HelloWorld(benchmark.name)
			}
		})
	}
}
