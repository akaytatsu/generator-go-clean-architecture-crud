package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Entity<%= entityUpCase %> struct {
	ID                 int       `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name" gorm:"varchar(50)"`
	Active             bool      `json:"active" gorm:"default:true"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewEntity<%= entityUpCase %>(entity<%= entityUpCase %>Param Entity<%= entityUpCase %>) (*Entity<%= entityUpCase %>, error) {
	u := &Entity<%= entityUpCase %>{
		ID:        entity<%= entityUpCase %>Param.ID,
		Name:      entity<%= entityUpCase %>Param.Name,
		Active:    entity<%= entityUpCase %>Param.Active,
		CreatedAt: entity<%= entityUpCase %>Param.CreatedAt,
		UpdatedAt: entity<%= entityUpCase %>Param.UpdatedAt,
	}

	return u, nil
}

func (u *Entity<%= entityUpCase %>) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntity<%= entityUpCase %>Params struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
