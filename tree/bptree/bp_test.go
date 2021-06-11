package bptree

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"testing"
	"time"
)

var items = []int64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42}

func Test_Add(t *testing.T) {
	tree, _ := newTree(4)
	for _, v := range items {
		tree.Insert(v, nil)
	}

	tree.printTree()
}

func Test_Get(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	// tree.printTree()
	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	t.Log(tree.Search(50))
}

func Test_LessThan(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	// tree.printTree()
	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	fmt.Println("LT")
	for _, v := range tree.SearchLessThan(20, false, 6, 16) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nLQT\n")
	for _, v := range tree.SearchLessThan(20, true, 6, 16) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nLT\n")
	for _, v := range tree.SearchLessThan(20, false, 3, 16) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nLQT\n")
	for _, v := range tree.SearchLessThan(20, true, 3, 16) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()
}

func Test_GreaterThan(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	// tree.printTree()
	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	fmt.Println("GT")
	for _, v := range tree.SearchGreaterThan(90, false, 3, 0) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nGQT\n")
	for _, v := range tree.SearchGreaterThan(90, true, 3, 0) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nGT\n")
	for _, v := range tree.SearchGreaterThan(90, false, 3, 5) {
		fmt.Printf("%d ", v.Key())
	}

	fmt.Printf("\nGQT\n")
	for _, v := range tree.SearchGreaterThan(90, true, 3, 5) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()
}

func Test_FromTo(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	// tree.printTree()
	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	for _, v := range tree.SearchFromTo(20, true, 30, true, 0, 0) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, true, 30, false, 0, 0) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, false, 30, true, 0, 0) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, false, 30, false, 0, 0) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, true, 30, true, 5, 5) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, true, 30, false, 5, 5) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, false, 30, true, 5, 5) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()

	for _, v := range tree.SearchFromTo(20, false, 30, false, 5, 5) {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()
}

func Test_Sort(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	for _, v := range tree.Sort() {
		fmt.Printf("%d ", v.Key())
	}
	fmt.Println()
}

func Test_SmallestAndLargest(t *testing.T) {
	tree, _ := newTree(4)
	for i := 0; i < 100; i++ {
		tree.Insert(int64(i), i)
	}

	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	t.Log(tree.Smallest(), tree.Largest())
}

func Test_Height(t *testing.T) {
	tree, _ := newTree(4)

	for i := 0; i < 1000; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		index := b.Int64()

		tree.Insert(index, nil)
	}

	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	t.Log("height:", tree.Height())
}

func Test_AddRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		tree, _ := newTree(4)

		for i := 0; i < 1000; i++ {
			b, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				t.Error(err)
				return
			}

			index := b.Int64()

			tree.Insert(index, nil)
		}

		if re := tree.checkAllNodes(); re != nil {
			t.Error("add:", re)
		}
	}
}

func Test_GCRemoveAll(t *testing.T) {
	// 需要开启 newLeaf() newInternal() 中的 runtime.SetFinalizer
	tree, _ := newTree(4)
	var removes []int64
	for i := 0; i < 100; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(1000000))
		if err != nil {
			t.Error(err)
			return
		}

		index := b.Int64()

		err = tree.Insert(index, nil)
		if err != nil {
			t.Error("add node:", index, "error:", err)
		} else {
			removes = append(removes, b.Int64())
		}
	}

	tree.printTree()

	t.Log("tree size:", tree.Size())
	t.Log("tree len:", tree.len())
	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	for _, v := range removes {
		err := tree.Delete(v)
		if err != nil {
			t.Error("remove node", v, "error:", err)
		}
	}

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()

	time.Sleep(200 * time.Millisecond)

	t.Log("tree size:", tree.Size())
	t.Log("tree len:", tree.len())
	if re := tree.checkAllNodes(); re != nil {
		t.Error("remove:", re)
	}
	t.Log("root is nil", tree.root == nil)
}

func Test_RemoveRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		tree, _ := newTree(4)

		var removes []int64

		for i := 0; i < 1000; i++ {
			b, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				t.Error(err)
				return
			}

			index := b.Int64()

			tree.Insert(index, nil)

			if i%100 == 0 {
				removes = append(removes, b.Int64())
			}
		}

		if re := tree.checkAllNodes(); re != nil {
			t.Error("add:", re)
		}

		for _, v := range removes {
			// NOTE removes 中有非常小的机率会出现相同数字，导致重复删除报错。
			err := tree.Delete(v)
			if err != nil {
				t.Error("remove node", v, "error:", err)
			}
		}

		if re := tree.checkAllNodes(); re != nil {
			t.Error("remove:", re)
		}
	}
}

func Test_MemStress(t *testing.T) {
	tree, _ := newTree(4)
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
	t.Log("tree len:", tree.len())
	traceMemStatsMark("full")

	if re := tree.checkAllNodes(); re != nil {
		t.Error("add:", re)
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
	t.Log("tree len:", tree.len())
	if re := tree.checkAllNodes(); re != nil {
		t.Error("remove:", re)
	}
}

func Test_Clear(t *testing.T) {
	tr, _ := newTree(4)
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

	t.Log("tree size:", tr.Size())
	t.Log("tree len:", tr.len())
	traceMemStatsMark("full")

	if re := tr.checkAllNodes(); re != nil {
		t.Error("add:", re)
	}

	tr.Clear()

	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()

	time.Sleep(200 * time.Millisecond)
	traceMemStatsMark("end")
}
