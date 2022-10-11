package main

import (
	"fmt"
	"log"
	"net" // standard network package
	"strings"
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

	var request = string(buf)
	var spitedRequest = strings.Split(request, " ")

	var method = spitedRequest[0]
	var path = spitedRequest[1]

	// write response
	const statusLine string = "HTTP/1.1 200 OK\n"
	const header string = "Content-Type: text/html; charset=utf-8\n"
	var body string
	switch method {
	case "GET":
		switch path {
		case "/":
			body = "default route"
		case "/favicon.ico":
			body = "favicon"
		case "/hello":
			body = "Hello, World!"
		case "/bye":
			body = "Good bye!"
		case "/hello.jp":
			body = "こんにちは！"
		default:
			body = "not found"
		}
	case "POST":
	}
	conn.Write([]byte(statusLine + header + "\n" + body))
}
