package sort

import (
	"math/rand"
	"testing"
	"time"
)

/*
   Execute command go test -bench=. -benchmem
*/

func BenchmarkSortByNative(b *testing.B) {
	vendorList := generatorInputForBenchmark()
	// Realiza la operación b.N veces
	for n := 0; n < b.N; n++ {
		_ = sortByNative(vendorList)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	vendorList := generatorInputForBenchmark()
	// Realiza la operación b.N veces
	for n := 0; n < b.N; n++ {
		_ = BubbleSort(vendorList)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	vendorList := generatorInputForBenchmark()
	// Realiza la operación b.N veces
	for n := 0; n < b.N; n++ {
		_ = InsertionSort(vendorList)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	vendorList := generatorInputForBenchmark()
	// Realiza la operación b.N veces
	for n := 0; n < b.N; n++ {
		_ = QuickSort(vendorList)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	vendorList := generatorInputForBenchmark()
	// Realiza la operación b.N veces
	for n := 0; n < b.N; n++ {
		_ = MergeSort(vendorList)
	}
}

func generatorInputForBenchmark() VendorList {
	vendorList := VendorList{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 30000; i++ {
		vendorList = append(vendorList, Vendor{ID: "11111111", Distance: rand.Float32() * 9000})
	}
	return vendorList
}
