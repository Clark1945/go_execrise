package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var helloList = []string{
	"Hello",
	"哈囉",
	"こんにちは",
	"Hi",
	"Bonjour",
}

func main() {
	fmt.Println("Hello Golang")
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err) // 強制結束程式
	}
	fmt.Println(msg)
	declareVariable()
	calculation()
	stringTest()
	if err := testExcpetion(); err != nil {
		fmt.Print(err)
	}
}

func execrise2() {
	rand.Seed(time.Now().Local().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Print(stars)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range" + strconv.Itoa(index))
	}
	return helloList[index], nil
}

func variableTest() {
	var int1 int = 0
	var str1 string = ""
	fmt.Println(int1)
	fmt.Println(str1)
}

func testSeed() {
	var seed = 123456789
	rand.Seed(int64(seed))
}

func test1_2_5() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}

func declareVariable() {
	var foo string = "string"
	boo := "string"
	var goo = "g"

	var ( // 一次定義多個
		demo1 string = "demo1"
		demo2 string
	)
	demo3, demo4, demo5 := "demo3", "demo4", "demo5"

	demo3, demo4, demo5 = "Demo3", "Demo4", "Demo5"
	fmt.Println(foo, boo, goo)
	fmt.Println(demo1, demo2)
	fmt.Println(demo3, demo4, demo5)
}

func calculation() {
	var a int = 10
	b := a / 3
	fmt.Println(b)
}

func stringTest() {
	concateName := "ABC" + " " + "DEF"
	fmt.Println(concateName)

	var discount float64
	var emails []string
	fmt.Printf("Discount : %#v \n Extra: %#v \n Type: %#v ", discount, emails, discount)

	fmt.Println("\n##########################")

	var nilString1 *string
	fmt.Println(nilString1)

	var count1 *int    //宣告指標變數(nil)
	count2 := new(int) //宣告指標變數(0)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v %d\n ", count2, *count2) // 宣告指標後用*取得指標的值
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)

}
func constTest() {
	const str string = ""

}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

func testExcpetion() error {
	if 3 > 2 {
		return errors.New("必定錯誤")
	} else {
		return nil
	}
}
