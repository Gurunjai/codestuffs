package main

import (
	// "fmt"
	"sync"
	"testing"
)


type ReceiveBuff [4096]byte

func (r *ReceiveBuff) Dump() bool {
	return true
}

var rPool = sync.Pool{
	New: func() interface{} {
		return new(ReceiveBuff)
	},
}

func BenchmarkBytePool( b *testing.B ) {
	for i := 0; i < b.N; i++ {
		var buf ReceiveBuff
		buf.Dump()
	}
}

func BenchmarkByRPool( b *testing.B ) {
	for i := 0; i < b.N; i++ {
		buf := rPool.Get().(*ReceiveBuff)
		defer rPool.Put(buf)
		buf.Dump()
	}
}