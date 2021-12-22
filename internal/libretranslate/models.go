package libretranslate

import "fmt"

type LanguageList struct { //languageList
	Languages []Language
}
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Translation struct {
	Text string `json:"translatedText"`
}

type CustomError struct {
	Code    int `json:"code"`
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
