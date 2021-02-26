package models

// Query is a basic struct for the input we expect to receive.
type Query struct {
	Token       string `json:"token"`
	Body        string `json:"body"`
	TempVersion string `json:"ver"` // a temporary parameter for version testing until we implement dbs
}
