package index_bench

import (
	"fmt"
	goart "github.com/plar/go-adaptive-radix-tree"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestArt(t *testing.T) {
	tree := goart.New()

	t.Log("pid = ", os.Getpid())

	now := time.Now()
	for i := 0; i < 1000000; i++ {
		tree.Insert(GetTestKey(i), RandomValue(4096))
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

var art = goart.New()

func BenchmarkPut_ART(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		art.Insert(GetTestKey(i), RandomValue(4096))
	}
}

func BenchmarkGet_ART(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		art.Insert(GetTestKey(i), RandomValue(16))
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		art.Search(GetTestKey(rand.Intn(1000000)))
	}
}

func BenchmarkDelete_ART(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		art.Insert(GetTestKey(i), RandomValue(16))
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		art.Delete(GetTestKey(rand.Intn(1000000)))
	}
}
