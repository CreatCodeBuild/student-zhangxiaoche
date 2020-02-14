package search

import "testing"
import "../search"
import "container/list"

func TestSearch(t *testing.T) {
	a := "D:/MyProject/homework/search/ceshi.txt"
	b := "坑人"
	c := "，"
	shouldBe := "太坑人了！墨瞳之所以决定潜入秘府"
	var real list.List
	real = search.Search(a, b, c);
	r1 :=real.Back().Value
	if  r1 == shouldBe {
		t.Errorf("Search(%s, %s) should be %s, but is:%s\n", a, b, shouldBe, r1)
	}
}
