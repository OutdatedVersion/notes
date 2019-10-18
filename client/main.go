package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	requestBody := map[string]string{"content": strings.Join(os.Args[1:], " ")}
	json, err := json.Marshal(requestBody)

	if err != nil {
		log.Fatalln("We could not convert the input into JSON", err)
	}

	response, err := http.Post("https://bens.wtf/notes/api/@me", "application/json", bytes.NewBuffer(json))

	if err != nil {
		log.Fatalln("We could not make the request", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("Failed to read response body", err)
	}

	log.Print("http response: ", string(body))
}
