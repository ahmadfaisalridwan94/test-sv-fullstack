package models

import (
	"errors"
	"html"
	"strings"
	"time"	
	"github.com/jinzhu/gorm"
)

type Post struct {
	Id        uint32    `gorm:"primary_key;auto_increment;column:Id;" json:"Id"`
	Title     string    `gorm:"size:200;not null;unique;column:Title;" json:"Title"`
	Content   string    `gorm:"type:text;not null;column:Content;" json:"Content"`
	Category  string    `gorm:"size:100;not null;column:Category;" json:"Category"`	
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:CreatedAt;" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:UpdatedAt;" json:"UpdatedAt"`
	Status    string    `gorm:"size:100;not null;column:Status;" json:"Status"`	
}

type Errors map[string][]string

func (p *Post) Prepare() {
	p.Id = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Category = html.EscapeString(strings.TrimSpace(p.Category))
	p.Status = html.EscapeString(strings.TrimSpace(p.Status))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() Errors {
	errors := make(Errors)

	if p.Title == "" {
		errors["Title"] = append(errors["Title"], "Required Title")
	}
	if p.Content == "" {
		errors["Content"] = append(errors["Content"], "Required Content")
	}
	if p.Category == "" {
		errors["Category"] = append(errors["Category"], "Required Category")
	}
	if p.Status == "" {
		errors["Status"] = append(errors["Status"], "Required Status")
	}

	if p.Title != "" && len(p.Title) < 20 {
		errors["Title"] = append(errors["Title"], "Title must be at least 20 characters")
	}
	if p.Content != "" && len(p.Content) < 200 {
		errors["Content"] = append(errors["Content"], "Content must be at least 200 characters")
	}
	if p.Category != "" && len(p.Category) < 3 {
		errors["Category"] = append(errors["Category"], "Category must be at least 3 characters")
	}
	if p.Status != "publish" && p.Status != "draft" && p.Status != "trash" {
		errors["Status"] = append(errors["Status"], "Invalid status")
	}

	return errors
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}	
	return p, nil
}

func (p *Post) FindAllPosts(db *gorm.DB, limit int, offset int, status *string) (*[]Post, error) {
	var err error
	posts := []Post{}

	query := db.Debug().Model(&Post{}).Limit(limit).Offset(offset)

	if status != nil && *status != "" {
		query = query.Where("status = ?", *status)
	}

	err = query.Find(&posts).Error

	if err != nil {
		return &[]Post{}, err
	}

	return &posts, nil
}

func (p *Post) FindPostById(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}	
	return p, nil
}

func (p *Post) UpdateAPost(db *gorm.DB) (*Post, error) {

	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", p.Id).Updates(Post{Title: p.Title, Content: p.Content, Category: p.Category, Status: p.Status, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Post{}, err
	}	
	return p, nil
}

func (p *Post) DeleteAPost(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Post) CountData(db *gorm.DB, status *string) (int64, error) {
	var count int64

	query := db.Debug().Model(&Post{}).Where("status = ?", *status).Count(&count)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("Post not found")
		}
		return 0, query.Error
	}

	return count, nil
}