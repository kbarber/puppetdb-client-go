package puppetdb

import (
	"encoding/json"
	"net/http"
	"net/url"
	"bytes"
	"io/ioutil"
	"strings"
)

type PayloadObject struct {
	Command string `json:"command"`
	Version	int `json:"version"`
	Payload	interface{} `json:"payload"`
}

type CommandObject struct {
	Payload PayloadObject `json:"payload"`
}

type CommandResponse struct {
	Uuid string `json:"uuid"`
}

type FactsWireFormat struct {
	Name string `json:"name"`
	Values map[string]string `json:"values"`
}

type Server struct {
	BaseUrl string
}

func NewServer(baseUrl string) Server {
	return Server{baseUrl}
}

func (server *Server) SubmitCommand(command string, version int, payload string) CommandResponse {
	baseUrl := server.BaseUrl
	commandsUrl := strings.Join([]string{baseUrl, "v3/commands"}, "")

	payloadObject := PayloadObject{command, version, payload}
	postPayload, _ := json.Marshal(payloadObject)

	data := url.Values{}
	data.Set("payload", string(postPayload[:]))

	req, _ := http.NewRequest("POST", commandsUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, _ := client.Do(req)

	bodyRC, _ := ioutil.ReadAll(resp.Body)

	var commandResponse CommandResponse
	json.Unmarshal(bodyRC, &commandResponse)

	return commandResponse
}

func (server *Server) ReplaceFacts(certname string, facts map[string]string) string {
        factsPayload := FactsWireFormat{certname, facts}
	factsJson, _ := json.Marshal(factsPayload)

	commandResponse := server.SubmitCommand("replace facts", 1, string(factsJson[:]))
	uuid := commandResponse.Uuid

	return uuid
}
