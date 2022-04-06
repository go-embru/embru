package server

import "net"

type processor interface {
	//read(conn net.Conn) []byte
	//decode(message []byte) string
	//authorize(message string) bool
	handle(conn net.Conn) string
	//encode(message string) []byte
	//write(conn net.Conn, bs []byte)
}
