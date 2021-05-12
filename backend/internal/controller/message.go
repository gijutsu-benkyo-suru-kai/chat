package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gijutsu-benkyo-suru-kai/chat/internal/dao"
)

type MessageController struct {
	messageDao *dao.MessageDao
}

func NewMessageController(messageDao *dao.MessageDao) *MessageController {
	return &MessageController{
		messageDao: messageDao,
	}
}

// GET /message
func (c *MessageController) Get(w http.ResponseWriter, r *http.Request) {}

// POST /message
func (c *MessageController) Post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	log.Printf("body: %#v\n", body)

	message := dao.Message{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	if err = c.messageDao.Create(message); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

// PUT /message
func (c *MessageController) Put(w http.ResponseWriter, r *http.Request) {}

// DELETE /message/:id
func (c *MessageController) Delete(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter :id
	substr := strings.TrimPrefix(r.URL.Path, "/message")
	_, id_str := filepath.Split(substr)
	id, err := strconv.ParseUint(id_str, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	log.Printf("[DELETE /message] :id = %d\n", id)

	if err := c.messageDao.Delete(id); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
