package container

type Container[T any] interface {
	Size() int
	Empty() bool
	Clear()
	Values() []T
	String() string
}
