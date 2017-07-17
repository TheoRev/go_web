package mux

import (
	"net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

type MyMux struct {
	myRutes map[string]customeHandler
}

func (this *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := this.myRutes[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (this *MyMux) AddFunc(ruta string, fn customeHandler) {
	this.myRutes[ruta] = fn
}
func (this *MyMux) AddHandle(ruta string, handle http.Handler) {
	this.myRutes[ruta] = handle.ServeHTTP
}

func CreateMux() *MyMux {
	newMap := make(map[string]customeHandler)
	return &MyMux{newMap}
}
