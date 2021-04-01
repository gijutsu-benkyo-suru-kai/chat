import "./Frame.css"

interface BaseMessage {
  username: string
  timestamp: number
}

interface ControlMessage extends BaseMessage {
  type: "Enter" | "Leave"
}
function isControlMessage(message: Message): message is ControlMessage {
  return (message as ControlMessage).type === "Enter" || (message as ControlMessage).type === "Leave"
}

interface UserMessage extends BaseMessage {
  content: string
}
function isUserMessage(message: Message): message is UserMessage {
  return typeof (message as UserMessage).content === "string"
}

type Message = ControlMessage | UserMessage


const messages: Message[] = [
  {
    username: "名無し",
    type: "Enter",
    timestamp: 1617279200000
  } as ControlMessage,
  {
    username: "名無し",
    content: "こんちは",
    timestamp: 1617279302000
  } as UserMessage,
  {
    username: "名無し",
    content: "誰もいないんか…",
    timestamp: 1617279304000
  } as UserMessage,
  {
    username: "名無し",
    type: "Leave",
    timestamp: 1617279304000
  } as ControlMessage
]

function toDateString(timestamp: number): string {
  return (new Date(timestamp)).toLocaleString()
}

function MessageDom(message: Message) {
  const username = message.username
  const timestamp = toDateString(message.timestamp)
  let content = ""
  let messageClass = "message"

  if (isControlMessage(message)) {
    messageClass = `${messageClass} control`
    if (message.type === "Enter") {
      content = "さんが入室しました"
    } else if (message.type === "Leave") {
      content = "さんが退出しました"
    }
  } else if (isUserMessage(message)) {
    content = message.content
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
