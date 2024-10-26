package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    req, err := http.ReadRequest(bufio.NewReader(conn))
    if err != nil {
        log.Println("Error reading request:", err)
        return
    }
    log.Println("request made!")
    if req.URL.Scheme == "" {
        req.URL.Scheme = "http"
    }
    if req.URL.Host == "" {
        req.URL.Host = req.Header.Get("Host")
    }
    resp, err := http.DefaultTransport.RoundTrip(req)
    if err != nil {
        log.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    conn.Write([]byte("HTTP/1.1 " + resp.Status + "\r\n"))
    for key, values := range resp.Header {
        for _, value := range values {
            conn.Write([]byte(key + ": " + value + "\r\n"))
        }
    }
    conn.Write([]byte("\r\n"))
    io.Copy(conn, resp.Body)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	log.Println("Proxy server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
