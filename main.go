package main

import (
	"fmt"
	"notifications/internal/di"
	"time"
)

func main() {

	handler, err := di.Initialize()
	err = handler.Send("status", "user", "news 1")
	if err != nil {
		fmt.Println("Error:", err)
	}
	time.Sleep(5 * time.Second)

	err = handler.Send("status", "user", "news 2")
	if err != nil {
		fmt.Println("Error:", err)
	}
	time.Sleep(6 * time.Second)

	err = handler.Send("status", "user", "news 3")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = handler.Send("status", "user", "news 4")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = handler.Send("news", "user", "news 5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("status", "user", "news 6")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = handler.Send("status", "user", "news 7")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = handler.Send("news", "another user", "news 8")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("news", "user", "status 9")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
