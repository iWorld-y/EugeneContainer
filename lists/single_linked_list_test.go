package lists

import "testing"

var (
	testArr    = []int{1, 3, 5, 7, 9}
	l          = New(testArr...)
	testArrLen = len(testArr)
)

func TestList_Append(t *testing.T) {
	if l.size != len(testArr) {
		t.Fatal("链表长度错误")
	}
	cnt := 0
	for _, i := range testArr {
		curr := l.Iter()
		if curr == nil {
			break
		}
		if curr.value != i {
			t.Fatal("链表内元素错误")
		}
		cnt++
	}
	if cnt != l.size {
		t.Fatal("未完全遍历链表")
	}
}
func TestList_InsertHead(t *testing.T) {
	l.InsertHead(testArr...)
	if l.size != testArrLen*2 {
		t.Fatal("头插失败")
	}
	cnt := 0
	for _, i := range append(testArr, testArr...) {
		curr := l.Iter()
		if curr == nil {
			break
		}
		if curr.value != i {
			t.Fatal("链表内元素错误")
		}
		cnt++
	}
	if cnt != l.size {
		t.Fatal("未完全遍历链表")
	}
}
