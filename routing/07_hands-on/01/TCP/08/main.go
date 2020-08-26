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

	var i int
	var rMethon, rURI string
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			xs := strings.Fields(ln)
			rMethon = xs[0]
			rURI = xs[1]

			fmt.Println("METHON:", rMethon)
			fmt.Println("URI:", rURI)
		}

		if ln == "" {
			fmt.Println("This is the end")
			break
		}
		i++
	}

	body := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Code Gangsta</title>
			</head>
			<body>
				<h1>"THIS IS LOW LEVEL"</h1>
			</body>
			</html>
	`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length:", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}
