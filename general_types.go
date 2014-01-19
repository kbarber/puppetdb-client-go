package puppetdb

/*
Response to servertime query end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/server-time.html#get-v3server-time
*/
type ServerTime struct {
	ServerTime string `json:"server-time"`
}

/*
Response to version query end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/version.html#get-v3version
*/
type Version struct {
	Version string `json:"version"`
}
