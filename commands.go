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
func (server *Server) SubmitCommand(command string, version int, payload interface{}) CommandResponse {
	baseUrl := server.BaseUrl
	commandsUrl := strings.Join([]string{baseUrl, "v3/commands"}, "")

	commandObject := CommandObject{command, version, payload}
	commandJson, _ := json.Marshal(commandObject)

	data := url.Values{}
	data.Set("payload", string(commandJson[:]))

	req, _ := http.NewRequest("POST", commandsUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, _ := client.Do(req)

	bodyRC, _ := ioutil.ReadAll(resp.Body)

	var commandResponse CommandResponse
	json.Unmarshal(bodyRC, &commandResponse)

	return commandResponse
}

/*
Submit a new 'replace facts' command to PuppetDB.

This function will submit a 'replace facts' command. It accepts a certificate
name and a map of facts (key/value pairs).

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#replace-facts-version-1
*/
func (server *Server) ReplaceFacts(certname string, facts map[string]string) CommandResponse {
        factsPayload := FactsWireFormat{certname, facts}
	factsJson, _ := json.Marshal(factsPayload)

	commandResponse := server.SubmitCommand("replace facts", 1, string(factsJson[:]))
	return commandResponse
}

/*
Submit a new 'deactivate node' command to PuppetDB.

This function will submit a 'deactivate node' command. It accepts a certificate
name as an argument to indicate which node to deactivate.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#deactivate-node-version-1
*/
func (server *Server) DeactivateNode(certname string) CommandResponse {
	certnameJson, _ := json.Marshal(certname)

	commandResponse := server.SubmitCommand("deactivate node", 1, string(certnameJson[:]))
	return commandResponse
}

/*
Submit a new 'replace catalog' command to PuppetDB.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#replace-catalog-version-3
*/
func (server *Server) ReplaceCatalog(catalog CatalogWireFormat) CommandResponse {
	commandResponse := server.SubmitCommand("replace catalog", 3, catalog)
	return commandResponse
}
