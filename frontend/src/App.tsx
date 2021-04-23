import React, { useEffect, useState } from 'react'
import './App.css'

type Message = {
  userId: number
  userName: string
  content: string
  time: number
}

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
let myUserId: number

function App() {
  const [text, setText] = useState('')
  const fetchMessages = async () => {
    const response = await (await fetch('/messages')).json()
    return response as Message[]
  }
  const [messageList, setMessageList] = useState([] as Message[])
  useEffect(() => {
    setInterval(() => fetchMessages().then((res) => setMessageList(res)), 1000)
    myUserId = Math.ceil(Math.random() * 10000000)
  }, [])

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value)
  }

  const handleClick = async (e: React.MouseEvent<HTMLButtonElement>) => {
    if (text === '') return
    const newMessage = {
      userId: myUserId,
      content: text,
      time: Number(new Date()),
      userName: 'じろう',
    }
    setMessageList([...messageList, newMessage])
    setText('')
    const response = await fetch('/message', {
      method: 'POST',
      body: JSON.stringify(newMessage),
    })
    fetchMessages().then((res) => setMessageList(res))
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
