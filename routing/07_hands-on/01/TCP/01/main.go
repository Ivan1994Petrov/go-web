package main

import (
	"io"
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
		io.WriteString(conn, "I see you connected.\n\n")
		conn.Close()
	}
}
