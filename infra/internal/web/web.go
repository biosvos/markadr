package web

import (
	"fmt"
	"github.com/biosvos/markadr/flow"
	"github.com/savsgio/atreugo/v11"
)

type router struct {
	navigator *flow.Navigator
}

type Web struct {
	navigator *flow.Navigator
	server    *atreugo.Atreugo
}

func NewWeb(port uint16, navigator *flow.Navigator) *Web {
	address := fmt.Sprintf(":%v", port)
	server := atreugo.New(atreugo.Config{
		Addr: address,
	})
	return &Web{
		server:    server,
		navigator: navigator,
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
	routing := &router{
		navigator: w.navigator,
	}
	w.server.GET("/", routing.homepage)
	w.server.GET("/pages/{title}", routing.page)
}
