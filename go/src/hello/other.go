package main

import (
        "net/http"
)

func write_text(wrt http.ResponseWriter, r *http.Request) {
        wrt.Write([]byte("This is a GO Line HTTP Server\n"))
        wrt.Write([]byte("Hello World"))
}

func main() {
        http.HandleFunc("/", write_text)
        if err := http.ListenAndServe(":8080", nil); err != nil {
                panic(err)
        }
}
