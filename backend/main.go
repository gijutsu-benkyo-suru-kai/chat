package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[
  {
    "userId": 2,
    "userName": "たろう",
    "content": "あ〜〜〜〜〜",
    "time": 1617371835000
  },
  {
    "userId": 1,
    "userName": "じろう",
    "content": "う〜〜〜",
    "time": 1617372835000
  },
  {
    "userId": 2,
    "userName": "たろう",
    "content": "画像が入る",
    "time": 1617373835000
  }
]`)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
