package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"imp"
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"os"
)


var  Poll []*imp.Wsmsg

func wsHandler(w http.ResponseWriter , r *http.Request){
	fmt.Println(r.RemoteAddr)
	up:=&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	var (
		wsconn *websocket.Conn
		err error

	)

	if wsconn,err= up.Upgrade(w,r,nil);err!=nil{
		return
	}
	ws:=imp.Initwsmsg(wsconn)
	Poll=append(Poll, ws)

	go ws.WriteToAll(&Poll)
	go ws.Read()

	return
}

func main2(){
	beego.Run()
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":5555",nil)
}



func main(){
	main2()
	//StartClient1()
}

func StartClient1() {
	tcpAddress, err := net.ResolveTCPAddr("tcp4", "192.168.100.145:51680")
	if err != nil {
		//errs.Error_exit(err)
		fmt.Fprintf(os.Stderr,"Fatal errs %s",err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		//errs.Error_exit(err)
		fmt.Fprintf(os.Stderr,"Fatal errs %s",err.Error())

	}

	writeChan := make(chan []byte, 1024)
	readChan := make(chan []byte, 1024)

	go writeConnection(conn, writeChan)
	go readConnection(conn, readChan)

	//go handleReadChannel(readChan)

	for {
		var s string
		fmt.Scan(&s)
		writeChan <- []byte(s)
	}

}

func readConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr,"Fatal errs2 %s",err.Error())
			//errs.Error_print(err)
			return
		}
		println("Received from:", conn.RemoteAddr().String(), string(buffer[:n]))
		//channel <- buffer[:n]
	}

}

func writeConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()
	for {
		select {
		case data := <- channel:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Fprintf(os.Stderr,"Fatal errs %s",err.Error())
				//errs.Error_exit(err)
			}
			println("Write to:", conn.RemoteAddr(), string(data))
		}
	}
}
