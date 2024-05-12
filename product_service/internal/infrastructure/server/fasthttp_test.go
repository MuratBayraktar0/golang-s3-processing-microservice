package server

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/valyala/fasthttp"
)

func TestHTTPServer_RegisterHandler(t *testing.T) {
	Convey("Given a new HTTP server", t, func() {
		server := NewHTTPServer(":8080")

		Convey("When registering a handler for GET /hello", func() {
			handler := func(ctx *fasthttp.RequestCtx) {
				fmt.Fprint(ctx, "Hello, World!")
			}

			server.RegisterHandler("GET", "/hello", handler)

			Convey("Then the handler should be registered)", func() {
				req := &fasthttp.RequestCtx{}
				req.Request.Header.SetMethod("GET")
				req.Request.SetRequestURI("/hello")

				server.router.Handler(req)

				expected := "Hello, World!"
				So(string(req.Response.Body()), ShouldEqual, expected)
			})
		})
	})
}
