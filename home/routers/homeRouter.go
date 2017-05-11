package routers

import (
    "net/http"
    "github.com/msfy/fengyun.cc/home/templates"
    "encoding/json"
)

type User struct {
    Name string
}

func (u User) MarshalBinary() ([]byte, error) {
    return json.Marshal(u)
}

func homeRouter(mux *http.ServeMux) {
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        templates.Render(w, templates.HomeTemplate, nil)
    })
}
