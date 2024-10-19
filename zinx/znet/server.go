package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

// implement iServer interface, define a Server class
type Server struct {
	//server name
	Name string
	//tcp4 or other
	IPVersion string
	//bind IP address
	IP string
	//服务绑定端口
	Port int
}

// turn on network service
func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	//open a go to doing server-side Listener business
	go func() {
		//1 acquire a tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("resolve tcp addr err: %v\n", err)
			return
		}
		//2 Monitor server address
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server", s.Name, "succ, now listening.....")
		//3 Start the server network connection service
		for {
			//3.1 Block waiting for client to establish connection request
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			//3.2 Temporarily providing a maximum 512 byte echo service
			go func() {
				buf := make([]byte, 512)
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}
					//Echo back the data received to the client
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name ", s.Name)
	//TODO: Server.Stop() cleanup
}

func (s *Server) Serve() {
	s.Start()
	//Block the server to prevent it from exiting
	for {
		time.Sleep(10 * time.Second)
	}
}

// create a server handle
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777, // Corrected to integer
	}
	return s
}
