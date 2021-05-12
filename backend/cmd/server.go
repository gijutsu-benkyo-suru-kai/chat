package main

import (
	"github.com/gijutsu-benkyo-suru-kai/chat/internal/controller"
	"github.com/gijutsu-benkyo-suru-kai/chat/internal/dao"
	"log"
	"net/http"
)

const serverAddr = "127.0.0.1:8080"

var (
	messageDao = dao.NewMessageDao()
	message    = controller.NewMessageController(messageDao)
	messages   = controller.NewMessagesController(messageDao)
)

func main() {
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			messages.Get(w, r)
		default:
			w.WriteHeader(405)
		}
	})
	http.HandleFunc("/message/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			message.Post(w, r)
		case "DELETE":
			message.Delete(w, r)
		default:
			w.WriteHeader(405)
		}
	})

	log.Printf("Listen on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
