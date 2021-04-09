import React from 'react'
import './App.css'

type Message = {
  userId: number
  userName: string
  content: string
  time: number
}

const messageList: Message[] = [
  {
    userId: 2,
    userName: 'たろう',
    content: 'あ〜〜〜〜〜',
    time: 1617371835000,
  },
  {
    userId: 1,
    userName: 'じろう',
    content: 'う〜〜〜',
    time: 1617372835000,
  },
  {
    userId: 2,
    userName: 'たろう',
    content: '画像が入る',
    time: 1617373835000,
  },
]

function MessageWrapper(props: {
  message: Message
  myUserId: number
}) {
  const time = new Date(props.message.time)
  if (props.message.userId === props.myUserId) {
    return (
      <div className="my-message-wrapper">
        <div className="message">{props.message.content}</div>
        <div className="time">{`${time.getHours()}:${time.getMinutes()}`}</div>
      </div>
    )
  } else {
    return (
      <div className="message-wrapper">
        <div className="message">{props.message.content}</div>
        <div className="time">{`${time.getHours()}:${time.getMinutes()}`}</div>
      </div>
    )
  }
}

function App() {
  const myUserId = 1

  return (
    <div className="App">
      <div className="chat-wrapper">
        {messageList.map((message, index) => (
          <MessageWrapper message={message} myUserId={myUserId} key={index} />
        ))}
      </div>
    </div>
  )
}

export default App
