/*
This package provides a PuppetDB Client in Golang.

An example for submitting a 'replace facts' command:

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

An example for submitting a 'deactivate node' command:

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
*/
package puppetdb
