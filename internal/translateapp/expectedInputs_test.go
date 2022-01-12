package translateapp_test

import (
	"translateapp/internal/libretranslate"
	"translateapp/internal/translateapp"
)

var expectedGetOutput = &libretranslate.LanguageList{
	Languages: []libretranslate.Language{
		{
			Name: "polish",
			Code: "pl",
		},
		{
			Name: "english",
			Code: "en",
		},
	},
}

var expectedPostOutput = &libretranslate.Translation{Text: "mysz"}

var mockedClientInput = libretranslate.Input{
	Word:   "mouse",
	Source: "en",
	Target: "pl",
}
var mockedServiceInput = translateapp.Input{
	Word:   "mouse",
	Source: "en",
	Target: "pl",
}
