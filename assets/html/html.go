package html

import _ "embed"

//go:embed index.html
var Index string

//go:embed page.html
var Page string

//go:embed page.css
var PageCSS []byte

//go:embed kanban.css
var KanbanCSS []byte

//go:embed kanban.js
var KanbanJavascrpt []byte