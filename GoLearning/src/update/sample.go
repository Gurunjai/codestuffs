package main

import (
	"fmt"
	"math"
	"math/rand"

	// _ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

const MAX_SAMPLES = 4

type mapper struct {
	key   []string
	value float64
}

func (m *mapper) reset() {
	m.key = nil
	m.value = 0
}

func (m *mapper) set(k []string, v float64) {
	m.key = k
	m.value = v
}

type MapperPool struct {
	mCh chan *mapper
}

func NewMapperPool(size int) *MapperPool {
	mp := &MapperPool{
		mCh: make(chan *mapper, size),
	}

	return mp
}

func (mp *MapperPool) Get() (r *mapper) {
	select {
	case r = <-mp.mCh:
	default:
		r = &mapper{}
	}

	return
}

func (mp *MapperPool) Put(r *mapper) {
	r.reset()
	select {
	case mp.mCh <- r:
	default:
	}
}

type Mapify []*mapper

var mapPool = NewMapperPool(5)

type Provider interface {
	Name() string

	RegisterMapperFunc(func() Mapify)

	Start()

	Stop() bool
}

type provider struct {
	name  string
	pfunc func() Mapify
	f     *Family
	t     *time.Ticker
	close chan bool
}

type Updater interface {
	Process()
}

var instanceLock sync.RWMutex
var instances []Provider

func NewProvider(name string, pf func() Mapify, m *Family) Provider {
	ret := &provider{
		name:  name,
		pfunc: pf,
		f:     m,
		t:     time.NewTicker(1 * time.Second),
		close: make(chan bool),
	}

	instanceLock.Lock()
	instances = append(instances, ret)
	instanceLock.Unlock()

	return ret
}

func ProviderInstances() []Provider {
	instanceLock.RLock()
	defer instanceLock.RUnlock()

	return instances
}

func (p *provider) Name() string {
	return p.name
}

func (p *provider) RegisterMapperFunc(pf func() Mapify) {
	p.Stop()
	p.pfunc = pf
	p.t = time.NewTicker(1 * time.Second)
	p.Start()
}

func (p *provider) Start() {
	st := make(chan bool)
	go func(p *provider, st chan bool) {
		close(st)
	lLoop:
		for {
			select {
			case <-p.t.C:
				p.process()
			case <-p.close:
				break lLoop
			}
		}

	}(p, st)

	<-st
}

func (p *provider) Stop() bool {
	// Signal stop to make sure we clean up the ticker
	// do not access the ticker if it returns false
	if p.t != nil {
		p.close <- true
		close(p.close)
		p.t.Stop()
		return true
	}

	return false
}

func (p *provider) process() {
	out := p.pfunc()

	// Put each mapper back into the pool
	for i := 0; i < len(out); i++ {
		// f := *p.f
		val := out[i]
		fmt.Printf("Key: {%+v}, value: {%.2f}\n", val.key, val.value)
		// f.With(val.key).Set(val.value)
		mapPool.Put(val)
	}
}

type Metric interface {
	Dump()
}

type valueType struct {
	val       [MAX_SAMPLES]float64
	activeIdx uint8
}

func (v *valueType) set(f float64) {
	v.val[v.activeIdx] = f
	v.activeIdx = ((v.activeIdx + 1) % MAX_SAMPLES)
}

func (v *valueType) read(idx int) float64 {
	if idx == -1 {
		i := v.activeIdx
		if i != 0 {
			i -= 1
		}
		return v.val[i]
	} else {
		return v.val[idx]
	}
}

type Family interface {
	Set(float64)

	Read(idx int) float64
}

type Mem struct {
	valueType
	f func() float64
}

type Net struct {
	Rxb valueType
	Txb valueType
	Rxr valueType
	Txr valueType
	f   func() []float64
}

func StartPMem() Family {
	m := &Mem{}
	return m
}

func StartMem(f func() float64) Metric {
	m := &Mem{f: f}
	m.Do()
	return m
}

func (m *Mem) Set(v float64) {
	m.set(v)
}

func (m *Mem) Read(idx int) float64 {
	return m.read(idx)
}

func (m *Mem) Do() {
	wait := make(chan bool)
	tick := time.NewTicker(1 * time.Second)
	go func(m *Mem, wait chan<- bool, t *time.Ticker) {
		wait <- true
		close(wait)
		for range t.C {
			m.set(m.f())
		}
	}(m, wait, tick)

	<-wait
}

func (m *Mem) Dump() {
	fmt.Printf("Memory Values are:\t")
	for _, v := range m.val {
		fmt.Printf("%v MB, ", v)
	}
	fmt.Printf("\n")
}

func StartPNet() Family {
	return &Net{}
}

func StartNet(f func() []float64) Metric {
	n := &Net{f: f}
	n.Do()
	return n
}

func (n *Net) Do() {
	wait := make(chan bool)
	tick := time.NewTicker(1 * time.Second)
	go func(n *Net, wait chan<- bool, t *time.Ticker) {
		wait <- true
		close(wait)
		for range t.C {
			out := n.f()
			if len(out) != 2 {
				panic("invalid length of response")
			}
			rx, tx := out[0], out[1]

			prx, ptx := n.Rxb.read(-1), n.Txb.read(-1)

			n.Rxr.set((rx - prx) * 8)
			n.Txr.set((tx - ptx) * 8)
			n.Rxb.set(rx)
			n.Txb.set(tx)
		}

	}(n, wait, tick)

	<-wait
}

func (n *Net) Set(v float64) {
	return
}

func (n *Net) Read(idx int) float64 {
	return 0
}

func (n *Net) Dump() {
	fmt.Printf("Network Interface values are:")
	fmt.Printf("\n\tRx Bytes:\t")
	for i := 0; i < MAX_SAMPLES; i++ {
		fmt.Printf("%0.2f, ", n.Rxb.read(i))
	}

	fmt.Printf("\n\tTx Bytes: \t")
	for i := 0; i < MAX_SAMPLES; i++ {
		fmt.Printf("%0.2f, ", n.Txb.read(i))
	}

	fmt.Printf("\n\tRx Rate (bps): \t")
	for i := 0; i < MAX_SAMPLES; i++ {
		fmt.Printf("%0.2f, ", n.Rxr.read(i))
	}

	fmt.Printf("\n\tTx Rate (bps): \t")
	for i := 0; i < MAX_SAMPLES; i++ {
		fmt.Printf("%0.2f, ", n.Txr.read(i))
	}

	fmt.Printf("\n")
}

func getMemStats() Mapify {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	var out []*mapper

	m := mapPool.Get()
	m.set([]string{"mem"}, float64(mem.Sys/(1<<20)))
	out = append(out, m)

	return out
}

func getMemStatsFloat() float64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return float64(mem.Sys / (1 << 20))
}

func getPNetStats() Mapify {
	r := rand.New(rand.NewSource(rand.Int63n((1 << 19))))
	ifs := []string {"eth0", "eth1", "eth2", "eth3", "eth4", "eth5", "eth6"}
	var out []*mapper

	for _, v := range ifs {
		m := mapPool.Get()
		m.key = []string {v, "rxB"}
		m.value = math.Abs(r.NormFloat64())

		out = append(out, m)

		m1 := mapPool.Get()
		m1.key = []string {v, "txB"}
		m1.value = math.Abs(r.NormFloat64())

		out = append(out, m1)
	}

	return out
}

func getNetStats() []float64 {
	// Mimic what /proc/net/dev provides
	r := rand.New(rand.NewSource(rand.Int63n((1 << 19))))
	rxB := math.Abs(r.NormFloat64())
	txB := math.Abs(r.NormFloat64())

	return []float64{rxB, txB}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	var out []Metric
	out = append(out, StartMem(getMemStatsFloat))
	out = append(out, StartNet(getNetStats))

	go func(o []Metric, wg *sync.WaitGroup) {
		defer wg.Done()

		// t := time.NewTimer(40 * time.Second)
		st := make(chan bool)
		stop := make(chan bool)

		go func() {
			st <- true
			close(st)

			t := time.NewTicker(63 * time.Second)
			for range t.C {
				fmt.Println("Ended!!!!!")
				t.Stop()
				stop <- true
				close(stop)
				break
			}
		}()

		<-st
		go func(o []Metric, stop <-chan bool) {
		lLoop:
			for {
				select {
				case <-time.After(10 * time.Second):
					for _, v := range o {
						v.Dump()
					}
					fmt.Printf("\n")
				case <-stop:
					break lLoop
				}
			}
		}(o, stop)

		<-stop
	}(out, &wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		m := StartPMem()
		NewProvider("mem", getMemStats, &m).Start()

		n := StartPNet()
		NewProvider("net", getPNetStats, &n).Start()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		t := time.NewTicker(24 * time.Second)

		<-t.C
		for _, val := range ProviderInstances() {
			val.Stop()
		}

		t.Stop()
	}(&wg)

	wg.Wait()
}