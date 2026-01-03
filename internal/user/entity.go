package user

import (
	"database/sql/driver"
	"fmt"
	"rest-fiber/internal/post"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "ADMIN"
	Member Role = "MEMBER"
)

type User struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string      `gorm:"type:varchar(255);not null;column:name"`
	Email     string      `gorm:"type:varchar(255);unique;not null;column:email"`
	Password  string      `gorm:"type:varchar(255);not null;column:password"`
	Role      Role        `gorm:"type:role_type;not null;default:'MEMBER'"`
	Posts     []post.Post `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt *time.Time  `gorm:"column:deleted_at;default:null"`
}

func (u *User) TableName() string {
	return "users"
}

func (r *Role) Scan(value any) error {
	if value == nil {
		*r = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*r = Role(string(v))
	case string:
		*r = Role(v)
	default:
		return fmt.Errorf("cannot scan %T into Role", value)
	}
	return nil
}

func (r Role) IsValid() bool {
	return r == Admin || r == Member
}

func (r Role) IsAdmin() bool {
	return r == Admin
}

func (r Role) IsMember() bool {
	return r == Member
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
