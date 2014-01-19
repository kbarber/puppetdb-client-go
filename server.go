package puppetdb

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
