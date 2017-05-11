package main

import (
    "log"
    "net/http"
    . "github.com/msfy/fengyun.cc/ttd2/routers"
)

func main() {

    mux := http.NewServeMux()

    Routers(mux)

    log.Print("App is listening port :8884")
    log.Fatal(http.ListenAndServe(":8884", mux))
}
