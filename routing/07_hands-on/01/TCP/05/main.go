package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			fmt.Println("This is the end")
			break
		}
	}
	fmt.Fprintln(conn, "Here we WRITE to the response.")
	io.WriteString(conn, "2Here we WRITE to the response.")
}
