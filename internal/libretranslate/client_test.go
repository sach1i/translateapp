package libretranslate_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
)

func TestClient_GetLanguages(t *testing.T) {
	expectedGetOutput := &libretranslate.LanguageList{
		Languages: []libretranslate.Language{
			{
				Name: "English",
				Code: "en",
			},
			{
				Name: "Polish",
				Code: "pl",
			},
		},
	}
	const BaseURLLibre = "http://localhost:5000"
	realClient := libretranslate.NewClient(logging.DefaultLogger(), BaseURLLibre)
	response, err := realClient.GetLanguages(context.Background())
	assert.Equal(t, expectedGetOutput, response)
	assert.Nil(t, err)
}
func TestClient_Translate(t *testing.T) {
	expectedPostOutput := &libretranslate.Translation{
		Text: "mysz",
	}
	testInput := libretranslate.Input{
		Word:   "mouse",
		Source: "en",
		Target: "pl",
	}
	const BaseURLLibre = "http://localhost:5000"
	realClient := libretranslate.NewClient(logging.DefaultLogger(), BaseURLLibre)
	response, err := realClient.Translate(context.Background(), testInput)
	assert.Equal(t, expectedPostOutput, response)
	assert.Nil(t, err)
}
