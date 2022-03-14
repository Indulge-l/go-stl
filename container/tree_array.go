package container

import "go-stl/util"

type SumTreeArray struct {
	arr  []int
	size int
}

func NewSumTreeArray(n int) (str *SumTreeArray) {
	return &SumTreeArray{arr: make([]int, n+1), size: n + 1}
}

func lowBit(x int) (val int) {
	return x & (-x)
}

func (str *SumTreeArray) Update(index, val int) {
	for index <= str.size {
		str.arr[index] += val
		index += lowBit(index)
	}
}

func (str *SumTreeArray) Sum(index int) (sum int) {
	for index > 0 {
		sum += str.arr[index]
		index -= lowBit(index)
	}
	return sum
}

func (str *SumTreeArray) RangeSum(left, right int) (sum int) {
	return str.Sum(right) - str.Sum(left-1)
}

type MaxTreeArray struct {
	arr  []int
	size int
}

func NewMaxTreeArray(n int) (mtr *MaxTreeArray) {
	return &MaxTreeArray{arr: make([]int, n+1), size: n + 1}
}

func (mtr *MaxTreeArray) Update(index, val int) {
	for index <= mtr.size {
		mtr.arr[index] = util.Max(mtr.arr[index], val)
		index += lowBit(index)
	}
}

func (mtr *MaxTreeArray) Max(index int) (maxNum int) {
	for index > 0 {
		maxNum = util.Max(mtr.arr[index], maxNum)
		index -= lowBit(index)
	}
	return maxNum
}
