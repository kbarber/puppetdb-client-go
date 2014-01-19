package puppetdb

/*
Top level struct representing a PuppetDB commands payload object.

See here for more details on the protocol: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
type CommandObject struct {
	// Command name, such as 'replace facts' or 'deactivate node'
	Command string `json:"command"`
	// Command version as an integer
	Version	int `json:"version"`
	// Command payload, may contain different data types depending on command
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
	// A UUID returned by the server uniquely identifying a command submission
	Uuid string `json:"uuid"`
}
