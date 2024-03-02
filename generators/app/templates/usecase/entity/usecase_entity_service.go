package usecase_<%= entityLowerCase %>

import "app/entity"

type Usecase<%= entityUpCase %> struct {
	repo IRepository<%= entityUpCase %>
}

func NewService(repository IRepository<%= entityUpCase %>) *Usecase<%= entityUpCase %> {
	return &Usecase<%= entityUpCase %>{repo: repository}
}

func (u *Usecase<%= entityUpCase %>) Get(id int) (*entity.Entity<%= entityUpCase %>, error) {
	return u.repo.GetFromID(id)
}

func (u *Usecase<%= entityUpCase %>) GetAll(searchParams entity.SearchEntity<%= entityUpCase %>Params) (response []entity.Entity<%= entityUpCase %>, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *Usecase<%= entityUpCase %>) Create(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {
	return u.repo.Create(<%= entityLowerCase %>)
}

func (u *Usecase<%= entityUpCase %>) Update(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {
	return u.repo.Update(<%= entityLowerCase %>)
}

func (u *Usecase<%= entityUpCase %>) Delete(id int) error {
	return u.repo.Delete(id)
}
