package lists

type List[T any] struct {
	first, last *node[T]
	size        int
}
type node[T any] struct {
	value T
	next  *node[T]
}

func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (l *List[T]) Add(values ...T) {
	for _, value := range values {
		currNode := &node[T]{
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
