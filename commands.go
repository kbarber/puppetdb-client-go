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
Top level struct representing a PuppetDB commands payload object.

It consists of:

* command
* version
* payload

Where payload contains a variable structure depending on the command itself.

See here for more details on the protocol: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
type CommandObject struct {
	Command string `json:"command"`
	Version	int `json:"version"`
	Payload	interface{} `json:"payload"`
}

/*
Response to a commands submission request.

This struct contains the fields that are returned when a command was
successfully submitted. This does not indicate the command was processed,
just an acknowledgement it was received and will be processed in the future.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#command-submission
*/
type CommandResponse struct {
	Uuid string `json:"uuid"`
}

/*
Fact command submission struct for submitting the 'replace facts' command
to PuppetDB.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/facts_format.html
*/
type FactsWireFormat struct {
	Name string `json:"name"`
	Values map[string]string `json:"values"`
}

/*
Representation of a PuppetDB server instance.

Use NewServer to create a new instance.
*/
type Server struct {
	BaseUrl string
}

/*
Create a new instance of a Server for usage later.

This is usually the main entry point of this SDK, where you would create
this initial object and use it to perform activities on the instance in
question.
*/
func NewServer(baseUrl string) Server {
	return Server{baseUrl}
}

/*
Generic command submission support, for submitting commands to a PuppetDB instance.

This is ordinarily not used, instead its recommended to use the various direct
functions instead.

More detail here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
func (server *Server) SubmitCommand(command string, version int, payload string) CommandResponse {
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
