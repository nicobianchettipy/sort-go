package sort

import "sort"

func sortByNative(vendorsIn VendorList) VendorList {
	sort.Sort(VendorList(vendorsIn))
	return vendorsIn
}

// BubbleSort
func BubbleSort(items VendorList) VendorList {
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if items[j].Distance > items[j+1].Distance {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

// InsertionSort
func InsertionSort(items VendorList) VendorList {
	n := len(items)
	for i := 1; i < n; i++ {
		key := items[i]
		j := i - 1
		for j >= 0 && items[j].Distance > key.Distance {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = key
	}
	return items
}

func QuickSort(items VendorList) VendorList {
	quickSort(items, 0, len(items)-1)
	return items
}

func quickSort(items VendorList, low, high int) {
	if low < high {
		pivotIndex := partition(items, low, high)
		quickSort(items, low, pivotIndex-1)
		quickSort(items, pivotIndex+1, high)
	}
}

func partition(items VendorList, low, high int) int {
	pivot := items[high]
	i := low - 1
	for j := low; j < high; j++ {
		if items[j].Distance < pivot.Distance {
			i++
			items[i], items[j] = items[j], items[i]
		}
	}
	items[i+1], items[high] = items[high], items[i+1]
	return i + 1
}

func MergeSort(items VendorList) VendorList {
	if len(items) <= 1 {
		return items
	}

	mid := len(items) / 2
	MergeSort(items[:mid])
	MergeSort(items[mid:])

	merge(items, mid)
	return items
}

// merge es una función auxiliar para combinar las dos mitades ordenadas.
func merge(items VendorList, mid int) {
	leftSize := mid
	rightSize := len(items) - mid

	left := make([]Vendor, leftSize)
	right := make([]Vendor, rightSize)

	copy(left, items[:mid])
	copy(right, items[mid:])

	i := 0
	j := 0
	k := 0

	for i < leftSize && j < rightSize {
		if left[i].Distance <= right[j].Distance {
			items[k] = left[i]
			i++
		} else {
			items[k] = right[j]
			j++
		}
		k++
	}

	for i < leftSize {
		items[k] = left[i]
		i++
		k++
	}

	for j < rightSize {
		items[k] = right[j]
		j++
		k++
	}
}
