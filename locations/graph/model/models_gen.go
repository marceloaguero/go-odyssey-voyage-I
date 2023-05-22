// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Location struct {
	ID string `json:"id"`
	// The name of the location
	Name string `json:"name"`
	// A short description about the location
	Description string `json:"description"`
	// The location's main photo as a URL
	Photo string `json:"photo"`
}

type LocationInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}
