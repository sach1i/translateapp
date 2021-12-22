package translateapp

type Input struct {
	Word   string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
}
