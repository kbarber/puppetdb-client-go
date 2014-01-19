package puppetdb

/*
A representation of a report wire format.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/report_format.html#report-interchange-format
*/
type ReportWireFormat struct {
	Certname string `json:"certname"`
	PuppetVersion string `json:"puppet-version"`
	ReportFormat int `json:"report-format"`
	ConfigurationVersion string `json:"configuration-version"`
	StartTime string `json:"start-time"`
	EndTime string `json:"end-time"`
	ResourceEvents []ResourceEvent `json:"resource-events"`
	TransactionUuid string `json:"transaction-uuid"`
}

/*
A representation of a resource even from a report.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/report_format.html#data-type-resource-event
*/
type ResourceEvent struct {
	ResourceType string `json:"resource-type"`
	ResourceTitle string `json:"resource-title"`
	Property string `json:"property"`
	Timestamp string `json:"timestamp"`
	Status string `json:"status"`
	OldValue string `json:"old-value"`
	NewValue string	`json:"new-value"`
	Message string `json:"message"`
	File string `json:"file"`
	Line int `json:"line"`
	ContainmentPath []string `json:"containment-path"`
}

/*
Response data structure for event-counts query end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/event-counts.html#get-v3event-counts
*/
type EventCounts struct {
	SubjectType string `json:"subject-type"`
	Subject string `json:"subject"`
	Failures string `json:"failures"`
	Successes string `json:"successes"`
	Noops string `json:"noops"`
	Skips string `json:"skips"`
}

/*
Response data structure for aggregate-event-counts query end-points.

More detail here: http://docs.puppetlabs.com/puppetdb/1.6/api/query/v3/aggregate-event-counts.html#get-v3aggregate-event-counts
*/
type AggregateEventCounts struct {
	Failures string `json:"failures"`
	Successes string `json:"successes"`
	Noops string `json:"noops"`
	Skips string `json:"skips"`
	Total string `json:"total"`
}
