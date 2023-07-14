package web

import (
	"fmt"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/savsgio/atreugo/v11"
)

type router struct {
	repository repository.Repository
}

type Web struct {
	repository repository.Repository
	server     *atreugo.Atreugo
}

func NewWeb(port uint16, repository repository.Repository) *Web {
	address := fmt.Sprintf(":%v", port)
	server := atreugo.New(atreugo.Config{
		Addr: address,
	})
	return &Web{
		server:     server,
		repository: repository,
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
		repository: w.repository,
	}
	w.server.GET("/", routing.homepage)
	w.server.GET("/pages/{title}", routing.page)
	w.server.GET("/{file}", routing.sendFile)
	w.server.PUT("/pages/{title}", routing.updateADRStatus)
}
