// @Title  
// @Description  
// @Author  fibbery  2021/5/1 上午9:25
package bloomfilter

import (
	"math/rand"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestNewBloomfilter(t *testing.T) {
	filter := NewBloomfilter(1<<31, 4)
	filter.Add("April")
	if !filter.IsExist("April") {
		t.Fatalf("April must in filter")
	}
}

func BenchmarkBloomfilter_Add(b *testing.B) {
	filter := NewBloomfilter(1<<32, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Add("Hello World")
	}
}

func BenchmarkBloomfilter_IsExist(b *testing.B) {
filter := NewBloomfilter(1<<32, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.IsExist("Hello World")
	}
}
