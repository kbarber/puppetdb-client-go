package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	server := puppetdb.NewServer("http://localhost:8080/")
	response, _ := server.QueryServerTime()
	fmt.Printf("Server Time: %v\n", response.ServerTime)
}
