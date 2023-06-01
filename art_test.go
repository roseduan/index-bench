package index_bench

import (
	goart "github.com/plar/go-adaptive-radix-tree"
	"math/rand"
	"os"
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
	time.Sleep(time.Hour)
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
