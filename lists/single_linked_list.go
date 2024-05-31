package lists

type List[T any] struct {
	first, last *Node[T]
	size        int
	curr        *Node[T]
}
type Node[T any] struct {
	value T
	next  *Node[T]
}

func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

func (l *List[T]) Append(values ...T) {
	for _, value := range values {
		currNode := &Node[T]{
			value: value,
		}
		if l.size == 0 {
			l.first = currNode
			l.last = currNode
		} else {
			l.last.next = currNode
			l.last = currNode
		}
		l.size++
	}
}

func (l *List[T]) InsertHead(values ...T) {
	if l.size == 0 {
		l.Append(values...)
		return
	}
	tmpL := New(values...)
	tmpL.last.next = l.first
	l.first = tmpL.first
	l.size += tmpL.size
}

func (l *List[T]) Iter() *Node[T] {
	if l.curr == nil {
		if l.first == nil {
			return nil
		}
		l.curr = l.first
		return l.curr
	}
	if l.curr.next == nil {
		return nil
	}
	l.curr = l.curr.next
	return l.curr
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.Size() == 0
}

func (l *List[T]) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}
