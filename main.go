package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	port = flag.Uint("p", 0, "http server port")
	dir  = flag.String("d", "", "http file directory")
)

func main() {
	flag.Parse()
	portS := func() string {
		if port == nil {
			return ":0"
		} else {
			return fmt.Sprintf(":%d", *port)
		}
	}()
	dirS := func() string {
		if dir == nil || *dir == "" {
			d, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			return d
		} else {
			return *dir
		}
	}()

	ln, err := net.Listen("tcp", portS)
	if err != nil {
		log.Fatal(err)
	}

	addr := ln.Addr().(*net.TCPAddr)

	fmt.Println(dirS, "    ", addr.String())
	showAddrs(addr.AddrPort().Port())

	err = http.Serve(ln, http.FileServer(http.Dir(dirS)))
	if err != nil {
		log.Fatal(err)
	}
}
func showAddrs(port uint16) {
	show := func(ip string) {
		fmt.Printf("http://%s:%d\n", ip, port)
	}

	show("0.0.0.0")

	a, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range a {
		ip := v.(*net.IPNet)
		if !ip.IP.IsGlobalUnicast() {
			continue
		}
		show(ip.IP.String())
	}
}
