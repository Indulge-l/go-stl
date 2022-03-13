package sort

func MergeSortInt(arr []int, low, height int) {
	temp := make([]int, len(arr))
	mergeSortInt(arr, temp, low, height)
}

func mergeSortInt(arr, temp []int, low, height int) {
	if low >= height {
		return
	}
	mid := (low + height) / 2
	mergeSortInt(arr, temp, low, mid)
	mergeSortInt(arr, temp, mid+1, height)
	l, r := low, mid+1
	for i := low; i <= height; i++ {
		if r > height || (l <= mid && arr[l] <= arr[r]) {
			temp[i] = arr[l]
			l++
		} else {
			temp[i] = arr[r]
			r++
		}
	}
	for i := low; i <= height; i++ {
		arr[i] = temp[i]
	}
}
