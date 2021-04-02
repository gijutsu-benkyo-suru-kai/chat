import React from 'react'
import './App.css'

function App() {
  return (
    <div className="App">
      <div className="chat-wrapper">
        <div className="message-wrapper">
          <div className="message">あ〜〜〜〜〜</div>
          <div className="time">20:00</div>
        </div>
        <div className="my-message-wrapper">
          <div className="message">う〜〜〜</div>
          <div className="time">20:00</div>
        </div>
        <div className="message-wrapper">
          <div className="message">
            <div>画像が入る</div>
          </div>
          <div className="time">20:00</div>
        </div>
      </div>
    </div>
  )
}

export default App
