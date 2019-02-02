package traverser

import "testing"

type TestStruct struct {
	A string
	B int
	C map[string]string
	D []int
}

func TestTraverse(t *testing.T) {
	var a = &TestStruct{"hello", 0, map[string]string{"hi": "how are you"}, []int{3, 4, 5, 6, 9}}
	Traverse(a, IsZeroValue)
}
