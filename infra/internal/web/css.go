package web

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"os"
	"strings"
)

func (r *router) css(ctx *atreugo.RequestCtx) error {
	filename := ctx.UserValue("file").(string)
	bytes, err := os.ReadFile(fmt.Sprintf("%v/%v", "assets/html", filename))
	if err != nil {
		return errors.WithStack(err)
	}
	setHeaderByFileSuffix(ctx, filename)
	ctx.SetBody(bytes)
	return nil
}

func setHeaderByFileSuffix(ctx *atreugo.RequestCtx, filename string) {
	switch {
	case strings.HasSuffix(filename, ".css"):
		ctx.Response.Header.SetContentType("text/css")
	case strings.HasSuffix(filename, ".js"):
		ctx.Response.Header.SetContentType("text/javascript")
	}
}
