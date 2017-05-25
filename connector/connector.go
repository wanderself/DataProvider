package connector

import (
	"net"
	//"github.com/golang/glog"
	"log"
)

func Connection() net.Conn {
	//conn, err := net.Dial("tcp", "10.5.19.147:6100")
	conn, err := net.Dial("tcp", "10.2.45.52:6100")
	//conn, err := net.Dial("tcp", "10.2.45.56:514")
	//conn, err := net.Dial("tcp", "172.16.16.166:6100")
	//handle.Err(err)

	if err != nil {
		log.Println(err.Error())
	}
	//glog.Infoln("connection established")
	return conn
}
