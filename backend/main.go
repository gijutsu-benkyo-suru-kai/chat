package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     uint64 `json:"time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
		messageJson, _ := json.Marshal(messageList)

		fmt.Fprintf(w, string(messageJson))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
