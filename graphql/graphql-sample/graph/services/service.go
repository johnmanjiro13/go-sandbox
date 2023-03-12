package services

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/johnmanjiro13/graphql-sample/graph/model"
)

type Services interface {
	UserService
}

type services struct {
	*userService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}
