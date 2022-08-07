package dtos

type Payload struct {
	Body   []byte `json:"body,omitempty"`
	Url    string `json:"url"`
	Method string `json:"method"`
}
