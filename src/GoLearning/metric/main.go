package main

import (
	"expvar"
	_ "expvar"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type Stats struct {
	SysMem   uint64
	IndexCnt interface{}
	CbrCnt   interface{}
}

type Updater struct {
	sync.RWMutex
	lastUpdateTime uint32
}

var indexCnt = expvar.NewInt("indexcounter")
var cbrCnt = expvar.NewInt("cbrcounter")
var stat = &Stats{}

func (s *Stats) sample() {
	var mStat runtime.MemStats
	runtime.ReadMemStats(&mStat)
	s.SysMem = (mStat.Sys >> 20)
	s.IndexCnt = expvar.Get("indexcounter")
	s.CbrCnt = expvar.Get("cbrcounter")
}

func (u *Updater) statsUpdater() {
	curTime := uint32(time.Now().UnixNano() / int64(time.Nanosecond*time.Millisecond))
	if curTime-u.lastUpdateTime < INTERVAL_IN_MS {
		return
	}

	u.Lock()
	defer u.Unlock()
	stat.sample()

	u.lastUpdateTime = uint32(time.Now().UnixNano() / int64(time.Nanosecond*time.Millisecond))
}

func main() {
	updater := make(chan bool)
	go func() {
		// Start the 1 second scheduler to update the samples
		updater <- true
		up := Updater{
			lastUpdateTime: uint32(time.Now().UnixNano() / int64(time.Nanosecond*time.Millisecond)),
		}
		for {
			select {
			case <-time.After(INTERVAL_IN_MS * time.Millisecond):
				up.statsUpdater()
			}
		}
	}()

	http.HandleFunc("/load", loadHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/cbr", cbrHandler)

	fmt.Println("Server is listening on port 8000")
	<-updater
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	timer := make(chan bool)
	go func(w http.ResponseWriter, r *http.Request) {
		indexCnt.Add(1)
		defer indexCnt.Add(-1)

		timer <- true
		select {
		case <-time.After(5 * time.Second):
			break
		}
	}(w, r)

	<-timer
	w.Write([]byte("index succeed"))
}

func cbrHandler(w http.ResponseWriter, r *http.Request) {
	timer := make(chan bool)
	go func(w http.ResponseWriter, r *http.Request) {
		cbrCnt.Add(1)
		defer cbrCnt.Add(-1)

		timer <- true
		select {
		case <-time.After(10 * time.Second):
			break
		}
	}(w, r)

	<-timer
	w.Write([]byte("play succeed"))
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	if err := tpl.Execute(w, stat); err != nil {
		log.Fatalf("Failed to generate the template: %v\n", err)
	}
}
