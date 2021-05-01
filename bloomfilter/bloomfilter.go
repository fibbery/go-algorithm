// @Title  
// @Description  布隆过滤器实现
// @Author  fibbery  2021/4/30 下午2:51
package bloomfilter

import (
	"crypto/md5"
	"math"
)

type Bloomfilter struct {
	n uint64 // nums of insert element
	k int    // nums of hash function

	m    uint64 // length of data
	data []byte // bits of data
}

// NewBloomfilter return a new bloomfilter
func NewBloomfilter(m uint64, k int) *Bloomfilter {
	// use &(2^n - 1) instead of %
	log2 := math.Ceil(math.Log2(float64(m)))
	return &Bloomfilter{
		n:    0,
		k:    k,
		m:    1 << uint64(log2),
		data: make([]byte, m),
	}
}

func (b *Bloomfilter) Add(data string) {
	slots := b.slot([]byte(data))
	for i := 0; i < len(slots); i++ {
		index := slots[i] & (b.m - 1)
		b.data[index] |= 1
	}
	b.n++
}

func (b *Bloomfilter) IsExist(data string) bool {
	slots := b.slot([]byte(data))
	for i := 0; i < len(slots); i++ {
		index := slots[i] & (b.m - 1)
		if b.data[index] != 1 {
			return false
		}
	}
	return true
}

func (b *Bloomfilter) slot(data []byte) []uint64 {
	ret := make([]uint64, 0)
	hash := md5.New()
	for i := 0; i < len(ret); i++ {
		hash.Reset()
		hash.Write(data)
		hash.Write([]byte{byte(i)})
		key := hash.Sum(nil)
		cnt := uint64(0)
		for _, v := range key {
			cnt += uint64(v)
		}
		ret = append(ret, cnt)
	}
	return ret
}

// GetFalseRate return False Positive Rate
// (1−e^(−kn/m))^k
func (b *Bloomfilter) GetFalseRate() float64 {
	expoInner := -(float64)(uint64(b.k)*b.n) / float64(b.m)
	rate := math.Pow(1-math.Pow(math.E, expoInner), float64(b.k))
	return rate
}
