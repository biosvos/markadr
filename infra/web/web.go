package web

import (
	"fmt"
	"github.com/savsgio/atreugo/v11"
)

type Web struct {
	server *atreugo.Atreugo
}

func NewWeb(port uint16) *Web {
	address := fmt.Sprintf(":%v", port)
	server := atreugo.New(atreugo.Config{
		Addr: address,
	})
	return &Web{
		server: server,
	}
}

func (w *Web) Run() error {
	w.routing()
	return w.run()
}

func (w *Web) run() error {
	return w.server.ListenAndServe()
}

func (w *Web) routing() {
	w.server.GET("/", homepage)
}
