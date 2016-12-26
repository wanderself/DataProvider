package connector

import (
	"github.com/wanderself/DataProvider/handle"
	"net"
	//"github.com/golang/glog"
)

func Connection() net.Conn {
	conn, err := net.Dial("tcp", "10.5.19.147:6100")
	//conn, err := net.Dial("tcp", "10.2.45.52:6100")
	//conn, err := net.Dial("tcp", "172.16.16.166:6100")
	handle.Err(err)
	//glog.Infoln("connection established")
	return conn
}
