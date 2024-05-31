package main

import (
	"context"
	"github.com/cloudwego/netpoll"
	"time"
)

func main() {
	listener, err := netpoll.CreateListener("tcp", ":8080")
	if err != nil {
		panic("create netpoll listener failed")
	}
	//conn, err := listener.Accept()
	//if err != nil {
	//	panic("create netpoll listener failed")
	//}

	eventLoop, _ := netpoll.NewEventLoop(
		func(ctx context.Context, connection netpoll.Connection) error {
			return nil
		},
		//netpoll.WithOnPrepare(prepare),
		//netpoll.WithOnConnect(connect),
		netpoll.WithReadTimeout(time.Second),
	)
	eventLoop.Serve(listener)
}
