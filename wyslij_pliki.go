package main

import (
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
)

func przeslijNaSerwerPlikow(serwerPlikowURL string, sciezkaDoPliku string) (err error) {
	file, err := os.Open(sciezkaDoPliku)
	if err != nil {
		return
	}
	defer file.Close()

	res, err := http.Post(serwerPlikowURL, "binary/octet-stream", file)
	if err != nil {
		return
	}
	defer res.Body.Close()
	message, err := ioutil.ReadAll(res.Body)

	fmt.Println(message)
	return
}
