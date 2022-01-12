package libretranslate_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
	"translateapp/internal/mocks"
)

func TestClient_GetLanguages(t *testing.T) {
	testClient := mocks.Client{}
	expectedGetOutput := &libretranslate.LanguageList{
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
	testClient.On("GetLanguages", mock.Anything).Return(expectedGetOutput, nil)
	realClient := libretranslate.NewClient(logging.DefaultLogger())
	response, err := realClient.GetLanguages(context.Background())
	assert.Equal(t, expectedGetOutput, response)
	assert.Nil(t, err)

}
