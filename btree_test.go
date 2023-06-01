package index_bench

import (
	"bytes"
	"fmt"
	"github.com/google/btree"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

type item struct {
	key   []byte
	value []byte
}

func (it *item) Less(bi btree.Item) bool {
	return bytes.Compare(it.key, bi.(*item).key) < 0
}

func TestBTree(t *testing.T) {
	tree := btree.New(32)

	t.Log("pid = ", os.Getpid())

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		tree.ReplaceOrInsert(&item{key: GetTestKey(i), value: RandomValue(4096)})
	}

	t.Log("time cost:", time.Since(now))
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	// 转换为兆字节
	totalMemory := float64(stat.Alloc) / 1024 / 1024
	heapMemory := float64(stat.HeapAlloc) / 1024 / 1024

	fmt.Printf("总内存使用量：%.2f MB\n", totalMemory)
	fmt.Printf("堆内存使用量：%.2f MB\n", heapMemory)
}

var btreeInstance = btree.New(32)

func BenchmarkPut_BTree(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		btreeInstance.ReplaceOrInsert(&item{key: GetTestKey(i), value: RandomValue(4096)})
	}
}

func BenchmarkGet_BTree(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		btreeInstance.ReplaceOrInsert(&item{key: GetTestKey(i), value: RandomValue(16)})
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		btreeInstance.Get(&item{key: GetTestKey(rand.Intn(1000000))})
	}
}

func BenchmarkDelete_BTree(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		btreeInstance.ReplaceOrInsert(&item{key: GetTestKey(i), value: RandomValue(16)})
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		btreeInstance.Delete(&item{key: GetTestKey(rand.Intn(1000000))})
	}
}
