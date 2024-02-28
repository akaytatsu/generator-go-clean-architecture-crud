package usecase_<%= entityLowerCase %>

import "app/entity"

type IRepository<%= entityUpCase %> interface {
	Get<%= entityUpCase %>ByID(id int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll<%= entityUpCase %>() ([]entity.Entity<%= entityUpCase %>, error)
	Create<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error
	Update<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error
	Delete<%= entityUpCase %>(id int) error
}

type IUsecase<%= entityUpCase %> interface {
	Get<%= entityUpCase %>ByID(id int) (*entity.Entity<%= entityUpCase %>, error)
	GetAll<%= entityUpCase %>() ([]entity.Entity<%= entityUpCase %>, error)
	Create<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error
	Update<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error
	Delete<%= entityUpCase %>(id int) error
}
