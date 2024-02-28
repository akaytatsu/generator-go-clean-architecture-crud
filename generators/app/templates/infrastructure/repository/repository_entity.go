package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type Repository<%= entityUpCase %> struct {
	DB *gorm.DB
}

func New<%= entityUpCase %>Postgres(DB *gorm.DB) *Repository<%= entityUpCase %> {
	return &Repository<%= entityUpCase %>{DB: DB}
}

func (r *Repository<%= entityUpCase %>) Get<%= entityUpCase %>ByID(id int) (<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>, err error) {
	r.DB.First(&<%= entityLowerCase %>, id)

	return <%= entityLowerCase %>, err
}

func (r *Repository<%= entityUpCase %>) GetAll<%= entityUpCase %>s() (<%= entityLowerCase %>s []entity.Entity<%= entityUpCase %>, err error) {
	<%= entityLowerCase %>s = make([]entity.Entity<%= entityUpCase %>, 0)

	err = r.DB.Find(&<%= entityLowerCase %>s).Error

	return <%= entityLowerCase %>s, err
}

func (r *Repository<%= entityUpCase %>) Create<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {

	return r.DB.Create(&<%= entityLowerCase %>).Error
}

func (r *Repository<%= entityUpCase %>) Update<%= entityUpCase %>(<%= entityLowerCase %> *entity.Entity<%= entityUpCase %>) error {

	_, err := r.Get<%= entityUpCase %>ByID(<%= entityLowerCase %>.ID)

	if err != nil {
		return err
	}

	return r.DB.Save(&<%= entityLowerCase %>).Error
}

func (r *Repository<%= entityUpCase %>) Delete<%= entityUpCase %>(id int) error {

	<%= entityLowerCase %>, err := r.Get<%= entityUpCase %>ByID(id)

	if err != nil {
		return err
	}

	return r.DB.Delete(&<%= entityLowerCase %>).Error
}
