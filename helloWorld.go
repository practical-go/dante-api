package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type catFacts struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	res, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		log.Fatalln(err)
	}

	var JSONKitty []catFacts
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &JSONKitty)
	log.Println(JSONKitty)

	//Create basic sever and log cat facts
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(JSONKitty); i++ {
			fmt.Fprintf(w, JSONKitty[i].Text)
		}
	})
	http.ListenAndServe(":8080", nil)
}
