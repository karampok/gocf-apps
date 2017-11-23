package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"os"
)

var (
	buildstamp string
	githash    string
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	port := os.Getenv("PORT")
	log.Printf("Binary version [%s] built at [%s]", githash, buildstamp)
	log.Printf("Listening tcp at port %s", port)
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("couldn't start listening:", err)
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				log.Printf("couldn't accept: " + err.Error())
				continue
			}
			log.Printf("Connection: (Local) %v <-> %v (Remote)\n",
				client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	client.Write([]byte("How can I help you\n>"))
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(bytes.ToUpper(line))
		client.Write([]byte("\n>"))
	}
}
