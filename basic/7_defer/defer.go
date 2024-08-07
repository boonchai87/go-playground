package main

import "fmt"

func main() {
	// ทำงานเหมือน stack เข้าก่อนออกทีหลัง
	// statements are executed in LIFO (last-in, first-out) order when the function returns.
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("hello")
}
