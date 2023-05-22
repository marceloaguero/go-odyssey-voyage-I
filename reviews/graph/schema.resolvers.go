package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/graph/model"
)

// SubmitReview is the resolver for the submitReview field.
func (r *mutationResolver) SubmitReview(ctx context.Context, locationReview *model.LocationReviewInput) (*model.SubmitReviewResponse, error) {
	panic(fmt.Errorf("not implemented: SubmitReview - submitReview"))
}

// LatestReviews is the resolver for the latestReviews field.
func (r *queryResolver) LatestReviews(ctx context.Context) ([]*model.Review, error) {
	panic(fmt.Errorf("not implemented: LatestReviews - latestReviews"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
