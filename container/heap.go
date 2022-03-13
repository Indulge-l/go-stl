package container

type Heap struct {
	arr     []int
	size    int
	isEmpty bool
}

//func NewHeap() *Heap {
//	return new(Heap)
//}

func (h *Heap) Push(val int) {
	h.arr = append(h.arr, val)
	tail := h.size
	h.size++
	for tail >= 0 && h.arr[tail/2] > h.arr[tail] {
		h.arr[tail], h.arr[tail/2] = h.arr[tail/2], h.arr[tail]
		tail = tail / 2
	}
	h.arr[tail] = val
}

func (h *Heap) Empty() (isEmpty bool) {
	return h.isEmpty
}

func (h *Heap) Top() (val int) {
	return h.arr[0]
}

func (h *Heap) Pop() {
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.arr = h.arr[:h.size]
	if h.size == 0 {
		h.isEmpty = true
	}
	h.down()
}

func (h *Heap) down() {
	root := 0
	for root < h.size {
		l, r := root*2+1, root*2+2
		if l < h.size && r < h.size && h.arr[l] > h.arr[r] {
			h.arr[l], h.arr[r] = h.arr[r], h.arr[l]
		}
		if l < h.size && h.arr[l] < h.arr[root] {
			h.arr[l], h.arr[root] = h.arr[root], h.arr[l]
			root = l
		} else {
			break
		}
	}
}
