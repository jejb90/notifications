package repositories

import "fmt"

type Gateway struct{}

func (g *Gateway) Send(userID, message string) {
	fmt.Printf("Sending message to user %s: %s\n", userID, message)
}

func NewGatewayRepository() *Gateway {
	return &Gateway{}
}
