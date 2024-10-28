package main

import (
	"fmt"
	"log"
	"net/http"
    "edge-talk/db"
)

func main() {
    err := db.Initialize()
    if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    router()

    fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)
}

