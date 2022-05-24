package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	CONN_HOST = "0.0.0.0"
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

		if strings.HasPrefix(data, "*") {
			numberOfParameters, _ := strconv.Atoi(data[1:])
			parameters := []string{}
			for s.Scan() {
				parameter := s.Text()
				if !strings.HasPrefix(parameter, "$") {
					parameters = append(parameters, parameter)
				}
				if len(parameters) == numberOfParameters {
					break
				}
			}
			handleCommand(strings.Join(parameters, "-"), conn)
		}
	}
}

func handleCommand(inp string, conn net.Conn) {
	command := strings.ToLower(inp)
	log.Printf("> %s", command)

	result, err := ioutil.ReadFile(command)
	if err == nil {
		result := []byte(result)
		conn.Write(result)
	}
}
