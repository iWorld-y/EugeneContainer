package lists

import "testing"

func TestList_Add(t *testing.T) {
	l := New(1, 3, 5, 7, 9)
	if l.size != 5 {
		t.Fatal("New 单链表失败")
	}
	cnt := 0
	for i := 1; i < 10 && l.first != nil; i += 2 {
		if l.first.value != i {
			t.Fatal("New 单链表失败")
		}
		l.first = l.first.next
		cnt++
	}
	if cnt != l.size {
		t.Fatal("New 单链表失败")
	}
}
