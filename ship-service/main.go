package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"bytes"
	"log"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	signed, err := request("http://sign-service:8440", strings.NewReader("Jelly donut"))
	if err != nil {
		panic(err)
	}

	donut, err := request("http://donut-service-prod:8441", bytes.NewReader(signed))
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(w, bytes.NewReader(donut)); err != nil {
		panic(err)
	}
}

func main() {
	log.Print("Shipper is ready 1")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8442", nil)
}

func request(url string, b io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
