package main

import "net/http"

func mountRoutes() {
	http.HandleFunc("/train", handleTrain)
	http.HandleFunc("/recognize", handleRecognize)
}
