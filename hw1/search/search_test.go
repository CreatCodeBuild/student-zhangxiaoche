package search

import (
	"container/list"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	a := "./ceshi.txt"
	b := "坑人"
	c := "，"
	shouldBe := "太坑人了！墨瞳之所以决定潜入秘府，"
	var real list.List

	file, _ := os.Open(a)
	real = Search(file, b, c)
	r1 := real.Back().Value
	if !reflect.DeepEqual(r1, shouldBe) {
		t.Errorf("Search(%s, %s) should be %s, but is:%s\n", a, b, shouldBe, r1)
	}
}

func TestSearch2(t *testing.T) {
	a := "./ceshi.txt"
	b := "坑人"
	c := "，"
	var real list.List

	real = Search(ioutil.NopCloser(strings.NewReader(a)), b, c)
	r1 := real.Back()
	if r1 != nil {
		t.Error("Should not find result")
	}
}

func BenchmarkSearch(b *testing.B) {
	a1 := "./ceshi.txt"
	b1 := "坑人"
	c1 := "，"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, _ := os.Open(a1)
		Search(file, b1, c1)
	}
}
