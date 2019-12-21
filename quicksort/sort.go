package quicksort

func Sorts(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	sort(res, 0, len(arr)-1)
	return res
}

func sort(arr []int, a, z int) {
	if z <= a {
		return
	}
	pivot := arr[a+(z-a)/2]
	i, j := a, z
	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	sort(arr, a, j)
	sort(arr, j+1, z)
}
