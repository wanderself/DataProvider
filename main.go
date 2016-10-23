package main

import (
	"github.com/wanderself/DataProvider/connector"
	"github.com/wanderself/DataProvider/handle"
	"github.com/wanderself/DataProvider/provider"
)

func main() {
	dat := provider.Hatcher()
	handle.Log(dat)

	send(dat)
}

func send(dat []byte) {
	conn := connector.Connection()
	handle.Log("connection established!")

	conn.Write(dat)

	defer conn.Close()
}
