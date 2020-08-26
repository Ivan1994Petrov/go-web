package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}

		if ln == "" {
			fmt.Println("This is the end")
			break
		}
		i++
	}
	body := "RESPONSE BODY"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	body += "\n"

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintln(conn, "Content-Type: text/plain")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}
