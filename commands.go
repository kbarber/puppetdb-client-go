package puppetdb

import (
	"encoding/json"
	"net/http"
	"net/url"
	"bytes"
	"io/ioutil"
	"strings"
)

/*
Generic command submission support, for submitting commands to a PuppetDB instance.

This is ordinarily not used, instead its recommended to use the various direct
functions instead.

More detail here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
func (server *Server) SubmitCommand(command string, version int, payload interface{}) (*CommandResponse, error) {
	baseUrl := server.BaseUrl
	commandsUrl := strings.Join([]string{baseUrl, "v3/commands"}, "")

	commandObject := CommandObject{command, version, payload}
	commandJson, err := json.Marshal(commandObject)
	if(err != nil) {
		return nil, err
	}

	data := url.Values{}
	data.Set("payload", string(commandJson[:]))

	req, err := http.NewRequest("POST", commandsUrl, bytes.NewBufferString(data.Encode()))
	if(err != nil) {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, err := client.Do(req)
	if(err != nil) {
		return nil, err
	}

	bodyRC, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		return nil, err
	}

	var commandResponse CommandResponse
	err = json.Unmarshal(bodyRC, &commandResponse)
	if(err != nil) {
		return nil, err
	}

	return &commandResponse, nil
}

/*
Submit a new 'replace facts' command to PuppetDB.

This function will submit a 'replace facts' command. It accepts a certificate
name and a map of facts (key/value pairs).

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#replace-facts-version-1
*/
func (server *Server) ReplaceFacts(certname string, facts map[string]string) (*CommandResponse, error) {
        factsPayload := FactsWireFormat{certname, facts}
	factsJson, err := json.Marshal(factsPayload)
	if(err != nil) {
		return nil, err
	}

	commandResponse, err := server.SubmitCommand("replace facts", 1, string(factsJson[:]))
	return commandResponse, err
}

/*
Submit a new 'deactivate node' command to PuppetDB.

This function will submit a 'deactivate node' command. It accepts a certificate
name as an argument to indicate which node to deactivate.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#deactivate-node-version-1
*/
func (server *Server) DeactivateNode(certname string) (*CommandResponse, error) {
	certnameJson, err := json.Marshal(certname)
	if(err != nil) {
		return nil, err
	}

	commandResponse, err := server.SubmitCommand("deactivate node", 1, string(certnameJson[:]))
	return commandResponse, err
}

/*
Submit a new 'replace catalog' command to PuppetDB.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#replace-catalog-version-3
*/
func (server *Server) ReplaceCatalog(catalog CatalogWireFormat) (*CommandResponse, error) {
	commandResponse, error := server.SubmitCommand("replace catalog", 3, catalog)
	return commandResponse, error
}

/*
Submit a new 'store report' command to PuppetDB.

More details here: http://docs.puppetlabs.com/puppetdb/1.6/api/commands.html#store-report-version-2
*/
func (server *Server) StoreReport(report ReportWireFormat) (*CommandResponse, error) {
	commandResponse, error := server.SubmitCommand("store report", 2, report)
	return commandResponse, error
}
