package container

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	lens = 10
)

func TestSumTreeArray(t *testing.T) {
	var nums []int
	tr := NewSumTreeArray(lens)
	for i := 0; i < lens; i++ {
		nums = append(nums, rand.Int()%10+1)
		tr.Update(i+1, nums[i])
	}
	fmt.Println(nums)
	fmt.Println(tr.Sum(5))
	fmt.Println(tr.RangeSum(5, 5))
}

func TestMaxTreeArray(t *testing.T) {
	var nums []int
	tr := NewMaxTreeArray(lens)
	for i := 0; i < lens; i++ {
		nums = append(nums, rand.Int()%10+1)
		tr.Update(i+1, nums[i])
	}
	fmt.Println(nums)
	fmt.Println(tr.Max(5))
	tr.Update(5, -1)
	fmt.Println(tr.Max(6))
}
