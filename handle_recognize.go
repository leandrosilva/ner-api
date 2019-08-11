package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type recognizeRequest struct {
	Model   string `json:model`
	Content string `json:content`
}

type recognizeResponse struct {
	Success  bool          `json:success`
	Model    string        `json:model`
	Entities []entityCount `json:entities`
}

func handleRecognize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	request, err := getRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := recognize(request.Content, request.Model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(recognizeResponse{
		Success:  true,
		Model:    request.Model,
		Entities: result.Entities})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getRequest(r *http.Request) (recognizeRequest, error) {
	var request recognizeRequest

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(body, &request)

	return request, err
}
