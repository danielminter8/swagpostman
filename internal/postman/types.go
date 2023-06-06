package postman

type CollectionInfo struct {
	Schema string `json:"schema"`
}

type Request struct {
	Method  string                 `json:"method"`
	Header  []interface{}          `json:"header"`
	URL     string                 `json:"url"`
	Body    map[string]interface{} `json:"body"`
	Auth    interface{}            `json:"auth"`
	Proxy   interface{}            `json:"proxy"`
	TLS     interface{}            `json:"tls"`
	Version string                 `json:"version"`
}

type Item struct {
	Name     string        `json:"name"`
	Protocol string        `json:"protocol"`
	Request  Request       `json:"request"`
	Response []interface{} `json:"response"`
}

type Collection struct {
	Collection struct {
		Info CollectionInfo `json:"info"`
		Item []Item         `json:"item"`
	} `json:"collection"`
}
