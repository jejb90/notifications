package main

import (
	"fmt"
	"notifications/internal/di"
)

func main() {

	handler, err := di.Initialize()
	if err != nil {
		panic("fatal err: " + err.Error())
	}

	err = handler.Send("news", "user", "news 1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("news", "user", "news 2")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("news", "user", "news 3")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("news", "another user", "news 1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.Send("status", "user", "status 1")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
