package index_bench

import (
	"github.com/hashicorp/go-immutable-radix/v2"
	"math/rand"
	"os"
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
	time.Sleep(time.Hour)
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
