package main

import (
	"wander/DataProvider/connector"
	"wander/DataProvider/provider"
	"fmt"
	"log"
	"time"
)

const (
	AIRCON = 0        //air conditioner
	STOVE = 1        //hang stove
	FRIDGE = 2        //fridge
	RCKR = 3        //rice cooker
	//tet = 4
)

func main() {

	dat := provider.Hatcher(AIRCON)

	//dat := provider.Hatcher(RCKR)
	//datBase := base64.StdEncoding.EncodeToString(dat)
	//d := []byte("!!" + datBase + "\n")

	send(dat)
}

func send(dat []byte) {
	conn := connector.Connection()
	fmt.Println("connection established!")
	for i := 1; i < 50; i++ {
		log.Println(i)
		time.Sleep(5 * time.Second)
		conn.Write(dat)
	}

	defer conn.Close()
}
