package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// Simulate client
func ClientTest() {
	fmt.Println("Client Test ... start")
	//after 3 steps open a test request,
	//provide a opportunity for server to start service
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("write buf error ")
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}
		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

// Test functions for the Server module
func TestServer(t *testing.T) {
	/*
		Server testing
	*/
	//1. Create a server handle s
	s := NewServer("[zinx V0.1]")
	/*
		Client testing
	*/
	go ClientTest()
	//2 starting service
	s.Serve()
}
