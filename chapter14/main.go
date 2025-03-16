package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hi</h1>"
	w.Write([]byte(msg))
}

func servePage1(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello</h1>"
	w.Write([]byte(msg))
}

func helloMsg(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	name, OK := v1["name"]
	if !OK {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<h1>MissingArgument</h1>"))
		return
	}
	w.Write([]byte(fmt.Sprintln("PASS!", name)))
}

// HTML模板
var templateStr = ` 
<html>
<head>
</head>
<body>
<h1>Hello</h1>
<p>Good to see ya {{.Name}}!!!!!</p>ともに行こう
</body>
</html>`

type Customer struct {
	ID   int
	Name string
}

func helloMsg2(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	name, OK := v1["name"]
	if !OK {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<h1>MissingArgument</h1>"))
		return
	}
	tmpl, _ := template.New("Execrise").Parse(templateStr)
	customer := Customer{Name: name[0]}
	tmpl.Execute(w, customer)
}

func main() {
	//
	// log.Fatal(http.ListenAndServe(":8080", &hello{})) // 建立一個 8080的服務，任何路徑任何HTTP方法皆回傳Hi
	//

	// http.HandleFunc("/page1", servePage1) // 設定的是http套件的DefaultServeMux結構，所以沒有改變
	// log.Fatal(http.ListenAndServe(":8080", hello{}))

	// http.HandleFunc("/page1", servePage1) // 設定的是http套件的DefaultServeMux結構，DefaultServeMux是http套件預設的ServeMux結構，與hello{}相同
	// http.Handle("/", hello{})
	// log.Fatal(http.ListenAndServe(":8080", nil)) // 使用預設的DefaultServerMux

	// mux := http.NewServeMux() // 自訂ServeMux結構
	// mux.HandleFunc("/page1", servePage1)
	// mux.HandleFunc("/passInfo", helloMsg)
	// mux.HandleFunc("/passInfo2", helloMsg2)
	// mux.HandleFunc("/htmlTemplate", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "chapter14/index.html") // mux本身並沒有serveFile()函式，必須使用http.ServeFile方法
	// })
	// mux.Handle("/", hello{})
	// log.Fatal(http.ListenAndServe(":8080", mux)) // 使用預設的DefaultServerMux << 會阻塞Goroutine導致8081不會啟動

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("chapter14/public"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
