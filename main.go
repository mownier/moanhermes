package main

import (
	"./server"
)

func main() {
	moanhermes := server.NewMoanhermes()
	moanhermes.StartServing("")
}