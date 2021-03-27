package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demographql/database"
	"demographql/graph/generated"
	"demographql/graph/model"
	"demographql/utils"
	"fmt"
	"log"
	"strconv"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link database.Links
	link.Title = input.Title
	link.Address = input.Address
	res := r.DB.Create(&link)

	if res.Error != nil {
		log.Panic("Insert Error.")
	}
	return &model.Link{ID: strconv.FormatInt(int64(link.ID), 10), Title: link.Title, Address: link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []database.Links
	link := database.Links{}
	err := link.GetAll(r.DB, &links, 100, 0)
	if err != nil {
		return nil, err
	}
	reslinks := []*model.Link{}
	for _, linkr := range links {
		tmp := &model.Link{
			ID:      string(rune(linkr.ID)),
			Title:   linkr.Title,
			Address: linkr.Address,
		}
		if linkr.UserID != nil {
			tmp.User = &model.User{
				ID:   string(rune(*linkr.UserID)),
				Name: string(linkr.Users.Username),
			}
		}
		reslinks = append(reslinks, tmp)
	}

	if len(reslinks) > 0 {
		return reslinks, nil
	} else {
		if err != nil {
			return nil, utils.New("error")
		}
	}
	return []*model.Link{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
