// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LocationReviewInput struct {
	// Written text
	Comment string `json:"comment"`
	// A number from 1 - 5 with 1 being lowest and 5 being highest
	Rating int `json:"rating"`
	// Location Id
	LocationID string `json:"locationId"`
}

type Review struct {
	ID string `json:"id"`
	// Written text
	Comment string `json:"comment"`
	// A number from 1 - 5 with 1 being lowest and 5 being highest
	Rating int `json:"rating"`
	// The location the review is about
	LocationID string `json:"locationID"`
}

func (Review) IsEntity() {}

type SubmitReviewResponse struct {
	// Similar to HTTP status code, represents the status of the mutation
	Code int `json:"code"`
	// Indicates whether the mutation was successful
	Success bool `json:"success"`
	// Human-readable message for the UI
	Message string `json:"message"`
	// Newly created review
	LocationReview *Review `json:"locationReview,omitempty"`
}
