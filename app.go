package sample_go_app

import (
	"fmt"
	"net/http"
	"time"
)

var start = time.Now()

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "User-Agent: %s\n", r.Header.Get("User-Agent"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Current time: %s\n", time.Now().String())
	fmt.Fprintf(w, "Uptime: %s\n", time.Since(start).String())
}
