package main

import (
	"encoding/json"
	"net/http"
	"time"
	"os"
	"fmt"
	"go/scanner"
)

var Key struct {
	token string `json:"token"`
}

func main() {
	getApiKey()
	//getJson("http://api.openweathermap.org/data/2.5/forecast", Key)
}

func getApiKey() {
	token, _ := os.Open("./keys.json")
	defer token.Close()

	decoder := json.NewDecoder(token)

	if err = decoder.Decode(&Key); err != nil {
		scanner.PrintError()
	}

	fmt.Println(Key)
	return
}

var client = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := client.Get(url)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
