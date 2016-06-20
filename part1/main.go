/*
   main.go

   This is the stub.

   This code is released to the public domain. Originally prepared for the LCA 2015 conference
   by Mark Smith <mark@qq.is>.
*/

package main

import (
	"bufio"
	"net"
	"net/http"
)

func main() {
	// 1. listen for connections forever
	if ln, err := net.Listen("tcp", ":8080"); err == nil {
		//f2. Accept Request
		if conn, err := ln.Accept(); err == nil {
			reader := bufio.NewReader(conn)

			if req, err := http.ReadRequest(reader); err == nil {
				if be, err := net.Dial("tcp", "127.0.0.1:8081"); err == nil {
					be_reader := bufio.NewReader(be)
				}
			}

		}

	}
}
