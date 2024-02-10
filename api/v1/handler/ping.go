package handler

import "net/http"

type PingHandler struct{}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
