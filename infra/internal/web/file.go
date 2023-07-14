package web

import (
	"errors"
	"github.com/biosvos/markadr/assets/html"
	"github.com/savsgio/atreugo/v11"
	"strings"
)

func (r *router) sendFile(ctx *atreugo.RequestCtx) error {
	filename := ctx.UserValue("file").(string)
	bytes, err := getFileBytes(filename)
	if err != nil {
		return ctx.ErrorResponse(err, 404)
	}
	setHeaderByFileSuffix(ctx, filename)
	ctx.SetBody(bytes)
	return nil
}

func getFileBytes(filename string) ([]byte, error) {
	switch filename {
	case "kanban.css":
		return html.KanbanCSS, nil
	case "kanban.js":
		return html.KanbanJavascript, nil
	case "favicon.ico":
		return html.Favicon, nil
	default:
		return nil, errors.New("not found")
	}
}

func setHeaderByFileSuffix(ctx *atreugo.RequestCtx, filename string) {
	switch {
	case strings.HasSuffix(filename, ".css"):
		ctx.Response.Header.SetContentType("text/css")
	case strings.HasSuffix(filename, ".js"):
		ctx.Response.Header.SetContentType("text/javascript")
	case strings.HasSuffix(filename, ".ico"):
		ctx.Response.Header.SetContentType("image/png")
	}
}
