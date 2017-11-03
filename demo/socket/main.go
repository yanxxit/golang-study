package main

import (
	"fmt"
	"io"
	"bufio"
	"net/http"
	"code.google.com/p/go.net/websocket"
	"container/list"
)

var connid int
var conns *list.List

func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()

	connid++
	id := connid

	fmt.Printf("connection id: %d\n", id)

	item := conns.PushBack(ws)
	defer conns.Remove(item)

	name := fmt.Sprintf("user%d", id)

	SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))

	r := bufio.NewReader(ws)

	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Printf("disconnected id: %d\n", id)
			SendMessage(item, fmt.Sprintf("%s offline\n", name))
			break
		}

		fmt.Printf("%s: %s", name, data)

		SendMessage(item, fmt.Sprintf("%s\t> %s", name, data))
	}
}

func SendMessage(self *list.Element, data string) {
	// for _, item := range conns {
	for item := conns.Front(); item != nil; item = item.Next() {
		ws, ok := item.Value.(*websocket.Conn)
		if !ok {
			panic("item not *websocket.Conn")
		}

		if item == self {
			continue
		}

		io.WriteString(ws, data)
	}
}


//该代码片段来自于: http://www.sharejs.com/codes/go/4374