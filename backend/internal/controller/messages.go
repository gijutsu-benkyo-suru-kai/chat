package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gijutsu-benkyo-suru-kai/chat/internal/dao"
)

type MessagesController struct {
	messageDao *dao.MessageDao
}

func NewMessagesController(messageDao *dao.MessageDao) *MessagesController {
	return &MessagesController{
		messageDao: messageDao,
	}
}

// GET /messages
func (c *MessagesController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	messageJson, err := json.Marshal(c.messageDao.ReadAll())
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, string(messageJson))
}

// POST /messages
func (c *MessagesController) Post(w http.ResponseWriter, r *http.Request) {}

// PUT /messages
func (c *MessagesController) Put(w http.ResponseWriter, r *http.Request) {}

// DELETE /messages
func (c *MessagesController) Delete(w http.ResponseWriter, r *http.Request) {}
