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

function InputBar(props: {}) {
  return (
    <div className="input-bar">
      <div className="input-wrapper">
        <input type="text" name="text" id="text" placeholder="メッセージを入力" />
      </div>
      <div className="material-icons">send</div>
    </div>
  )
}

function App() {
  const myUserId = 1

  return (
    <div className="App">
      <div className="chat-wrapper">
        {messageList.map((message, index) => (
          <MessageWrapper key={index} message={message} myUserId={myUserId} />
        ))}
      </div>
      <InputBar />
    </div>
  )
}

export default App
