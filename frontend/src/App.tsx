import React, { useState } from 'react'
import './App.css'

type Message = {
  userId: number
  userName: string
  content: string
  time: number
}

const defaultMessageList: Message[] = [
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

function MessageWrapper(props: { message: Message; myUserId: number }) {
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
  const [text, setText] = useState('')

  const [messageList, setMessageList] = useState(defaultMessageList)

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value)
  }

  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    console.log(text)
    if (text === '') return
    setMessageList([
      ...messageList,
      {
        userId: myUserId,
        content: text,
        time: Number(new Date()),
        userName: 'じろう',
      },
    ])
    setText('')
  }

  return (
    <div className="App">
      <div className="chat-wrapper">
        {messageList.map((message, index) => (
          <MessageWrapper message={message} myUserId={myUserId} key={index} />
        ))}
        <textarea className="input-form" value={text} onChange={handleChange} />
        <button onClick={handleClick}>送信</button>
      </div>
    </div>
  )
}

export default App
