package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type trainResponse struct {
	Success bool   `json:success`
	Label   string `json:label`
}

func handleTrain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath, err := uploadFile(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := train(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(trainResponse{
		Success: true,
		Label:   result.Label})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func uploadFile(w http.ResponseWriter, r *http.Request) (string, error) {
	// Reference: https://tutorialedge.net/golang/go-file-upload-tutorial/
	log.Println("Uploading file...")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `dataset`,
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, meta, err := r.FormFile("dataset")
	if err != nil {
		return "", err
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", meta.Filename)
	log.Printf("File Size: %+v\n", meta.Size)
	log.Printf("MIME Header: %+v\n", meta.Header)

	// Create a temporary file within our temp directory that follows
	// a particular naming pattern
	datasetDir := filepath.Join(".", "datasets")
	os.MkdirAll(datasetDir, os.ModePerm)
	tempFile, err := ioutil.TempFile(datasetDir, "dataset-train-"+meta.Filename)
	if err != nil {
		log.Println("Cannot create temp file:", err)
		return "", err
	}
	defer tempFile.Close()

	// Read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Cannot read temp file:", err)
		return "", err
	}

	// Write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Println("Cannot write to temp file:", err)
	}

	return tempFile.Name(), nil
}
