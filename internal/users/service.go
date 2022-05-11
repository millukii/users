package users

import "context"

type UserService interface {
	Save(ctx context.Context, user *User) error
	Get(ctx context.Context,id string) (*User, error)
	GetAll(ctx context.Context,) ([]User, error)
	Delete(ctx context.Context,id string) error
	Update(ctx context.Context,user *User) error
}

type svc struct {
	client Client
	repo Repository
}

func NewUserService(client Client, repo Repository) UserService {
	return &svc{
		client: client,
		repo: repo,
	}
}

func (svc *svc)	Save(ctx context.Context, user *User) error {
	svc.client.Get(ctx, user.Name, user.ID)

	//enriquecer datos
	//validaciones
	return svc.repo.Save(ctx, user)
}
func (svc *svc)		Get(ctx context.Context,id string) (*User, error) {
	return svc.repo.Get(ctx, id)
}
func (svc *svc)		GetAll(ctx context.Context,) ([]User, error){
	return svc.repo.GetAll(ctx)
}
func (svc *svc)		Delete(ctx context.Context,id string) error{
	return svc.repo.Delete(ctx, id)
}
func (svc *svc)		Update(ctx context.Context,user *User) error{
	return svc.repo.Update(ctx, user)
}