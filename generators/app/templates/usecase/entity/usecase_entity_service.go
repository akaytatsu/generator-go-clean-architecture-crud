package usecase_<%= entityLowerCase %>

import "app/entity"

type UseCase<%= entityUpCase %> struct {
	repo IRepository<%= entityUpCase %>
}

func NewService(repository IRepository<%= entityUpCase %>) *UseCase<%= entityUpCase %> {
	return &UseCase<%= entityUpCase %>{repo: repository}
}

func (u *UseCase<%= entityUpCase %>) Get<%= entityUpCase %>ByID(id int) (*entity.Entity<%= entityUpCase %>, error) {
	return u.repo.Get<%= entityUpCase %>ByID(id)
}

func (u *UseCase<%= entityUpCase %>) GetAll<%= entityUpCase %>s() ([]entity.Entity<%= entityUpCase %>, error) {
	return u.repo.GetAll<%= entityUpCase %>s()
}

func (u *UseCase<%= entityUpCase %>) Create<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {
	return u.repo.Create<%= entityUpCase %>(<%= entityLowerCase %>)
}

func (u *UseCase<%= entityUpCase %>) Update<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {
	return u.repo.Update<%= entityUpCase %>(<%= entityLowerCase %>)
}

func (u *UseCase<%= entityUpCase %>) Delete<%= entityUpCase %>(id int) error {
	return u.repo.Delete<%= entityUpCase %>(id)
}
