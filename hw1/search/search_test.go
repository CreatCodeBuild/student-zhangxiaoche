package search

import (
	"container/list"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	a := "./ceshi.txt"
	b := "坑人"
	c := "，"
	shouldBe := "太坑人了！墨瞳之所以决定潜入秘府"
	var real list.List

	file, _ := os.Open(a)
	real = Search(file, b, c)
	r1 := real.Back().Value
	if r1 == shouldBe {
		t.Errorf("Search(%s, %s) should be %s, but is:%s\n", a, b, shouldBe, r1)
	}
}
func BenchmarkSearch(b *testing.B) {
	a1 := "./ceshi.txt"
	b1 := "坑人"
	c1 := "，"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(a1, b1, c1)
	}
}
