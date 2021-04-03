import React, { useState } from "react"
import useInterval from "use-interval"
import "./Frame.css"

interface Message {
  username: string
  type: "Enter" | "Leave" | "Content"
  timestamp: number
  content: string
}

function isMessage(obj: any): obj is Message {
  return typeof obj === "object"
    && typeof obj.username === "string"
    && typeof obj.type === "string" && (obj.type === "Enter" || obj.type === "Leave" || obj.type === "Content")
    && typeof obj.timestamp === "number"
    && typeof obj.content === "string"
}

function toDateString(timestamp: number): string {
  return (new Date(timestamp)).toLocaleString()
}

function MessageDom(message: Message) {
  const username = message.username
  const timestamp = toDateString(message.timestamp)
  let content = ""
  let messageClass = "message"

  if (message.type === "Enter") {
    messageClass = `${messageClass} control`
    content = "さんが入室しました"
  } else if (message.type === "Leave") {
    messageClass = `${messageClass} control`
    content = "さんが退出しました"
  } else if (message.type === "Content") {
    content = message.content || ""
  }

  return (
    <div className={messageClass}>
      <span className="name">{username}</span>
      {content}
      <span className="timestamp">{timestamp}</span>
    </div>
  )
}


function Frame() {
  const [messages, setMessages] = useState<Message[]>([]);

  // fetch("/messages")
  //   .then(response => response.json())
  //   .then(response => {
  //     setMessages(response as Message[])
  //   })
  // useInterval(() => fetchMessages(setMessages), 5000)

  return (
    <div className="messages">
      {messages.map((message: Message) => MessageDom(message))}
    </div>
  )
}

export default Frame;
