package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gijutsu-benkyo-suru-kai/chat/internal/controller"
	"github.com/gijutsu-benkyo-suru-kai/chat/internal/dao"
)

const (
	defaultPort = "8080"
	defaultHost = "0.0.0.0"
)

var (
	messageDao = dao.NewMessageDao()
	message    = controller.NewMessageController(messageDao)
	messages   = controller.NewMessagesController(messageDao)
)

func generateServerAddr() string {
	port := defaultPort
	host := defaultHost
	if value, ok := os.LookupEnv("SERVER_HOST"); ok {
		host = value
	}
	if value, ok := os.LookupEnv("SERVER_PORT"); ok {
		port = value
	}
	return host + ":" + port
}

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

	serverAddr := generateServerAddr()
	log.Printf("Listen on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
