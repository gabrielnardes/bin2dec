package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"bin2dec"
)

func main() {
    mux := http.NewServeMux()
    mux.Handle("/api/bin2dec/", &ApiHandler{})
    fmt.Println("bin2dec server started")
    http.ListenAndServe(":8080", mux)
}

type ApiHandler struct{}

var (
    ApiRegex = regexp.MustCompile(`^/api/bin2dec/\d+$`)
)

func (h *ApiHandler) Hello(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    binary := parts[len(parts)-1]
    result := bin2dec.Calculate(binary)
    w.Write([]byte(fmt.Sprint(result)))
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == http.MethodGet && ApiRegex.MatchString(r.URL.Path):
        h.Hello(w, r)
        return
    default:
        return
    }
}

