package users

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user *User) error
}
type repo struct {
}

func NewUserRepository() Repository {
	return &repo{}
}

func (r *repo) Save(ctx context.Context, user *User) error {
	return nil
}
func (r *repo) Get(ctx context.Context, id string) (*User, error) {
	return nil, nil
}
func (r *repo) GetAll(ctx context.Context) ([]*User, error) {
	return nil, nil
}
func (r *repo) Delete(ctx context.Context, id string) error {
	return nil
}
func (r *repo) Update(ctx context.Context, user *User) error {
	return nil
}