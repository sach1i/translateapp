package models

type Input struct {
	Word   string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type ClientSuccessResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type Data struct {
	RespData string `json:"field"`
}

type ClientErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code int    `json:"code"`
	Data []byte `json:"data"'`
}
