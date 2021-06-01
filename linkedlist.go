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
