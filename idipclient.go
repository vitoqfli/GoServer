package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	//"strconv"
)

var host = flag.String("host", "192.168.72.1", "host")
var port = flag.String("port", "80", "port")
var cmd = flag.String("cmd", "cmd=10451006&area=1000&partition=1&platid=1&openid=C6258EE95108132A9DAD4D421821145C&charac_no=1&source=11&serial=M-PAYO-20140414124009-58382166\n", "cmd")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
func handleWrite(conn net.Conn, done chan string) {
	for i := 1; i > 0; i-- {
		_, e := conn.Write([]byte(*cmd))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 2048)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"
}
