package entity

import (
	"io/ioutil"
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
