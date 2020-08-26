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
		log.Fatal(err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("This is the end!")
				break
			}
		}

		fmt.Println("code got here")
		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
