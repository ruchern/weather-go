package main

import (
	"encoding/json"
	"net/http"
	"time"
	"os"
)

var Key struct {
	Token string `json:"token"`
}

func main() {
	getApiKey()
	//url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?id=524901&APPID=%s", Key.Token)
	//weather := getJson(url, &Weather{})

	//fmt.Println(weather)
}

func getApiKey() {
	token, _ := os.Open("./keys.json")
	defer token.Close()

	decoder := json.NewDecoder(token)
	decoder.Decode(&Key)

	return
}

var client = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	response, err := client.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}
