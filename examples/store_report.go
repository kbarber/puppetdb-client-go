package main

import (
	"fmt"
	"github.com/kbarber/puppetdb-client-go"
)

func main() {
	// Create report
	var report puppetdb.ReportWireFormat

	report.Certname = "foobar"
	report.PuppetVersion = "3.1.1"
	report.ReportFormat = 1
	report.ConfigurationVersion = "aaa"
	report.StartTime = "2013-10-28T12:35:00.000Z"
	report.EndTime = "2013-10-28T12:35:00.000Z"
	report.TransactionUuid = "aaa"

	var e1 puppetdb.ResourceEvent
	e1.ResourceType = "File"
	e1.ResourceTitle = "/etc/hosts"
	e1.Property = "content"
	e1.Timestamp = "2013-10-28T12:35:00.000Z"
	e1.Status = "changed"
	e1.OldValue = "foo"
	e1.NewValue = "bar"
	e1.Message = "foo has now changed to bar"
	e1.File = "/etc/puppet/manifests/site.pp"
	e1.Line = 1
	e1.ContainmentPath = []string{"main"}

	report.ResourceEvents = []puppetdb.ResourceEvent{e1}

	// Store report
	server := puppetdb.NewServer("http://localhost:8080/")
	response, _ := server.StoreReport(report)
	fmt.Printf("UUID: %v\n", response.Uuid)
}
