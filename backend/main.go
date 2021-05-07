package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
}

func main() {
	messageList := []Message{
		{
			Id:       1,
			UserId:   2,
			UserName: "たろう",
			Content:  "あ〜〜〜〜〜",
			Time:     1617371835000,
		},
		{
			Id:       2,
			UserId:   1,
			UserName: "じろう",
			Content:  "う〜〜〜",
			Time:     1617372835000,
		},
		{
			Id:       3,
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
			message := Message{}
			err = json.Unmarshal(body, &message)

			if err != nil {
				log.Println(err)
			}
			message.Time = time.Now().Unix() * 1000
			message.Id = uint(len(messageList)) + 1
			messageList = append(messageList, message)
			log.Printf("%+v\n", messageList)
		}
	})
	http.HandleFunc("/message/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(405)
			return
		}
		// eq. '/message/3' -> '/3'
		substr := strings.TrimPrefix(r.URL.Path, "/message")
		_, id_str := filepath.Split(substr)
		id, err := strconv.ParseUint(id_str, 10, 64)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		found := false
		for i, message := range messageList {
			if message.Id == uint(id) {
				// messageList[i] の要素を削除
				messageList = append(messageList[:i], messageList[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(200)
		return
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
