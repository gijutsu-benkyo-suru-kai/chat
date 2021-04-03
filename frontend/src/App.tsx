import React, { useState } from 'react'
import './App.css'

type Message = {
  userId: number
  userName: string
  content: string
  time: number
}

type OnChangeHandler = (event: React.ChangeEvent<HTMLInputElement>) => void

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


function MessageWrapper(props: { message: Message, myUserId: number }) {
  const time = new Date(props.message.time)
  const wrapperClass = (props.message.userId === props.myUserId) ? "my-message-wrapper" : "message-wrapper"
  return (
    <div className={wrapperClass}>
      <div className="message">{props.message.content}</div>
      <div className="time">{`${time.getHours()}:${time.getMinutes()}`}</div>
    </div>
  )
}

function InputBar(props: {
  value: string
  onChange: OnChangeHandler
}) {
  return (
    <div className="input-bar">
      <div className="input-wrapper">
        <input type="text" name="text" id="text" placeholder="メッセージを入力" value={props.value} onChange={props.onChange} />
      </div>
      <div className="material-icons">send</div>
    </div>
  )
}

function App() {
  const myUserId = 1

  const [text, setText] = useState("")

  const handleChangeText: OnChangeHandler = (e) => setText(e.target.value)

  return (
    <div className="App">
      <div className="chat-wrapper">
        {messageList.map((message, index) => (
          <MessageWrapper key={index} message={message} myUserId={myUserId} />
        ))}
      </div>
      <InputBar value={text} onChange={handleChangeText} />
    </div>
  )
}

export default App
