package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	server := puppetdb.NewServer("http://localhost:8080/")

	facts := map[string]string{
		"foo": "bar",
	}

	response := server.ReplaceFacts("foobar", facts)
	fmt.Printf("UUID: %v\n", response.Uuid)
}
