package main

import (
	"log"
	"net/http"

	"github.com/temporalio/screencast-remote-codec-server-go/codec"
	"go.temporal.io/sdk/converter"
)

func NewPayloadCodecCORSHTTPHandler(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Namespace")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := converter.NewPayloadCodecHTTPHandler(codec.NewSnappyCodec())
	handler = NewPayloadCodecCORSHTTPHandler("http://localhost:8233", handler)

	http.Handle("/", handler)

	err := http.ListenAndServe(":8234", nil)
	log.Fatal(err)
}
