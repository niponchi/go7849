package main

import (
	"fmt"
	"github.com/niponchi/go7849/lib/proto"
	"github.com/rs/xid"
)

func main() {
	fmt.Println(xid.New().String())

	client := svc.NewChatClient(nil)
	fmt.Println(client)
}
