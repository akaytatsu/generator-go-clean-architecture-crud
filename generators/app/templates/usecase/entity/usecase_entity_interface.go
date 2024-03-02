package usecase_<%= entityLowerCase %>

import "app/entity"

type IRepository<%= entityUpCase %> interface {
	GetFromYear(year int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll(searchParams entity.SearchEntity<%= entityUpCase %>Params) (response []entity.Entity<%= entityUpCase %>, totalRegisters int64, err error)
	Create(*entity.Entity<%= entityUpCase %>) error
	Update(*entity.Entity<%= entityUpCase %>) error
	Delete(id int) error
}

type IUsecase<%= entityUpCase %> interface {
	Get(year int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll(searchParams entity.SearchEntity<%= entityUpCase %>Params) (response []entity.Entity<%= entityUpCase %>, totalRegisters int64, err error)
	Create(*entity.Entity<%= entityUpCase %>) error
	Update(*entity.Entity<%= entityUpCase %>) error
	Delete(id int) error
}
