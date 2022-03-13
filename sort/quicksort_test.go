package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSortInt(t *testing.T) {
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Int()%1000)
	}
	fmt.Printf("befor quick sort: %v\n", nums)
	QuickSortInt(nums, 0, len(nums)-1)
	fmt.Printf("after quick sort: %v\n", nums)
}
