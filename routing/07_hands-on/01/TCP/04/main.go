package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
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
	fmt.Printf("%t\n", conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("This is the end of the http request headers")
			break
		}
	}
}
