package httpHandler

import (
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fileBytes, err := ioutil.ReadFile("internal/services/httpHandler/templates/index.html")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		_, err = w.Write(fileBytes)
		if err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
