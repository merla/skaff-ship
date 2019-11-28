package main

import (
	"io/ioutil"
	"net/http"

	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	_, _ = w.Write([]byte(fmt.Sprintf("SIGNED-%s", string(body))))
}

func main() {
	log.Print("Signer is ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8440", nil)
}
