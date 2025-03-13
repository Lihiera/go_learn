package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree {
	var t *tree
	for _, v := range values {
		t = add(t, v)
	}
	return t
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	var visit func(t *tree)
	visit = func(t *tree) {
		if t == nil {
			return
		}
		visit(t.left)
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", t.value)
		visit(t.right)
	}
	visit(t)
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	t := Sort([]int{7, 2, 5, 6, 1, 3, 9, 8, 0, -2, 189})
	fmt.Println(t)
}
