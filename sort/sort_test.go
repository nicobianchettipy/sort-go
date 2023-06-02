package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingNative(t *testing.T) {
	vendorsSorted := sortByNative(getInputForTest())
	assert.Equal(t, vendorsSorted, getOutputExpectedForTest())
}

func TestBubbleSort(t *testing.T) {
	vendorsSorted := BubbleSort(getInputForTest())
	assert.Equal(t, vendorsSorted, getOutputExpectedForTest())
}

func TestInsertionSort(t *testing.T) {
	vendorsSorted := InsertionSort(getInputForTest())
	assert.Equal(t, vendorsSorted, getOutputExpectedForTest())
}

func TestQuickSort(t *testing.T) {
	vendorsSorted := QuickSort(getInputForTest())
	assert.Equal(t, vendorsSorted, getOutputExpectedForTest())
}

func TestSortingMerge(t *testing.T) {
	vendorsSorted := MergeSort(getInputForTest())
	assert.Equal(t, vendorsSorted, getOutputExpectedForTest())
}

func getInputForTest() VendorList {
	return VendorList{
		Vendor{ID: "1", Distance: 2.3},
		Vendor{ID: "2", Distance: 1.3},
		Vendor{ID: "3", Distance: 2.4},
		Vendor{ID: "4", Distance: 2.7},
		Vendor{ID: "5", Distance: 5.3},
		Vendor{ID: "6", Distance: 7.3},
		Vendor{ID: "7", Distance: 2.5},
	}
}

func getOutputExpectedForTest() VendorList {
	return VendorList{
		Vendor{ID: "2", Distance: 1.3},
		Vendor{ID: "1", Distance: 2.3},
		Vendor{ID: "3", Distance: 2.4},
		Vendor{ID: "7", Distance: 2.5},
		Vendor{ID: "4", Distance: 2.7},
		Vendor{ID: "5", Distance: 5.3},
		Vendor{ID: "6", Distance: 7.3},
	}
}
