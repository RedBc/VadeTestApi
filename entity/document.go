package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Document struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetDocuments returns the JSON file content if available else returns an error.
func GetDocuments() ([]byte, error) {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

// AddDocument adds an input document to the document list in JSON document.
func AddDocument(document Document) error {
	// Load existing documents and append the data to document list
	var documents []Document
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	// Load our JSON file to memory using array of documents
	err = json.Unmarshal(data, &documents)
	if err != nil {
		return err
	}
	// Add new Document to our list
	documents = append(documents, document)

	// Write Updated JSON file
	updatedData, err := json.Marshal(documents)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// ErrNoDocument is used if no document found
var ErrNoDocument = errors.New("no document found")

// removeElement is used to remove element from document array at given index
func removeElement(arr []Document, index int) []Document {
	ret := make([]Document, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}

// Deletedocument takes id as input and deletes the corresponding document, else it returns ErrNoDocument error.
func DeleteDocument(id string) error {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	// read documents
	var documents []Document
	err = json.Unmarshal(data, &documents)
	if err != nil {
		return err
	}
	// iterate through document array
	for i := 0; i < len(documents); i++ {
		// if we find one document with the given ID
		if documents[i].ID == id {
			documents = removeElement(documents, i)
			// Write Updated JSON file
			updatedData, err := json.Marshal(documents)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoDocument
}
