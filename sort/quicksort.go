package sort

func QuickSortInt(arr []int, low, height int) {
	if low >= height {
		return
	}
	l, r := low, height
	pos := arr[(low+height)/2]
	for l <= r {
		for arr[l] < pos {
			l++
		}
		for arr[r] > pos {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	QuickSortInt(arr, low, r)
	QuickSortInt(arr, l, height)
}
