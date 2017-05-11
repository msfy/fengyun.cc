package main

import (
    "log"
    "net/http"
    . "github.com/msfy/fengyun.cc/home/routers"
)

func main() {

    mux := http.NewServeMux()

    Routers(mux)

    log.Print("App is listening port :8885")
    log.Fatal(http.ListenAndServe(":8885", mux))
}
