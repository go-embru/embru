package holder

import (
	"fmt"
	"github.com/go-embru/embru/util"
	"net"
)

var holder = make(map[string]net.Conn)

func init() {
	fmt.Println(holder)
	holder["127.0.0.1:12312"] = nil
}

func AddConn(name string, conn net.Conn) {
	if !util.IsBlank(name) {
		holder[name] = conn
	}
}
