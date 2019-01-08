package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/zokypesch/example-ceria/model"
)

// ListMigration list of migration
type ListMigration struct{}

// NewListMigration migration service
func NewListMigration() *ListMigration {
	return &ListMigration{}
}

// Migrate for migartion data list
func (list *ListMigration) Migrate(db *gorm.DB) error {
	db.AutoMigrate(&model.Article{}, &model.Comment{})

	return nil
}
