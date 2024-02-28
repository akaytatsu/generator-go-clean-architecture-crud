package entity

type Entity<%= entityUpCase %> struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	CreatedAt       int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
