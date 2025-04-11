package sr

// QuickSort sorts a slice of integers using the quicksort algorithm
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

// quickSort is the recursive implementation of quicksort
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivot := partition(arr, low, high)

		// Recursively sort elements before and after the pivot
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

// partition rearranges the elements and returns the pivot index
func partition(arr []int, low, high int) int {
	// Choose the rightmost element as pivot
	pivot := arr[high]

	// Index of smaller element
	i := low - 1

	for j := low; j < high; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap arr[i+1] and arr[high] (pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
