package model

import "github.com/jinzhu/gorm"

// Article struct for table article
type Article struct {
	gorm.Model
	Title  string  `json:"title" binding:"required"`
	Tag    string  `json:"tag"`
	Body   string  `json:"body" binding:"required"`
	Author *Author `gorm:"foreignkey:AuthorID" json:"author" binding:"-" ceria:"ignoreStructField"`
	//field Author is type non slice hirarcy struct please add tag ignore elastic, because elastic cannot convert to string
	AuthorID uint      `json:"author_id"` //  binding:"nefield=Author.ID"
	Comments []Comment `gorm:"foreignkey:ArticleID" ceria:"ignoreStructField"`
	Custom   string
}

// Comment struct for table article
type Comment struct {
	gorm.Model
	ArticleID uint   `json:"article_id" binding:"required"`
	Fullname  string `json:"fullname" binding:"required"`
	Body      string `json:"body" binding:"required"`
}

// Author struct for table author
type Author struct {
	gorm.Model
	Fullname string `json:"fullname" binding:"required"`
}
