package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	server := puppetdb.NewServer("http://localhost:8080/")
	response, _ := server.QueryVersion()
	fmt.Printf("Version: %v\n", response.Version)
}
