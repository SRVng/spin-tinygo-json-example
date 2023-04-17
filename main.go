package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go_spin/message"
	"github.com/mailru/easyjson"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()

		router.GET("/", get_greeting)
		router.POST("/", post_echo)

		router.ServeHTTP(w, r)
	})
}

func get_greeting(w http.ResponseWriter, r *http.Request, ps spinhttp.Params) {
	m := &message.Message{Text: "Hello, World!"}

	w.Header().Set("Content-Type", "application/json")
	easyjson.MarshalToHTTPResponseWriter(m, w)
}

func post_echo(w http.ResponseWriter, r *http.Request, ps spinhttp.Params) {
	m := &message.Message{}
	// Read key "text" from body
	err := easyjson.UnmarshalFromReader(r.Body, m)

	if len(m.Text) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Missing key 'text' in request body", http.StatusUnprocessableEntity)
		return
	}

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		m = &message.Message{Text: m.Text + " ECHO"}
		easyjson.MarshalToHTTPResponseWriter(m, w)
		return
	}
}

func main() {}
