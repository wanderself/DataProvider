package main

import (
	"github.com/wanderself/DataProvider/connector"
	//g "github.com/wanderself/DataProvider/handle"
	"github.com/wanderself/DataProvider/provider"
	//"encoding/base64"
	"fmt"

	"log"
	"time"
)

const (
	AIRCON = 0        //air conditioner
	STOVE = 1        //hang stove
	FRIDGE = 2        //fridge
	RCKR = 3        //rice cooker
	tet = 4
)

func main() {

	//dat := provider.Hatcher(RCKR)
	//log.Println(" 长度：", len(dat),  "\n二进制包： ", dat)
	//send(dat)
	//for i := 0; i < 5; i++ {

	//datSTV := provider.Hatcher(STOVE)
	//log.Println(" 长度：", len(datSTV), "\n二进制包： ", datSTV)
	//send(datSTV)
	//time.Sleep(5 * time.Second)
	//
	dat := provider.Hatcher(tet)
	//send(datAC)
	//time.Sleep(5 * time.Second)
	//
	//datFRI := provider.Hatcher(FRIDGE)
	//send(datFRI)
	//time.Sleep(5 * time.Second)

	//dat := provider.Hatcher(RCKR)
	//datBase := base64.StdEncoding.EncodeToString(dat)
	//d := []byte("!!" + datBase + "\n")



	send(dat)

	//}

}

func send(dat []byte) {
	conn := connector.Connection()
	fmt.Println("connection established!")
	for i := 1; i < 5000000; i++ {
		log.Println(i)
		time.Sleep(0 * time.Second)
		conn.Write(dat)
	}

	defer conn.Close()
}
