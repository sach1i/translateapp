package models

type Input struct {
	Word   string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type GetRes struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    LanguageList `json:"data"`
}
type LanguageList struct { //languageList
	Languages []Language
}
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type PostRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    Translation `json:"data"`
}

type Translation struct {
	Text string `json:"translatedText"`
}

//type ClientSuccessResponse struct {
//	Code int    `json:"code"`
//	Data interface{} `json:"data"`
//}

type ClientErrResponse struct {
	Code    int `json:"code"`
	Message Error
}

type Error struct {
	Message string `json:"error"`
}
