package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Message struct {
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
}

func main() {
	messageList := []Message{
		{
			UserId:   2,
			UserName: "たろう",
			Content:  "あ〜〜〜〜〜",
			Time:     1617371835000,
		},
		{
			UserId:   1,
			UserName: "じろう",
			Content:  "う〜〜〜",
			Time:     1617372835000,
		},
		{
			UserId:   2,
			UserName: "たろう",
			Content:  "画像が入る",
			Time:     1617373835000,
		},
	}

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		messageJson, err := json.Marshal(messageList)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(messageJson))
	})
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
			}
			var message Message
			err = json.Unmarshal(body, &message)
			if err != nil {
				log.Println(err)
			}
			message.Time = time.Now().Unix() * 1000
			messageList = append(messageList, message)
			log.Printf("%+v\n", messageList)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
