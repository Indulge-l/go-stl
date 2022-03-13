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
	h := new(ListNode)
	t := new(ListNode)
	h.next = t
	t.pre = h
	return &List{head: h, tail: t}
}

func (l *List) AddHead(val int) {
	node := new(ListNode)
	node.val = val
	node.next = l.head
	l.head.pre = node
	l.head = node
	l.size++
}

func (l *List) GetHead() (val int) {
	return l.head.next.val
}

func (l *List) AddTail(val int) {
	node := new(ListNode)
	node.val = val
	node.pre = l.tail
	l.tail.next = node
	l.head = node
	l.size++
}

func (l *List) GetTail() (val int) {
	return l.tail.pre.val
}

func (l *List) AddNodeByNo(index, val int) {
	temp := l.head.next
	pos := 1
	for temp != l.tail {
		if pos == index {
			node := &ListNode{val: val}
			l.addNewNode(temp.pre, node, temp)
		}
	}
}

func (l *List) DeleteFirstByVal(val int) {
	temp := l.head.next
	for temp != l.tail {
		if temp.val == val {
			l.deleteNode(temp)
			break
		}
	}
}

func (l *List) DeleteAllByVal(val int) {
	temp := l.head.next
	for temp != l.tail {
		if temp.val == val {
			l.deleteNode(temp)
		}
	}
}

// DeleteByNo 删除指定位置节点
func (l *List) DeleteByNo(index int) {
	temp := l.head.next
	pos := 1
	for temp != l.tail {
		if pos == index {
			l.deleteNode(temp)
			break
		}
		pos++
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
