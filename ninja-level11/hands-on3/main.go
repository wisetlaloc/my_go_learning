package main

import (
	"fmt"
	"log"
)

type customErr struct {
	err string
}

func (c customErr) Error() string {
	return c.err
}

func foo(e error) {
	log.Println("fooified error: ", e)
}

func main() {
	c := customErr{
		err: "Custom Error",
	}
	fmt.Println(c.err)
	fmt.Printf("%T\n", c)
	foo(c)
}
