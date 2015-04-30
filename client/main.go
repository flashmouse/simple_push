package main

import (
	"fmt"
	"net"
	"flag"
	"strconv"
	"time"
	tomb "gopkg.in/tomb.v1"
)

var (
	ip string
	port int
)

//only for test
func main() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "server listene ip")
	flag.IntVar(&port, "port", 9987, "server listen port")
	flag.Parse()
	loop := 15000
	errLoop := 0
	for i := 0; i < loop; i++ {
		con , err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
		con.SetDeadline(time.Now().Add(100 * time.Second))
		if err != nil {
			if errLoop%1000 == 0 {
				fmt.Printf("error: %v \n", err)
			}
			errLoop++
		}
		con.Write([]byte("hello world"))
		go func() {
			for {
				buf := make([]byte, 100, 100)
				_, err := con.Read(buf)
				if err != nil {
					return
				}
			}
		}()
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}


	//	time.Sleep(10 * time.Second)
	tomb1 := tomb.Tomb{}
	tomb1.Wait()
}

