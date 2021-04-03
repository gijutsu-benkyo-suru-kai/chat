package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageType string

const (
	Enter   MessageType = "Enter"
	Leave   MessageType = "Leave"
	Content MessageType = "Content"
)

type Message struct {
	UserName  string      `json:"username"`
	Type      MessageType `json:"type"`
	Timestamp int64       `json:"timestamp"`
	Content   string      `json:"content"`
}

const PORT = 8080

var messages = []Message{
	{
		UserName:  "名無し",
		Type:      Enter,
		Timestamp: 1617279200,
	},
	{
		UserName:  "名無し",
		Type:      Content,
		Content:   "こんちは",
		Timestamp: 1617279302,
	},
	{
		UserName:  "名無し",
		Type:      Content,
		Content:   "誰もいないんか…",
		Timestamp: 1617279304,
	},
	{
		UserName:  "名無し",
		Type:      Leave,
		Timestamp: 1617279304,
	},
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Golang"))
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	jsonStr, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to JSON marshal"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/messages", handleMessages)

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", "127.0.0.1", PORT),
	}

	log.Printf("Listen at %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
