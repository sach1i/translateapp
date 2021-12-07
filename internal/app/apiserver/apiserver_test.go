package apiserver_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"translateapp/internal/app/apiserver"
)

func TestListLanguages(t *testing.T) {
	s := apiserver.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/languages", nil)
	w := httptest.NewRecorder()
	s.ListLanguages().ServeHTTP(w, req)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body) //nolint:wsl
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)

	}
	expected := map[string]string{"code": "code", "name": "name"}
	tc, err := json.Marshal(expected)
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(data, tc) {
		t.Errorf("Expected %s, got %s", tc, data)
	}
}

func TestTranslateInput(t *testing.T) {
	s := apiserver.NewServer()
	data := map[string]string{"word": "testword", "source": "en", "target": "es"}
	postData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/translate", bytes.NewBuffer(postData))
	w := httptest.NewRecorder()
	s.TranslateInput().ServeHTTP(w, req)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body) //nolint:wsl
	outputData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	expected := map[string]string{"TranslatedWord": "testword"}
	tc, err := json.Marshal(expected)
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(outputData, tc) {
		t.Errorf("Expected %s, got %s", tc, outputData)
	}
}
