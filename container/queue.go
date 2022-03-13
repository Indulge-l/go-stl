package container

type Queue struct {
	arr     []int
	isEmpty bool
	size    int
}

func (q *Queue) Push(val int) {
	q.arr = append(q.arr, val)
	q.size++
	q.isEmpty = false
}

func (q *Queue) Front() (val int) {
	return q.arr[0]
}

func (q *Queue) Pop() {
	q.size--
	q.arr = q.arr[1:]
	if q.size == 0 {
		q.isEmpty = true
	}
}

func (q *Queue) Empty() (isEmpty bool) {
	return q.isEmpty
}
