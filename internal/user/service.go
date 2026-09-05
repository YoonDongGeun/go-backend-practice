package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, command CreateUserCommand) (UserEntity, error) {
	return s.repo.Create(ctx, command)
}

func (s *Service) Get(ctx context.Context, id int64) (UserEntity, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, limit, offset int32) ([]UserEntity, error) {
	return s.repo.List(ctx, limit, offset)
}
