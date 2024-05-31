package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var a *int

	c := &a

	fmt.Println(c, &a, &c, c == &a)
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Close()
	fmt.Println(reply, err)

}
