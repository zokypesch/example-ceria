package model

import "github.com/jinzhu/gorm"

// Article struct for table article
type Article struct {
	gorm.Model
	Title    string    `json:"title" binding:"required"`
	Author   string    `json:"author" binding:"required"`
	Comments []Comment `gorm:"foreignkey:ArticleID"`
}

// Comment struct for table article
type Comment struct {
	gorm.Model
	ArticleID   uint   `json:"article_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
