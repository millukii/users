package users

import (
	"context"

)

type Repository interface {
	Save(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user *User) error
}
type repo struct {
	userRepo []User
}

func NewUserRepository(userRepo []User) Repository {
	return &repo{
		userRepo: userRepo,
	}
}

func (r *repo) Save(ctx context.Context, user *User) error {
	r.userRepo = append(r.userRepo, *user)
	return nil
}
func (r *repo) Get(ctx context.Context, id string) (*User, error) {
 var user *User

	for i := range r.userRepo {
			if r.userRepo[i].ID == id {
					// Found!
					user =  &r.userRepo[i]
					break
			}
	}

	return user, nil
}
func (r *repo) GetAll(ctx context.Context) ([]User, error) {

	return r.userRepo, nil
}
func (r *repo) Delete(ctx context.Context, id string) error {
	return nil
}
func (r *repo) Update(ctx context.Context, user *User) error {
	return nil
}