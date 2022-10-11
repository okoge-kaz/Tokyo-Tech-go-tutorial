package main

import (
	"fmt"
	"log"
	"net" // standard network package
)

func main() {
	// config
	port := 8000
	protocol := "tcp"

	// resolve TCP address
	addr, err := net.ResolveTCPAddr(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}

	// get TCP socket
	socket, err := net.ListenTCP(protocol, addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listen: ", socket.Addr().String())

	// keep listening
	for {
		// wait for connection
		conn, err := socket.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Connected by ", conn.RemoteAddr().String())

		// yield connection to concurrent process
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// close connection when this function ends
	defer conn.Close()

	buf := make([]byte, 1024)
	conn.Read(buf)

	log.Printf("Request\n----------\n%s\n----------", string(buf))
	// write response
	const statusLine string = "HTTP/1.1 200 OK\n"
	const header string = "Content-Type: text/html; charset=utf-8\n"
	const body string = "Hello, World!"
	conn.Write([]byte(statusLine + header + "\n" + body))
}
