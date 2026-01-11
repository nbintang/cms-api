package category

import "gorm.io/gorm"

type ScopeReturn func(db *gorm.DB) *gorm.DB

func SelectedFields(db *gorm.DB) *gorm.DB {
	return db.Select("id", "name")
}

func Paginate(limit, offset int) ScopeReturn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

func WhereID(id string) ScopeReturn {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}