package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != http.StatusOK {
		log.Fatal(r.StatusCode)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}

	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}

func postDataAndReturnMessage(msg messageData) messageData {
	jsonByte, _ := json.Marshal(msg)    // 將要傳遞的結構變成JSON字串
	buffer := bytes.NewBuffer(jsonByte) // JSON 字串轉換成 byte.Buffer

	response, err := http.Post("http://localhost:8080", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	messageData := messageData{}
	err = json.NewDecoder(response.Body).Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	return messageData
}

func postDataAndReturnMessage2(msg messageData) messageData {
	jsonByte, _ := json.Marshal(msg)    // 將要傳遞的結構變成JSON字串
	buffer := bytes.NewBuffer(jsonByte) // JSON 字串轉換成 byte.Buffer

	request, err := http.NewRequest("POST", "http://localhost:8080", buffer)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("X-Time", "2025/1/1")
	request.Header.Set("Accept", "application/json")
	http.DefaultClient.Timeout = time.Second * 5
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	messageData := messageData{}
	err = json.NewDecoder(resp.Body).Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	return messageData
}

func main() {
	// response := getDataAndReturnResponse()
	// fmt.Println("response:", response)

	msg := messageData{Message: "Hi, server!"}
	data := postDataAndReturnMessage(msg)
	fmt.Println(data)
	data = postDataAndReturnMessage2(msg)
	fmt.Println("Second:", data)
}
