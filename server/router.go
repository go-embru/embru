package server

import (
	"embru/src/util"
	"fmt"
	"net"
)

type router struct {
	conn         net.Conn
	chanel       chan string
	pushing      bool
	handing      bool
	pushExecutor *PushExecutor
}

func newRouter(conn net.Conn, executor *PushExecutor) *router {
	v := &router{
		conn:         conn,
		chanel:       nil,
		pushing:      true,
		handing:      true,
		pushExecutor: executor,
	}

	v.initPusher()
	v.initHandler()
	return v
}

func (router *router) push2self(m string) {
	router.chanel <- m
}

func (router *router) autoPush(name string, msg string) {
	router.pushExecutor.push(name, msg)
}

func (router *router) initHandler() {
	if router.handing {
		for router.handing {
			func() {
				b := make([]byte, 1024)
				length, _ := router.conn.Read(b)
				s := string(b[0:length])
				fmt.Println(s)

				name := "127.0.0.1:" + s
				router.autoPush(name, s)
			}()
		}
	}
}

func (router *router) initPusher() {

	if router.pushing {
		router.chanel = make(chan string)

		go func() {
			for router.pushing {
				v := <-router.chanel
				router.conn.Write(util.Str2bytes(v))
			}
		}()
	}
}
