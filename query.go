package puppetdb

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

/*
Generic query function.
*/
func (server *Server) Query(url string) ([]byte, error) {
	baseUrl := server.BaseUrl

	fullUrl := strings.Join([]string{baseUrl, url}, "")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if(err != nil) {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if(err != nil) {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

/*
Query the PuppetDB instance version end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/version.html
*/
func (server *Server) QueryVersion() (*Version, error) {
	body, err := server.Query("v3/version")
	if(err != nil) {
		return nil, err
	}

	var version Version
	json.Unmarshal(body, &version)

	return &version, err
}

/*
Query the PuppetDB instance server-time end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/server-time.html
*/
func (server *Server) QueryServerTime() (*ServerTime, error) {
	body, err := server.Query("v3/server-time")
	if(err != nil) {
		return nil, err
	}

	var serverTime ServerTime
	json.Unmarshal(body, &serverTime)

	return &serverTime, err
}

/*
Query the PuppetDB instance fact-names end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/fact-names.html
*/
func (server *Server) QueryFactNames() ([]string, error) {
	body, err := server.Query("v3/fact-names")
	if(err != nil) {
		return nil, err
	}

	var factNames []string
	json.Unmarshal(body, &factNames)

	return factNames, err
}

/*
Query the PuppetDB instance catalogs end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/catalogs.html
*/
func (server *Server) QueryCatalogs(certname string) (*CatalogWireFormat, error) {
	url := fmt.Sprintf("v3/catalogs/%v", certname)
	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var catalog CatalogWireFormat
	json.Unmarshal(body, &catalog)

	return &catalog, err
}

/*
Query the PuppetDB instance facts end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/facts.html#get-v3facts
*/
func (server *Server) QueryFacts(queryString string) (*[]Fact, error) {
	url := fmt.Sprintf("v3/facts?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var facts []Fact
	json.Unmarshal(body, &facts)

	return &facts, err
}

/*
Query the PuppetDB instance facts end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/facts.html#get-v3factsname
*/
func (server *Server) QueryFactsByName(name string, queryString string) (*[]Fact, error) {
	url := fmt.Sprintf("v3/facts/%v?%v", name, queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var facts []Fact
	json.Unmarshal(body, &facts)

	return &facts, err
}

/*
Query the PuppetDB instance facts end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/facts.html#get-v3factsnamevalue
*/
func (server *Server) QueryFactsByNameValue(name string, value string, queryString string) (*[]Fact, error) {
	url := fmt.Sprintf("v3/facts/%v/%v?%v", name, value, queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var facts []Fact
	json.Unmarshal(body, &facts)

	return &facts, err
}

/*
Query the PuppetDB instance resources end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/resources.html#get-v3resources
*/
func (server *Server) QueryResources(queryString string) (*[]CatalogResource, error) {
	url := fmt.Sprintf("v3/resources?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var resources []CatalogResource
	json.Unmarshal(body, &resources)

	return &resources, err
}

/*
Query the PuppetDB instance nodes end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/nodes.html#get-v3nodes
*/
func (server *Server) QueryNodes(queryString string) (*[]Node, error) {
	url := fmt.Sprintf("v3/nodes?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var nodes []Node
	json.Unmarshal(body, &nodes)

	return &nodes, err
}

/*
Query the PuppetDB instance reports end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/reports.html#get-v3reports
*/
func (server *Server) QueryReports(queryString string) (*[]Report, error) {
	url := fmt.Sprintf("v3/reports?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var reports []Report
	json.Unmarshal(body, &reports)

	return &reports, err
}

/*
Query the PuppetDB instance events end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/events.html#get-v3events
*/
func (server *Server) QueryEvents(queryString string) (*[]Event, error) {
	url := fmt.Sprintf("v3/events?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var event []Event
	json.Unmarshal(body, &event)

	return &event, err
}

/*
Query the PuppetDB instance event-counts end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/event-counts.html#get-v3event-counts
*/
func (server *Server) QueryEventCounts(queryString string) (*EventCounts, error) {
	url := fmt.Sprintf("v3/event-counts?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var eventCounts EventCounts
	json.Unmarshal(body, &eventCounts)

	return &eventCounts, err
}

/*
Query the PuppetDB instance aggregate-event-counts end-point.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/aggregate-event-counts.html#get-v3aggregate-event-counts
*/
func (server *Server) QueryAggregateEventCounts(queryString string) (*AggregateEventCounts, error) {
	url := fmt.Sprintf("v3/aggregate-event-counts?%v", queryString)

	body, err := server.Query(url)
	if(err != nil) {
		return nil, err
	}

	var aec AggregateEventCounts
	json.Unmarshal(body, &aec)

	return &aec, err
}
