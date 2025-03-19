package main

import (
	"fmt"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	if wg != nil {
		wg.Done()
	}
}

func main() {
	var s1, s2 int

	wg := &sync.WaitGroup{} //多執行緒結構
	wg.Add(1)               //有一項工作

	go sum(1, 100, wg, &s1) //以Goroutine形式執行的sum()
	sum(1, 100, nil, &s2)   // 正常執行的sum()
	wg.Wait()               // 等待一個Goroutine執行完成

	fmt.Println(s1, s2)
}
