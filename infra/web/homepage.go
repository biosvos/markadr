package web

import "github.com/savsgio/atreugo/v11"

func homepage(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("hi")
}
