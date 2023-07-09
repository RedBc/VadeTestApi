package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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

// DeleteDocumentHandler deletes the document with given ID.
func DeleteDocumentHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read document ID
		documentID := mux.Vars(r)["id"]
		err := entity.DeleteDocument(documentID)
		if err != nil {
			// Check if it is No document error or any other error
			if errors.Is(err, entity.ErrNoDocument) {
				// Write Header if no related document found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write Header with Accepted Status (done operation)
		rw.WriteHeader(http.StatusAccepted)
	}
}
