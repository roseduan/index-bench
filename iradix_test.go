package index_bench

import (
	"fmt"
	"github.com/hashicorp/go-immutable-radix/v2"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestIRadix(t *testing.T) {
	tree := iradix.New[[]byte]()

	t.Log("pid = ", os.Getpid())

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		tree, _, _ = tree.Insert(GetTestKey(i), RandomValue(4096))
	}
	t.Log("time cost:", time.Since(now))

	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	// 转换为兆字节
	totalMemory := float64(stat.Alloc) / 1024 / 1024
	heapMemory := float64(stat.HeapAlloc) / 1024 / 1024

	fmt.Printf("总内存使用量：%.2f MB\n", totalMemory)
	fmt.Printf("堆内存使用量：%.2f MB\n", heapMemory)
	//time.Sleep(time.Hour)
}

var tree = iradix.New[[]byte]()

func BenchmarkPut_IRadix(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tree, _, _ = tree.Insert(GetTestKey(i), RandomValue(4096))
	}
}

func BenchmarkGet_IRadix(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		tree, _, _ = tree.Insert(GetTestKey(i), RandomValue(16))
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tree.Get(GetTestKey(rand.Intn(1000000)))
	}
}

func BenchmarkDelete_IRadix(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		tree, _, _ = tree.Insert(GetTestKey(i), RandomValue(16))
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tree, _, _ = tree.Delete(GetTestKey(rand.Intn(1000000)))
	}
}
