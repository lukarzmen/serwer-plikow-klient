package main

import (
	"net/http"
	"os"
)

func przeslijNaSerwerPlikow(serwerPlikowURL string, sciezkaDoPliku string, nazwaPliku string, nazwaUzytkownika string) (err error) {
	file, err := os.Open(sciezkaDoPliku)
	if err != nil {
		return
	}
	defer file.Close()

	contentDisposition := zrobContentDisposition(nazwaPliku, nazwaUzytkownika)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", serwerPlikowURL, file)
	request.Header.Set("Content-Type", "binary/octet-stream")
	request.Header.Set("Content-Disposition", contentDisposition)
	response, _ := client.Do(request)

	if err != nil {
		return
	}
	defer response.Body.Close()
	return
}

func zrobContentDisposition(nazwaPliku string, nazwaUzytkownika string) (contentDisposition string) {
	contentDisposition = "attachment; "
	contentDisposition += "filename=" + nazwaPliku + "; "
	contentDisposition += "username" + nazwaUzytkownika
	return
}
