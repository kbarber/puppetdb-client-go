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
