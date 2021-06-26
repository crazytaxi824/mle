package example

import (
	"testing"

	"github.com/crazytaxi824/mle/tree/bptree"
)

var items = []int64{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

func Test_Insert(t *testing.T) {
	tree, _ := bptree.NewTree(4)
	for _, v := range items {
		tree.Insert(v, v*100)
	}

	t.Log(tree.Search(50))

	kvs := tree.SearchGreaterThan(90, true, 0, 0)
	t.Log("greater than")
	for _, v := range kvs {
		t.Log(v.Key(), v.Value())
	}

	kvs = tree.SearchLessThan(40, true, 0, 0)
	t.Log("less than")
	for _, v := range kvs {
		t.Log(v.Key(), v.Value())
	}

	tree.Delete(70)
}
