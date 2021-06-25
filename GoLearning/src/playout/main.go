package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Req struct {
	p string `xml:"p"`
	a string `xml:"a"`
	s string `xml:"s"`
	r uint32 `xml:"r"`
}

type Rsp struct {
	tp string `xml:"tp"`
	ti string `xml:"ti"`
	to string `xml:"to"`
}

func (r *Req) Build(p, a, s string, r uint32) {
	r.p, r.a, r.s, r.r = p, a, s, r
}

func ParseRequest(buf []byte) *Req {
	req := Req{}
	if err := xml.Unmarshal(buf, req); err != nil {
		log.Printf("Failed to Unmarshal request: %v\n", err)
	}

	return &req
}

func (r *Rsp) Build(tp, ti, to string) {
	r.tp, r.ti, r.to = tp, ti, to
}

func ParseResponse(buf []byte) *Res {
	res := Res{}
	if err := xml.Unmarshal(buf, res); err != nil {
		log.Printf("Failed to unmarshal the response: %v\n", err)
	}
}

func main() {
	fmt.Println("hi")
}
