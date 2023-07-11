package main

import (
	"github.com/biosvos/markadr/infra/web"
	"log"
)

func main() {
	server := web.NewWeb(8123)
	err := server.Run()
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panicf("%+v", err)
	}
}
