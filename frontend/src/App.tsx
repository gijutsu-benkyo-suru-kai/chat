import React from 'react';
import './App.css';
import Frame from "./Frame"

function getUserName(): string {
  return "user"
}

function App() {
  return (
    <div className="App">
      <header>
        <h1>cgi chat</h1>
        <div className="input">
          <div className="greeting">
            こんにちは、{getUserName()} さん！
          </div>
          <input placeholder="メッセージを入力" type="text"></input>
          <input type="button" value="送信"></input>
        </div>
      </header>
      <main>
        <Frame />
      </main>
    </div>
  );
}

export default App;
