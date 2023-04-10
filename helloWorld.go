package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		log.Fatalln(err)
	}

	type catFacts struct {
		Title   string
		Summary string
	}

	body, err := ioutil.ReadAll(res.Body)
	var data = string(body)
	var rawData any
	json.Unmarshal([]byte(data), &rawData)
	fmt.Println(rawData)

}
