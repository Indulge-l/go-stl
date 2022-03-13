package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Int()%1000)
	}
	fmt.Printf("befor merge sort: %v\n", nums)
	MergeSortInt(nums, 0, len(nums)-1)
	fmt.Printf("after merge sort: %v\n", nums)
}
