package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// Todo is item of todolist
type Todo struct {
	id    string
	value string
	done  bool
}

var (
	m = map[string]Todo{}
)

func main() {
	log.Println("Listening localhost:3000...")
	http.HandleFunc("/todo", todoHandler)
	http.ListenAndServe(":3000", nil)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v /todo", r.Method)

	switch r.Method {
	case http.MethodGet:
		for k, v := range m {
			fmt.Fprintf(w, "id: %s, value: %v\n", k, v)
		}
	case http.MethodPost:
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		todo := makeTodoStruct(buf.String())
		m[todo.id] = todo

		fmt.Fprintf(w, "id: %v\n", todo.id)
	}
}

func makeTodoStruct(val string) Todo {
	key := makeRandom()
	return Todo{id: key, value: val, done: false}

}

func makeRandom() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := 10

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
