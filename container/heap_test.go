package container

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMinHeap(t *testing.T) {
	const lens = 10
	var nums []int
	for i := 0; i < lens; i++ {
		nums = append(nums, rand.Int()%1000)
	}
	fmt.Printf("arr: %v\n", nums)
	h := new(Heap)
	for i := 0; i < lens; i++ {
		h.Push(nums[i])
	}
	for !h.Empty() {
		fmt.Println(h.Top())
		h.Pop()
	}
}
