package gocodecheck

func binarySearch(arr []int, key int, start int, end int) int {
	if len(arr) == 0 || start > end {
		return -1
	}
	middle := start + (end-start)/2
	val := arr[middle]

	if val < key {
		return binarySearch(arr, key, middle+1, end)
	} else if val > key {
		return binarySearch(arr, key, start, middle-1)
	} else {
		return middle
	}
}

// added: keep track of the sliced index
func binarySearchSlice(arr []int, key int, added int) int {
	if len(arr) == 0 {
		return -1
	}

	middle := len(arr) / 2
	val := arr[middle]

	if val < key {
		return binarySearchSlice(arr[middle+1:], key, added+middle+1)
	} else if val > key {
		return binarySearchSlice(arr[:middle], key, added)
	} else {
		return middle + added
	}
}

func binarySearchIterative(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1

	for start <= end {
		middle := start + (end-start)/2
		val := arr[middle]

		if val < key {
			start = middle + 1
		} else if val > key {
			end = middle - 1
		} else {
			return middle
		}
	}

	return -1
}

func chopIterative(key int, arr []int) int {
	return binarySearchIterative(arr, key)
}

func chopSlice(key int, arr []int) int {
	return binarySearchSlice(arr, key, 0)
}

func chop(key int, arr []int) int {
	return binarySearch(arr, key, 0, len(arr)-1)
}
