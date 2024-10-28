package main

import (
	web "edge-talk/web/base"
	"net/http"

	"github.com/a-h/templ"
)

func router() {
    component := web.Base("Edge talk")
    http.Handle("/", templ.Handler(component))
}
