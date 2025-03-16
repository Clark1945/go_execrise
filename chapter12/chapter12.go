package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "your name")
	flag.Parse() // 解析使用者的輸入旗標
	if *n == "" {
		fmt.Println("Name is require")
		flag.PrintDefaults()
		os.Exit(1)
	}
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("嗨\n"))
	f.WriteString("我今天10月就準備離職")
	err = os.WriteFile("test2.txt", []byte("嗨\n"), 0673)
	if err != nil {
		panic(err)
	}
}

func checkFile(filename string) (os.FileInfo, error) {
	f, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%v: 檔案不存在\n", filename)
			return nil, err
		}
		return nil, err
	}
	return f, nil
}
