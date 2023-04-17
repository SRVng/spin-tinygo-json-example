package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go_spin/message"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		m := message.Message{Text: "Hello, World!"}
		out, err := m.MarshalJSON()

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, string(out))
		}
	})
}

func main() {}
