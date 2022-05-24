package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "26379"
	CONN_TYPE = "tcp"
)

func main() {
	log.Println("starting sentinel")
	l, err := net.Listen(CONN_TYPE, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	if err != nil {
		log.Fatal("error listening")
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("error accepting")
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	s := bufio.NewScanner(conn)

	for s.Scan() {
		data := s.Text()
		log.Println(data)
		if !strings.HasPrefix(data, "*") && !strings.HasPrefix(data, "$") {
			handleCommand(data, conn)
		}
	}
}

func handleCommand(inp string, conn net.Conn) {
	command := strings.ToLower(inp)

	result, err := ioutil.ReadFile(command)
	if err == nil {
		conn.Write([]byte(result))
	} else {
		conn.Write([]byte(""))
	}
}
