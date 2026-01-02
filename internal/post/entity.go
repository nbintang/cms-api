package post

import (
	"database/sql/driver"
	"fmt"
	"rest-fiber/internal/category"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Published Status = "PUBLISHED"
	Draft     Status = "DRAFT"
)

type Post struct {
	ID         string            `gorm:"type:char(36);primaryKey;column:id"`
	Title      string            `gorm:"type:varchar(255);not null;column:title"`
	Body       string            `gorm:"type:text;not null;column:body"`
	UserID     string            `gorm:"type:char(36);not null;column:user_id"`
	CategoryID string            `gorm:"type:char(36);not null;column:category_id"`
	Status     Status            `gorm:"type:enum('PUBLISHED','DRAFT');not null;default:'DRAFT'"`
	Category   category.Category `gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt  time.Time         `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time         `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  *time.Time        `gorm:"column:deleted_at;default:null"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) IsPublished() bool {
	return p.Status == Published
}
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return nil
}

func (r *Status) Scan(value any) error {
	if value == nil {
		*r = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*r = Status(string(v))
	case string:
		*r = Status(v)
	default:
		return fmt.Errorf("cannot scan %T into Role", value)
	}
	return nil
}

func (r Status) Value() (driver.Value, error) {
	return string(r), nil
}
