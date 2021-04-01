import "./Frame.css"

interface Message {
  username: string
  type: "Enter" | "Leave" | "Content"
  timestamp: number
  content?: string
}

const messages: Message[] = [
  {
    username: "名無し",
    type: "Enter",
    timestamp: 1617279200000,
  },
  {
    username: "名無し",
    type: "Content",
    content: "こんちは",
    timestamp: 1617279302000
  },
  {
    username: "名無し",
    type: "Content",
    content: "誰もいないんか…",
    timestamp: 1617279304000
  },
  {
    username: "名無し",
    type: "Leave",
    timestamp: 1617279304000
  }
]

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
  messages.reverse()
  return (
    <div className="messages">
      {messages.map((message: Message) => MessageDom(message))}
    </div>
  )
}

export default Frame;
