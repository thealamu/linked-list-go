package linkedlists

type LinkedLists struct {
	head *node
}

type node struct {
	data interface{}
	next *node
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
	i := 0
	l.Walk(func(interface{}) { i++ })
	return i
}

// Returns true if list is empty
func (l *LinkedLists) Empty() bool {
	return l.Size() == 0
}

// Function to call for each iteration when walking
type iterFunc func(interface{})

// Walk walks the linked lists, calling the iterFunc each iteration
func (l *LinkedLists) Walk(f iterFunc) {
	curr := l.head
	for curr != nil {
		f(curr.data)
		curr = curr.next
	}
}
