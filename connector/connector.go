package connector

import (
	"github.com/wanderself/DataProvider/handle"
	"net"
)

func Connection() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8811")
	handle.Err(err)
	return conn
}
