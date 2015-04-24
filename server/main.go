package main

import (
	"github.com/flashmouse/simple_push/server/server"
	tomb "gopkg.in/tomb.v1"
	"flag"
	"strconv"
	"runtime"
)

var (
	ip string
	port int
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
//	runtime.GOMAXPROCS(1)
	flag.StringVar(&ip, "ip", "127.0.0.1", "server listene ip")
	flag.IntVar(&port, "port", 9987, "server listen port")
	flag.Parse()
	s := server.NewServer(ip, strconv.Itoa(port))
	s.Start()
	tomb1 := tomb.Tomb{}
	tomb1.Wait()
}

