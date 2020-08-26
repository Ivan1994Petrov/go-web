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
	var rMethod, rURI string

	scanner := bufio.NewScanner(conn)

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

		fmt.Println("++++++++++++", i)
	}

	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(conn)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(conn)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleIndex(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET INDEX</title>
	</head>
	<body>
		<h1>"GET INDEX"</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
	</body>
	</html>
	`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length:", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}

func handleApply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET DOG</title>
	</head>
	<body>
		<h1>"GET APPLY"</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
		<form action="/apply" method="POST">
		<input type="hidden" value="In my good death">
		<input type="submit" value="submit">
		</form>
	</body>
	</html>
	`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length:", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>POST APPLY</title>
	</head>
	<body>
		<h1>"POST APPLY"</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
	</body>
	</html>
	`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length:", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}

func handleDefault(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>default</title>
	</head>
	<body>
		<h1>"default"</h1>
	</body>
	</html>
	`
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Content-Length:", len(body))
	fmt.Fprintln(conn, "Content-Type: text/html")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, body)
}
