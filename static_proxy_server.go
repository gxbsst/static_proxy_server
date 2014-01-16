package main

import (
	"fmt"
	"static_proxy_server/lib/test"
)

func init() {
	fmt.Printf("我是1\n")
}

func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func GreetNames(names []string, suffix string) {
	for _, n := range names {
		Greet(n + suffix)
	}
}

func main() {
	test.Puts()
	name := []string{
		"Weston",
		"Lisa",
		"CC",
		"Leo",
	}
	comm := make(chan string)

	go func() {
		GreetNames(name, "<CR>")
		comm <- "Finish Greeting name"
		// comm
	}()
	GreetNames(name, "CM")
	fmt.Println(<-comm)

}
