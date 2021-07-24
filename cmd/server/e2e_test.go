// nolint:testpackage
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
)

func Test_server(t *testing.T) {
	srv := Server()
	type arg struct {
		path   string
		method string
		body   io.Reader
	}

	tests := []struct {
		name string
		arg  arg
		want int
	}{
		{
			name: "healthcheck",
			arg: arg{
				path:   "/health",
				method: http.MethodGet,
				body:   nil,
			},
			want: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			l, err := net.Listen("tcp", ":0") // nolint:gosec
			if err != nil {
				log.Fatal(err)
			}

			idleConnsClosed := make(chan struct{})
			go func() {
				if err2 := srv.Serve(l); err2 != nil && !errors.Is(err2, http.ErrServerClosed) {
					t.Errorf("failed boot: %v", err2)
				}

				close(idleConnsClosed)
			}()

			path := fmt.Sprintf("http://%s%s", l.Addr().String(), tt.arg.path)

			req, err := http.NewRequestWithContext(ctx, tt.arg.method, path, tt.arg.body)
			if err != nil {
				t.Fatal(err)
			}
			defer closeReqSafe(req)

			res, err := http.DefaultClient.Do(req) // nolint: bodyclose
			if err != nil {
				t.Fatal(err)
			}
			defer closeResSafe(res)

			if res.StatusCode != tt.want {
				t.Fatalf("%s %s returns get %d want %d", tt.arg.method, tt.arg.path, res.StatusCode, tt.want)
			}

			if err := srv.Shutdown(context.Background()); err != nil {
				t.Fatalf("failed to shutdown: %v", err)
			}

			<-idleConnsClosed
		})
	}
}

func closeReqSafe(req *http.Request) {
	if req != nil && req.Body != nil {
		req.Body.Close()
	}
}

func closeResSafe(res *http.Response) {
	if res != nil && res.Body != nil {
		res.Body.Close()
	}
}
