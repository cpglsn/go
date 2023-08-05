package main

import (
	"encoding/json"
	"net/http"
)

type envelop map[string]any

func WriteJSON(w http.ResponseWriter, data envelop) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)

	return nil
}
