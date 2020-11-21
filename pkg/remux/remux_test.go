package remux

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"
)

func TestReMux_RegisterPlain(t *testing.T) {
	type fields struct {
		mu              sync.RWMutex
		plain           map[Method]map[string]http.Handler
		regex           map[Method]map[*regexp.Regexp]http.Handler
		notFoundHandler http.Handler
	}

	remux := New()

	if err := remux.RegisterPlain(GET, "/get", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(GET))
	})); err != nil {
		t.Fatal(err)
	}
	if err := remux.RegisterPlain(POST, "/post", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(POST))
	})); err != nil {
		t.Fatal(err)
	}

	type args struct {
		method      Method
		path        string
		handler     http.Handler
		middlewares []Middleware
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "GET", args: args{method: GET, path: "/get"}, want: []byte(GET)},
		{name: "POST", args: args{method: POST, path: "/post"}, want: []byte(POST)},
	}
	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		remux.ServeHTTP(response, request)
		got := response.Body.Bytes()
		if !bytes.Equal(tt.want, got) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func TestReMux_NotFound(t *testing.T) {
	remux := New()

	type args struct {
		method Method
		path   string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "GET", args: args{method: GET, path: "/get"}, want: http.StatusNotFound},
		{name: "POST", args: args{method: POST, path: "/post"}, want: http.StatusNotFound},
	}

	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		remux.ServeHTTP(response, request)
		got := response.Result().StatusCode
		if tt.want != got {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
