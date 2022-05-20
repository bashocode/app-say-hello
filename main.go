package main

import (
	"fmt"

	go_say_hello "github.com/bashocode/go-say-hello"
)

func main() {
	name := "Karin"
	fmt.Println(go_say_hello.SayHello(name))
}
