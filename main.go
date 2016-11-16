package main

import (
	"github.com/wanderself/DataProvider/connector"
	//g "github.com/wanderself/DataProvider/handle"
	"github.com/wanderself/DataProvider/provider"
	"log"
	"time"
	"fmt"
)

const (
	AIRCON = 0        //air conditioner
	STOVE = 1        //hang stove
	FIRDGE = 2        //fridge
	RCKR = 3        //rice cooker
)

func main() {

	dat := provider.Hatcher(RCKR)
	log.Println(" 长度：", len(dat),  "\n二进制包： ", dat)
	send(dat)
	time.Sleep(5 * time.Second);


}

func send(dat []byte) {
	conn := connector.Connection()
	fmt.Println("connection established!")

	conn.Write(dat)

	defer conn.Close()
}
