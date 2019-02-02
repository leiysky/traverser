package traverser

import "testing"

type testStruct struct {
	a string
	b int
	c map[string]string
	d []int
}

func TestTraverse(t *testing.T) {
	var a = &testStruct{"hello", 123, map[string]string{"hi": "how are you"}, []int{3, 4, 5, 6, 9}}
	Traverse(a)
}
