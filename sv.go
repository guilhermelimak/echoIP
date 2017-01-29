package main

import (
    "net/http"
    "strings"
    "fmt"
    "log"
)


func main() {
    PORT := ":8080"

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        s := strings.Split(r.RemoteAddr, ":")
        ip := s[0]
        fmt.Fprintf(w, "%s\n", ip)
    })

    log.Printf("Listening on port %s", PORT)
    log.Fatal(http.ListenAndServe(PORT, nil))
}
