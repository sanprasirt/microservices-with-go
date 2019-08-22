package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Response struct {
	Message string
}

func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandler(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:8080/helloworld",
			"application/json",
			bytes.NewBuffer([]byte(`{"Name": "World"`)),
		)

		var response helloworldResponse
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&response)
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}
	for i := 1; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkHelloHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}
