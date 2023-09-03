package main

import (
	model "github.com/himanshu1221/GoShort/Model"
	"github.com/himanshu1221/GoShort/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
