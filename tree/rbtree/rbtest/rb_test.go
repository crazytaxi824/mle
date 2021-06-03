package rbtest

import (
	"crypto/rand"
	"math/big"
	"runtime"
	"testing"

	"github.com/crazytaxi824/mle/tree/rbtree"
)

var items = []int64{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70, 65}

func Test_Add(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Add(v, nil)
	}

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_Remove(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Add(v, nil)
	}

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}

	tree.Remove(30)
	tree.Remove(15)
	tree.Remove(40)

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_AddRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		tree := rbtree.NewTree()

		var countAddError int
		for i := 0; i < 1000; i++ {
			b, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				t.Error(err)
				return
			}

			index := b.Int64()

			err = tree.Add(index, nil)
			if err != nil {
				countAddError++
			}
		}

		if countAddError+tree.Size() != 1000 {
			t.Error("add count error")
		}

		if re := CheckAllNodes(tree); re != nil {
			t.Error("add:", re)
		}
	}
}

func Test_GCRemove(t *testing.T) {
	// 需要开启 Add() 中的 runtime.SetFinalizer
	tree := rbtree.NewTree()
	removes := make([]int64, 0, 10)
	for i := 0; i < 1000; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		// duplicated index will cause this error
		if err := tree.Add(b.Int64(), nil); err != nil {
			t.Logf("Add node %d error: %v\n", b.Int64(), err)
		}

		if i%100 == 0 {
			removes = append(removes, b.Int64())
		}
	}

	t.Log("tree size:", tree.Size())
	t.Log("tree len:", len(tree.Sort()))

	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}

	for _, v := range removes {
		// index not exist will cause this error
		err := tree.Remove(v)
		if err != nil {
			t.Logf("Add node %d error: %v\n", v, err)
		}
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()

	t.Log("tree size:", tree.Size())
	t.Log("tree len:", len(tree.Sort()))

	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_RemoveRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		tree := rbtree.NewTree()
		removes := make([]int64, 0, 10)
		for i := 0; i < 1000; i++ {
			b, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				t.Error(err)
				return
			}

			tree.Add(b.Int64(), nil)

			if i%100 == 0 {
				removes = append(removes, b.Int64())
			}
		}

		if re := CheckAllNodes(tree); re != nil {
			t.Error(re)
		}

		for _, v := range removes {
			err := tree.Remove(v)
			if err != nil {
				// not existed index will cause this error
				t.Log("remove node", v, "error:", err)
			}
		}

		if re := CheckAllNodes(tree); re != nil {
			t.Error(re)
			return
		}
	}
}

func Test_Sort(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Add(v, nil)
	}

	if err := CheckAllNodes(tree); err != nil {
		t.Error(err)
	}

	var ns []int64
	for _, v := range tree.Sort() {
		ns = append(ns, v.Index())
	}
	t.Log(ns)
}
