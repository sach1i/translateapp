package translateapp_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
	"translateapp/internal/mocks"
	"translateapp/internal/translateapp"
)

func TestService_Languages(t *testing.T) {
	expectedServiceOutput := translateapp.Response{Data: *expectedGetOutput}
	lt := mocks.Client{}
	lt.On("GetLanguages", mock.Anything).Return(expectedGetOutput, nil)
	tr := mocks.Translator{}
	tr.On("Translate")
	service := translateapp.NewService(&lt, &tr, logging.DefaultLogger())
	res, err := service.Languages(context.Background())
	require.NoError(t, err)
	require.Equal(t, &expectedServiceOutput, res)

}

func TestService_Translate(t *testing.T) {

	expectedServiceOutput := translateapp.Response{Data: "mysz"}
	lt := mocks.Client{}
	mockedClientInput := libretranslate.Input{
		Word:   "mouse",
		Source: "en",
		Target: "pl",
	}
	mockedServiceInput := translateapp.Input{
		Word:   "mouse",
		Source: "en",
		Target: "pl",
	}

	tr := mocks.Translator{}
	tr.On("Translate", mock.Anything, mockedClientInput).Return("mysz", nil)
	service := translateapp.NewService(&lt, &tr, logging.DefaultLogger())
	res, err := service.Translate(context.Background(), mockedServiceInput)
	require.NoError(t, err)
	require.Equal(t, &expectedServiceOutput, res)

}
