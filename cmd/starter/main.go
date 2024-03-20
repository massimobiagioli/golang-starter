package main

import (
	"fmt"

	"github.com/massimobiagioli/golang-starter/internal/message"
)

func main() {
	message := message.GetMessage()
	fmt.Println(message)
}