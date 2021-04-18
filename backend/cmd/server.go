package main

import (
	"encoding/json"
	"fmt"
	"github.com/gijutsu-benkyo-suru-kai/chat/internal"
	"io/ioutil"
	"log"
	"net/http"
)

var messageDao = internal.NewMessageDao()

func getMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	messageJson, err := json.Marshal(messageDao.ReadAll())
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, string(messageJson))
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		message := internal.Message{}
		err = json.Unmarshal(body, &message)
		if err != nil {
			log.Println(err)
		}

		if err = messageDao.Create(message); err != nil {
			log.Println(err)
		}

		log.Printf("%+v\n", messageDao.ReadAll())
	}
}

func main() {

	http.HandleFunc("/messages", getMessages)
	http.HandleFunc("/message", postMessage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
