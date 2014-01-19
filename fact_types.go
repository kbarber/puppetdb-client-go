package puppetdb

/*
Fact command submission struct for submitting the 'replace facts' command
to PuppetDB.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/facts_format.html
*/
type FactsWireFormat struct {
	// Certificate name of node to replace facts for
	Name string `json:"name"`
	// A map of fact key/value pairs
	Values map[string]string `json:"values"`
}

/*
Response for fact based query end-points.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/facts.html#get-v3facts
*/
type Fact struct {
	Certname string `json:"certname"`
	Name string `json:"name"`
	Value string `json:"value"`
}
