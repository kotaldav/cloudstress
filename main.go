package main

import (
  "io"
  "log"
  "net/http"
  "time"
)

func main() {

  http.Handle("/", Logger(http.HandlerFunc(mainHandler), "WEB"))

  log.Println("I'm listening...")
  log.Fatal(http.ListenAndServe(":8000", nil))
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // Save start time to calculate duration
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t\t%s\t\t%s\t\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}

func  mainHandler(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Test!\n")
}
