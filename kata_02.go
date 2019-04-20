package kata

func binarySearch(arr []int, key int, start int, end int) int {
	if start > end {
		return -1
	}
	middle := start + (end-start)/2
	val := arr[middle]

	if val < key {
		return binarySearch(arr, key, middle+1, end)
	} else if val > key {
		return binarySearch(arr, key, start, end-1)
	} else {
		return middle
	}
}

func chop(key int, arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return binarySearch(arr, key, 0, len(arr)-1)
}
