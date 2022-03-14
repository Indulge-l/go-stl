package container

type Stack struct {
	arr  []int
	size int
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Pop() {
	s.size--
	s.arr = s.arr[:s.size]
}

func (s *Stack) Top() (val int) {
	return s.arr[s.size-1]
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() (size int) {
	return s.size
}
