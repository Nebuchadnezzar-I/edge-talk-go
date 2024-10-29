package main

import (
	"edge-talk/db"
	web "edge-talk/web/base"
	"edge-talk/web/components"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func router() {
    // fetch negotiations

    negotiations, err := db.FetchNegotiations()
    if err != nil {
        log.Fatalf("Failed to fetch negotiations: %v", err)
    }

    wrapper := components.Wrapper(negotiations)
    component := web.Base("Edge talk", wrapper)
    http.Handle("/", templ.Handler(component))

    newContent := components.Content()
    http.Handle("/new-content", templ.Handler(newContent))
}
