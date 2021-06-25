package main

import (
	"expvar"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Dump struct {
	MemStats interface{}
	APICount interface{}
}

func main() {
	var tmp = expvar.NewInt("tempapi")
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("templates/*.html"))

		err := tpl.Execute(w, &Dump{MemStats: expvar.Get("memstats"), APICount: expvar.Get("tempapi")})
		if err != nil {
			log.Fatalf("Error in executing: ", err)
		}
	})

	http.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {
		go func(w http.ResponseWriter, r *http.Request) {
			tmp.Add(1)
			defer tmp.Add(-1)

		lLoop:
			for {
				select {
				case <-time.After(10 * time.Second):
					break lLoop
				}
			}

			w.Write([]byte("Done"))
		}(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
