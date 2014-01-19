package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	server := puppetdb.NewServer("http://localhost:8080/")
	response := server.DeactivateNode("foobar")
	fmt.Printf("UUID: %v\n", response.Uuid)
}
