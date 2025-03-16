package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type server struct{}

type messageData struct {
	Message string `json:"message"`
}

// 收到請求要執行的伺服器服務
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// msg := `{"message":"hello world"}`
	// w.Write([]byte(msg))
	messageData := messageData{}
	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(messageData)

	if r.Header.Get("X-Time") != "" {
		log.Println("X-Time is: ", r.Header.Get("X-Time"))
	}

	messageData.Message = strings.ToUpper(messageData.Message)
	jsonBytes, _ := json.Marshal(messageData)
	w.Write(jsonBytes)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
