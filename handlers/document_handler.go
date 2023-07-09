package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"vade.com/vade/goDocumentAPI/entity"
)

func GetDocumentsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := entity.GetDocuments()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write the body with JSON data
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

func CreateDocumentHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var document entity.Document
		err = json.Unmarshal(data, &document)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		err = entity.AddDocument(document)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// return after writing Body
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New document"))
	}
}
