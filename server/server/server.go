package server

import (
	"net"
	"fmt"
	"sync"
	"time"
	"github.com/flashmouse/go-channel-map/chmap"
)

type Server struct{
	ip       string
	port     string
	listener net.Listener
	clients  *chmap.Chmap
	running  bool
	counter  int64
	lock     sync.Mutex
}

func NewServer(ip string, port string) Server {
	s := Server{}
	s.ip = ip
	s.port = port
	s.counter = 0
	s.clients = chmap.NewMap(32)//map[string]net.Conn
	return s
}

func (s *Server) count() {
	for {
		fmt.Printf("now has clients %d %d \n", s.counter, s.clients.Count())
		time.Sleep(3 * time.Second)
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.ip+":"+s.port)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	s.listener = listener
	go s.accept()
	go s.count()
}

func (s *Server) addCount() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.counter ++
}

func (s *Server) minusCount() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.counter --
}

func (s *Server) accept() {
	for {
		client, err := s.listener.Accept()

		if err != nil {
			fmt.Printf("error: %v", err)
		}

		client.SetDeadline(time.Now().Add(100 * time.Second))
		s.clients.Put(client.RemoteAddr().String(), client)
		s.addCount()

		go func() {
			buf := make([]byte, 100, 100)
			for {
				_, err1 := client.Read(buf)
				if err1 != nil {
					key := client.RemoteAddr().String()
					err2 := client.Close()
					s.clients.Delete(key)
					s.minusCount()
					if err2 != nil {
					}
					return
				}
			}
		}()
	}
}
