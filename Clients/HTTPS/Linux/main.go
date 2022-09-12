package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/d-demirci/Project-Whis/Clients/HTTPS/Linux/core"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	fmt.Println("Linux Client")
	core.Boot()
	for {
		time.Sleep(15 * time.Millisecond)
	}
}
