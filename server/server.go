package server

import (
	"embru/src/util"
	"embru/src/zaplog"
	"fmt"
	"net"
	"strings"
)

func Start() {

	m := make(map[string]net.Conn)

	l, err := net.Listen("tcp", ":3446")
	zaplog.CheckError(err)

	for {
		c, err := l.Accept()
		zaplog.CheckError(err)

		go handleConn(c, m)
	}
}

func handleConn(conn net.Conn, m map[string]net.Conn) {
	fmt.Println(conn.RemoteAddr().String())

	m[conn.RemoteAddr().String()] = conn

	for {
		b := make([]byte, 1024)
		length, _ := conn.Read(b)
		s := string(b[0:length])
		fmt.Println(s)

		arr := strings.Split(s, " ")
		cmd := arr[0]
		tgt := arr[1]
		msg := arr[2]

		if cmd == "1" {
			go func() {
				name := "127.0.0.1:" + tgt
				m[name].Write(util.Str2bytes(msg))
			}()
		}

		if cmd == "2" {
			go func() {
				for k, _ := range m {
					fmt.Println(k)
				}
			}()
		}

		conn.Write(util.Str2bytes("OK"))

	}
}
