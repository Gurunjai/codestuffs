package memch_test

import (
	"encoding/xml"
	"sync"
	"testing"
)

type LReq struct {
	id int    `xml:"id"`
	p  string `xml:"provider"`
	a  string `xml:"assetid"`
	s  string `xml:"subtype"`
	r  uint32 `xml:"rate"`
}

var gid int

var lpool = NewRequestPool(10)

var reqPool = sync.Pool{
	New: func() interface{} {
		return &LReq{}
	},
}

func (l *LReq) constructReq(p, a, s string) {
	l.p = p
	l.a = a
	l.s = s
	l.r = 3750000
	l.id = gid
}

func (l *LReq) reset() {
	l.p = ""
	l.a = ""
	l.s = ""
	l.r = 0
	l.id = 0
}

type LReqPool struct {
	lc chan *LReq
}

func NewRequestPool(size int) *LReqPool {
	lp := &LReqPool{
		lc: make(chan *LReq, size),
	}
	return lp
}

func (lp *LReqPool) Get() (r *LReq) {
	select {
	case r = <-lp.lc:
	default:
		r = &LReq{}
	}

	return
}

func (lp *LReqPool) Put(r *LReq) {
	// Reset Request
	r.reset()
	select {
	case lp.lc <- r:
	default:
	}
}

func BenchmarkLocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var lr LReq
		lr.constructReq("testnobench", "assetnoch", "index")
		out, err := xml.Marshal(lr)
		if err != nil {
			b.Errorf("failed to marshal the data: %v\n", err.Error())
		}

		var lr2 LReq
		if err := xml.Unmarshal(out, &lr2); err != nil {
			b.Errorf("failed to unmarshal the data: %v\n", err)
		}
	}
}
func BenchmarkLocateChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lr := lpool.Get()
		// defer lpool.Put(lr)

		lr.constructReq("testbench", "assetch", "index")
		out, err := xml.Marshal(lr)
		if err != nil {
			b.Errorf("failed to marshal the data: %v\n", err.Error())
		}
		lpool.Put(lr)

		lr2 := lpool.Get()
		if err := xml.Unmarshal(out, lr2); err != nil {
			b.Errorf("failed to unmarshal the data: %v\n", err)
		}
	}
}

func BenchmarkLocatePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lr := reqPool.Get().(*LReq)
		// defer reqPool.Put(lr)

		lr.constructReq("testpool", "assetpool", "index")
		out, err := xml.Marshal(lr)
		if err != nil {
			b.Errorf("failed to marsha the data: %v\n", err.Error())
		}
		reqPool.Put(lr)

		lr2 := reqPool.Get().(*LReq)
		if err := xml.Unmarshal(out, lr2); err != nil {
			b.Errorf("failed to unmarshal the data: %v\n", err)
		}
	}
}
