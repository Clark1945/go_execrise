package main

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
)

func main() {
	fmt.Println("Chapter 3")
	memoryTest()
	testCollection()

	l1, l2, l3 := linked()
	fmt.Println("有連結:", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("沒連結:", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("有設容量有連結:", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("有設容量無連結:", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Println("使用Copy()無連結:", copy1, copy2)
	fmt.Printf("使用%v個元素", copied)
	a1, a2 := appendNoLink()
	fmt.Println("使用apend()無連結", a1, a2)

	users := map[string]string{"305": "Bob"}
	users["306"] = "Daisy"
	user, exist := users["307"]
	fmt.Println(users)
	fmt.Println(user, exist)
	delete(users, "306")
	fmt.Println(users)

	b := &point{}
	b.setPoint(10, 5)
	fmt.Println("B=", b)
}

func memoryTest() {
	var list []int
	for i := 0; i < 40000000; i++ {
		list = append(list, 100)
	}

	// 印出Heap記憶體用量
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)

	// 測試overflow
	var a int64 = math.MaxInt64
	bigInt8 := big.NewInt(a)
	fmt.Println(bigInt8)
	fmt.Println(bigInt8.Add(bigInt8, big.NewInt(1)))
}

func testCollection() {
	arr1 := [5]int{1, 1} // 建立長度為5的陣列並定義初始值
	fmt.Println(arr1)
	arr2 := [...]int{5: 8, 9: 7} // 對陣列索引9賦予0值，因此長度為10
	fmt.Println(arr2)

	var slice []string
	slice = append(slice, "A")
	slice = append(slice, "B", "C")
	fmt.Println(slice)
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1    // 直接複製切片
	s3 := s1[:] // 從切片複製同長度切片
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6) // 加入新值，超過s1容量，觸發置換陣列
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10, 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1) // 回傳成功複製數量
	s1[3] = 99

	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

type id string // 簡易自訂型別

func getIDS() (id, id, id) {
	var id1 id
	var id2 id = "1234"
	var id3 id
	id3 = "12345678"
	return id1, id2, id3
}

type student struct { // Go不是物件導向語言
	id   int
	name string
}

func createStudent() student {
	var student student = student{
		id:   1,
		name: "Clark",
	}
	return student
}

func createStudent2() bool {
	student := struct {
		id   int
		name string
	}{
		2,
		"Clark",
	}
	fmt.Print(student.name)
	return true
}

type name string

func (n name) printName() {
	fmt.Println("name:", n)
}

type point struct {
	x int
	y int
}

func (p *point) setPoint(x, y int) {
	p.x = x
	p.y = y
}

func doubler(v interface{}) (string, error) {
	switch s := v.(type) {
	case string:
		return "string" + s, nil
	case bool:
		return "true", nil
	}
}
