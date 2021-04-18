package internal

import (
	"time"
)

type Message struct {
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
}

type MessageDao struct {
	messageList []Message
}

func NewMessageDao() *MessageDao {
	return &MessageDao{
		messageList: []Message{
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
		},
	}
}

func (dao *MessageDao) ReadAll() *[]Message {
	return &dao.messageList
}

func (dao *MessageDao) Create(message Message) error {
	message.Time = time.Now().Unix() * 1000
	dao.messageList = append(dao.messageList, message)
	return nil
}
