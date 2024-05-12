package server

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type HTTPServer struct {
	addr   string
	router *fasthttprouter.Router
}

func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{addr: addr, router: fasthttprouter.New()}
}

func (s *HTTPServer) RegisterHandler(method, path string, handler fasthttp.RequestHandler) {
	s.router.Handle(method, path, handler)
}

func (s *HTTPServer) Start() {
	fmt.Println("Server is running on", s.addr)
	log.Fatal(fasthttp.ListenAndServe(s.addr, s.router.Handler))
}
