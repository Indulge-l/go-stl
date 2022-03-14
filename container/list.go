package container

type List struct {
	head *ListNode
	tail *ListNode
	size int
}

type ListNode struct {
	val  int
	pre  *ListNode
	next *ListNode
}

func NewList() (list *List) {
	head := new(ListNode)
	tail := new(ListNode)
	head.next = tail
	tail.pre = head
	return &List{head: head, tail: tail}
}

func (l *List) AddHead(val int) {
	node := new(ListNode)
	node.val = val
	l.addNewNode(l.head, node, l.head.next)
}

func (l *List) GetHead() (val int) {
	return l.head.next.val
}

func (l *List) AddTail(val int) {
	node := new(ListNode)
	node.val = val
	l.addNewNode(l.tail.pre, node, l.tail)
}

func (l *List) GetTail() (val int) {
	return l.tail.pre.val
}

func (l *List) GetNodeByNo(index int) (val int) {
	for temp, pos := l.head.next, 1; temp != l.tail; temp, pos = temp.next, pos+1 {
		if pos == index {
			val = temp.val
			break
		}
	}
	return val
}

func (l *List) AddNodeByNo(index, val int) {
	pos := 1
	node := &ListNode{val: val}
	for temp := l.head.next; temp != l.tail; temp = temp.next {
		if pos == index {
			l.addNewNode(temp.pre, node, temp)
			return
		}
		pos++
	}
	l.addNewNode(l.tail.pre, node, l.tail)
}

func (l *List) DeleteFirstByVal(val int) {
	for temp := l.head.next; temp != l.tail; temp = temp.next {
		if temp.val == val {
			l.deleteNode(temp)
			break
		}
	}
}

func (l *List) DeleteAllByVal(val int) {
	for temp := l.head.next; temp != l.tail; temp = temp.next {
		if temp.val == val {
			l.deleteNode(temp)
		}
	}
}

// DeleteByNo 删除指定位置节点
func (l *List) DeleteByNo(index int) {
	for temp, pos := l.head.next, 1; temp != l.tail; temp, pos = temp.next, pos+1 {
		if pos == index {
			l.deleteNode(temp)
			break
		}
	}
}

func (l *List) Size() (size int) {
	return l.size
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) deleteNode(node *ListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
}

func (l *List) addNewNode(pre, node, next *ListNode) {
	pre.next = node
	node.pre = pre
	node.next = next
	next.pre = node
	l.size++
}
