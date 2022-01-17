package translateapp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
	"translateapp/internal/mocks"
	"translateapp/internal/translateapp"
)

func TestApp_ListLanguages(t *testing.T) {
	expectedServiceOutput := translateapp.Response{Data: *expectedGetOutput}
	srvc := mocks.Service{}
	srvc.On("Languages", mock.Anything).Return(&expectedServiceOutput, nil)
	app := translateapp.NewApp(&srvc, logging.DefaultLogger())
	req, err := http.NewRequest("GET", "/languages", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := app.ListLanguages()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	var actualResponse translateapp.Response

	json.NewDecoder(rr.Body).Decode(&actualResponse)
	jsonbody, err := json.Marshal(actualResponse.Data)
	responseContents := libretranslate.LanguageList{}
	if err := json.Unmarshal(jsonbody, &responseContents); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", responseContents)
	actual := translateapp.Response{Data: responseContents}
	assert.Equal(t, expectedServiceOutput, actual)
}

func TestApp_TranslateInput(t *testing.T) {
	expectedServiceOutput := translateapp.Response{Data: *expectedPostOutput}
	srvc := mocks.Service{}
	srvc.On("Translate", mock.Anything, mockedServiceInput).Return(&expectedServiceOutput, nil)
	app := translateapp.NewApp(&srvc, logging.DefaultLogger())
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mockedClientInput); err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/translate", &body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := app.TranslateInput()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	var actualResponse translateapp.Response

	json.NewDecoder(rr.Body).Decode(&actualResponse)
	jsonbody, err := json.Marshal(actualResponse.Data)
	responseContents := libretranslate.Translation{}
	if err := json.Unmarshal(jsonbody, &responseContents); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", responseContents)
	actual := translateapp.Response{Data: responseContents}
	assert.Equal(t, expectedServiceOutput, actual)
}
