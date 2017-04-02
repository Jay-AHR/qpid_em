package main

import (
	"github.com/Jay-AHR/qpid_em/gobot"
	"github.com/Jay-AHR/qpid_em/http"
	"github.com/Jay-AHR/qpid_em/log"
	"github.com/Jay-AHR/qpid_em/prometheus"
	"github.com/Jay-AHR/qpid_em/twillio"
)

func main() {
	gb := gobot.NewController()

	l := log.New()

	p := prometheus.NewSink()

	t := twillio.New()

	s := http.NewServer(gb, t, l, p)
	err := s.ListenAndServe(":8081")
	if err != nil {
		panic(err)
	}
}
