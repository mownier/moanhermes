package main

import (
	"./server"
)

func main() {
	m := server.NewMoanhermes()
	m.StartServing(":8080")
}