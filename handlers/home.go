package handlers

import (
	"encoding/json"
	"fmt"
	"gocielo/types"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	api := struct {
		Title   string
		Version string
	}{
		Title:   "GoCielo",
		Version: "1.0",
	}

	json, err := json.Marshal(api)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
