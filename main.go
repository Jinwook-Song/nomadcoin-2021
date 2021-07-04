package main

import "fmt"


func main() {
	a := 2
	b := &a  // a의 메모리 주소
	a = 30
	fmt.Println(&a, *b)
}