package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	// Create catalog
        catalog := puppetdb.NewCatalogWireFormat()
	catalog.Metadata.ApiVersion = 1
	catalog.Data.Name = "foobar"
	catalog.Data.Version = "3"
	catalog.Data.TransactionUuid = "aaaab"

	var e1 puppetdb.CatalogEdge
	e1.Source = puppetdb.CatalogResourceSpec{"File", "/etc"}
	e1.Target = puppetdb.CatalogResourceSpec{"File", "/etc/hosts"}
	e1.Relationship = "required-by"

	catalog.Data.Edges = []puppetdb.CatalogEdge{e1}

	var r1 puppetdb.CatalogResource
	r1.Type = "File"
	r1.Title = "/etc"
	r1.Exported = false
	r1.File = "/etc/puppet/manifests/site.pp"
	r1.Line = 1
	r1.Tags = []string{"foo", "bar"}
	r1.Parameters = map[string]string{"foo":"bar"}

	var r2 puppetdb.CatalogResource
	r2.Type = "File"
	r2.Title = "/etc/hosts"
	r2.Exported = false
	r2.File = "/etc/puppet/manifests/site.pp"
	r2.Line = 1
	r2.Tags = []string{"foo", "bar"}
	r2.Parameters = map[string]string{"foo":"bar"}

	catalog.Data.Resources = []puppetdb.CatalogResource{r1, r2}

	// Submit catalog
	server := puppetdb.NewServer("http://localhost:8080/")
	response, _ := server.ReplaceCatalog(catalog)
	fmt.Printf("UUID: %v\n", response.Uuid)
}
