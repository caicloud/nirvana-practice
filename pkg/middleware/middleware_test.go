package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/service"
)

// responseWriter is a trivial implementation of service.ResponseWriter from Nirvana; it makes it
// convenient to build fake HTTP requests for testing purposes.
// TODO: move this to a dedicated package for testing utilities so it can be reused.
type responseWriter struct {
	code   int
	header http.Header
	buf    *bytes.Buffer
}

func newRW() *responseWriter {
	return &responseWriter{0, http.Header{}, bytes.NewBuffer(nil)}
}

func (r *responseWriter) Header() http.Header {
	return r.header
}

func (r *responseWriter) Write(d []byte) (int, error) {
	return r.buf.Write(d)
}

func (r *responseWriter) WriteHeader(code int) {
	r.code = code
}

func mustURL(rawURL string) *url.URL {
	u, _ := url.Parse(rawURL)
	return u
}

func TestReqlog(t *testing.T) {
	buildService := func() (service.Service, error) {
		builder := service.NewBuilder()
		builder.SetModifier(service.FirstContextParameter())
		if err := builder.AddDescriptor(
			definition.Descriptor{
				Path:        "/api/v1",
				Definitions: []definition.Definition{},
				Consumes:    []string{"application/json"},
				Produces:    []string{"application/json"},
				Middlewares: []definition.Middleware{Reqlog(log.NewStdLogger(0))},
				Children: []definition.Descriptor{
					{
						Path: "/foo",
						Definitions: []definition.Definition{
							{
								Method: definition.Create,
								Parameters: []definition.Parameter{
									definition.QueryParameterFor("query", ""),
									definition.BodyParameterFor(""),
								},
								Function: func(ctx context.Context, query string, body []byte) (string, error) {
									request := service.HTTPContextFrom(ctx).Request()
									return fmt.Sprintf(
										"%s %s %s %s",
										request.Method, request.URL.String(),
										query, string(body),
									), nil
								},
								Results: definition.DataErrorResults(""),
							},
						},
					},
				},
			},
		); err != nil {
			return nil, err
		}
		return builder.Build()
	}
	tests := []struct {
		req       http.Request
		wantRegex string
		respStr   string
		respCode  int
	}{
		{
			req: http.Request{
				Method: http.MethodPost,
				URL:    mustURL("/api/v1/foo?query=bar"),
				Header: http.Header{
					"Content-Type": []string{"application/json"},
					"Accept":       []string{"application/json"},
				},
				Body: ioutil.NopCloser(strings.NewReader("{\"key\": \"value\"}")),
			},
			wantRegex: `^INFO\s(.+) | POST 201 61 [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+ /api/v1/foo?query=bar$`,
			respStr:   `POST /api/v1/foo?query=bar bar {"key": "value"}`,
			respCode:  201,
		},
	}
	for _, tc := range tests {
		old := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w
		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		svc, err := buildService()
		if err != nil {
			t.Fatal(err)
		}
		rw := newRW()
		svc.ServeHTTP(rw, &tc.req)

		w.Close()
		out := <-outC
		os.Stderr = old

		if rw.buf.String() != tc.respStr {
			t.Fatal("unexpected response content, the middleware might have altered the request/response")
		}
		if rw.code != tc.respCode {
			t.Fatal("unexpected response code, the middleware might have altered the request/response")
		}
		if !regexp.MustCompile(tc.wantRegex).MatchString(out) {
			t.Fatalf("unexpected log content: \"%s\"", out)
		}
	}
}
