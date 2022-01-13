package translateapp

import (
	"encoding/json"
	"net/http"
)

// TranslateInput takes request payload and returns translated word
func (a *App) TranslateInput() http.HandlerFunc {
	return func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		a.Logger.Debug("TranslateFunction was triggered")

		ctx := r.Context()

		myService := a.Service
		var i Input
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(response).Encode("Couldn't read request body")
			a.Logger.Errorf("%s", err)
		}

		a.Logger.Debugf("users input:%s", i)

		clientResponse, customErr := myService.Translate(ctx, i)

		if customErr != nil {
			response.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(response).Encode(customErr.Error())
			if err != nil {
				a.Logger.Fatalf("%s", err)
			}

		} else {
			response.WriteHeader(http.StatusOK)
			err := json.NewEncoder(response).Encode(clientResponse)
			if err != nil {
				a.Logger.Fatalf("%s", err)
			}
		}
	}
}

// ListLanguages returns list of all available languages for translation in form {"code":*,"name":*}
func (a *App) ListLanguages() http.HandlerFunc {
	return func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		a.Logger.Debug("ListLanguages function was triggered")
		ctx := r.Context()
		myService := a.Service
		clientResponse, customErr := myService.Languages(ctx)

		if customErr != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(customErr.Error()))

		} else {
			response.WriteHeader(http.StatusOK)
			err := json.NewEncoder(response).Encode(clientResponse)
			if err != nil {
				a.Logger.Fatalf("%s", err)
			}
		}
	}
}
