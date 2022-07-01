package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testt/memory"
	"testt/presist"

	"github.com/go-chi/chi/v5"
)

type Store interface {
	Set(string, any)
	Get(string) (any, error)
	Delete(string)
}

type KeyValPair struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func getValue(w http.ResponseWriter, r *http.Request) {
	val, err := myStore.Get(chi.URLParam(r, "key"))
	if err != nil {
		resBody := formatResponseBody(nil)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write(resBody)
		return
	}

	resBody := formatResponseBody(val)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resBody)
}

func setValue(w http.ResponseWriter, r *http.Request) {
	var data KeyValPair
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	myStore.Set(data.Key, data.Value)
	w.WriteHeader(201)
}

func deleteValue(w http.ResponseWriter, r *http.Request) {
	// _, err := inMemoryStore.get(chi.URLParam(r, "key"))
	// if err != nil {
	// 	w.WriteHeader(404)
	// 	return
	// }
	myStore.Delete(chi.URLParam(r, "key"))
	w.WriteHeader(200)
}

func formatResponseBody(data any) []byte {
	responseBody, err := json.Marshal(struct {
		Value any `json:"value"`
	}{Value: data})

	if err != nil {
		log.Fatal(err)
	}

	return responseBody
}

var myStore Store

func main() {
	if len(os.Args) > 1 {
		flag := os.Args[1]
		switch flag {
		case "-p":
			myStore = presist.Connect()
		case "-m":
			myStore = memory.InMemoryStore
		case "-h":
			fmt.Println("use -p to presist in database or -m to presist in memory")
			return
		default:
			log.Fatal("Invalid input for help -h")
		}
	} else {
		log.Fatal("Invalid input for help -h")
	}

	r := chi.NewRouter()

	r.Get("/{key}", getValue)
	r.Post("/", setValue)
	r.Delete("/{key}", deleteValue)

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}

}

/*

in-memory

- rest http (done)
-  socket tcp
- grpc


presist

- rest http (done)
-  socket tcp
- grpc

*/
