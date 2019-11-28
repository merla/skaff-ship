package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"fmt"
	"log"
)

func handleMakeDonut(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	signed := checkRequestSignature(body)
	if !signed {
		fmt.Println("Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid request"))
		return
	}

	fmt.Println("Making donut...")
	_, _ = w.Write([]byte("Delicious Jelly donut"))
}

func main() {
	log.Print("Donut shop is open")
	http.HandleFunc("/", handleMakeDonut)
	http.ListenAndServe(":8441", nil)
}

func checkRequestSignature(b []byte) bool {
	return strings.Contains(string(b), "SIGNED")
}
