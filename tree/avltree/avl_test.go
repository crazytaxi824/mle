package avltree

import (
	"crypto/rand"
	"math/big"
	"runtime"
	"testing"
	"time"
)

func Test_GCRemove(t *testing.T) {
	// 需要开启 Add() 中的 runtime.SetFinalizer
	tree := newTree()
	removes := make([]int64, 0, 10) // to remove
	for i := 0; i < 1000; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		// duplicated index will cause this error
		if err := tree.Insert(b.Int64(), nil); err != nil {
			t.Logf("Add node %d error: %v\n", b.Int64(), err)
		}

		if i%100 == 0 {
			removes = append(removes, b.Int64())
		}
	}

	t.Log("tree size:", tree.Size())
	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}

	for _, v := range removes {
		// index not exist will cause this error
		err := tree.Delete(v)
		if err != nil {
			t.Logf("Add node %d error: %v\n", v, err)
		}
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()

	t.Log("tree size:", tree.Size())
	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}
}

func Test_RemoveRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		tree := newTree()
		removes := make([]int64, 0, 10)
		for i := 0; i < 1000; i++ {
			b, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				t.Error(err)
				return
			}

			tree.Insert(b.Int64(), nil)

			if i%100 == 0 {
				removes = append(removes, b.Int64())
			}
		}

		if er := tree.checkAllNodes(); er != nil {
			t.Error(er)
		}

		for _, v := range removes {
			err := tree.Delete(v)
			if err != nil {
				// not existed index will cause this error
				// NOTE removes 中有非常小的机率会出现相同数字，导致重复删除报错。
				t.Log("remove node", v, "error:", err)
			}
		}

		if er := tree.checkAllNodes(); er != nil {
			t.Error(er)
		}
	}
}

func Test_RemoveAll(t *testing.T) {
	tree := newTree()
	removes := make([]int64, 0, 10)
	for i := 0; i < 1000; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		err = tree.Insert(b.Int64(), nil)
		if err == nil {
			removes = append(removes, b.Int64())
		}
	}

	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}

	for _, v := range removes {
		err := tree.Delete(v)
		if err != nil {
			// not existed index will cause this error
			t.Log("remove node", v, "error:", err)
		}
	}

	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}
}

func Test_ClearTreeGC(t *testing.T) {
	tr := newTree()
	runtime.SetFinalizer(tr, func(tt *tree) {
		t.Log("tree GC")
	})

	traceMemStatsMark("start")

	for i := 0; i < 100; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		index := b.Int64()

		// value 设置为 1M
		err = tr.Insert(index, make([]byte, 1<<20))
		if err != nil {
			t.Error("add node:", index, "error:", err)
		}
	}

	if er := tr.checkAllNodes(); er != nil {
		t.Error(er)
	}

	traceMemStatsMark("full")

	tr.Clear()

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	time.Sleep(200 * time.Millisecond)
	traceMemStatsMark("end")
}

func Test_MemStress(t *testing.T) {
	tree := newTree()
	var removes []int64
	traceMemStatsMark("start")

	for i := 0; i < 100; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		index := b.Int64()

		// value 设置为 1M
		err = tree.Insert(index, make([]byte, 1<<20))
		if err != nil {
			t.Error("add node:", index, "error:", err)
		} else {
			removes = append(removes, b.Int64())
		}
	}

	t.Log("tree size:", tree.Size())
	traceMemStatsMark("full")

	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}

	for i := 0; i < 50; i++ {
		err := tree.Delete(removes[i])
		if err != nil {
			t.Error("remove node", removes[i], "error:", err)
		}
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	time.Sleep(200 * time.Millisecond)
	traceMemStatsMark("half")

	for i := 50; i < 100; i++ {
		err := tree.Delete(removes[i])
		if err != nil {
			t.Error("remove node", removes[i], "error:", err)
		}
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	time.Sleep(200 * time.Millisecond)
	traceMemStatsMark("empty")

	t.Log("tree size:", tree.Size())
	if er := tree.checkAllNodes(); er != nil {
		t.Error(er)
	}
	t.Log(tree.Root() == nil)
}
