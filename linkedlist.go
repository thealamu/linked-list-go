package linkedlists

type node struct {
	data interface{}
	next *node
}

type LinkedLists struct {
	head *node
}

func New() *LinkedLists {
	return &LinkedLists{}
}

// PushFront adds an item to the front of the list
func (l *LinkedLists) PushFront(item interface{}) {
	if l.head == nil {
		l.head = &node{data: item}
		return
	}

	newHead := &node{data: item}
	newHead.next = l.head
	l.head = newHead
}

// Size returns the length of the linked lists
func (l *LinkedLists) Size() int {
	// walk the list to report the size
	curr := l.head
	i := 0
	for curr != nil {
		i++
		curr = curr.next
	}
	return i
}
