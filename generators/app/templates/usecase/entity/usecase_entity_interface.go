package usecase_<%= entityLowerCase %>

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_<%= entityLowerCase %>.go -package=mocks app/usecase/<%= entityLowerCase %> IRepository<%= entityUpCase %>
type IRepository<%= entityUpCase %> interface {
	GetFromID(id int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll(searchParams entity.SearchEntity<%= entityUpCase %>Params) (response []entity.Entity<%= entityUpCase %>, totalRegisters int64, err error)
	Create(*entity.Entity<%= entityUpCase %>) error
	Update(*entity.Entity<%= entityUpCase %>) error
	Delete(id int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_<%= entityLowerCase %>.go -package=mocks app/usecase/<%= entityLowerCase %> IUsecase<%= entityUpCase %>
type IUsecase<%= entityUpCase %> interface {
	Get(id int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll(searchParams entity.SearchEntity<%= entityUpCase %>Params) (response []entity.Entity<%= entityUpCase %>, totalRegisters int64, err error)
	Create(*entity.Entity<%= entityUpCase %>) error
	Update(*entity.Entity<%= entityUpCase %>) error
	Delete(id int) error
}
