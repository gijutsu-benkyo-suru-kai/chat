package dao

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
	Id       uint64 `json:"-"`
	UserId   uint64 `json:"userId"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
}

type MessageDao struct {
	messageList []Message
	maximumId   uint64
}

func NewMessageDao() *MessageDao {
	return &MessageDao{
		messageList: []Message{
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
		},
		maximumId: 3,
	}
}

func (dao *MessageDao) ReadAll() *[]Message {
	log.Printf("[READ] all records")
	return &dao.messageList
}

func (dao *MessageDao) ReadById(id uint64) (*Message, error) {
	log.Printf("[READ] message.id = %d", id)
	for _, message := range dao.messageList {
		if message.Id == id {
			return &message, nil
		}
	}
	return nil, fmt.Errorf("cannot find record: Message.Id = %d", id)
}

func (dao *MessageDao) Create(message Message) error {
	dao.maximumId++
	message.Id = dao.maximumId
	message.Time = time.Now().Unix() * 1000
	dao.messageList = append(dao.messageList, message)
	log.Printf("[CREATE] message = %#v\n", message)
	return nil
}

func (dao *MessageDao) Delete(id uint64) error {
	log.Printf("[DELETE] message.id = %d", id)
	for i, message := range dao.messageList {
		if message.Id == id {
			dao.messageList = append(dao.messageList[:i], dao.messageList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("cannot find record: Message.Id = %d", id)
}
