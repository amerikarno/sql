package tables

import "gorm.io/gorm"

type Table[T any] struct{}

func NewTable[T any]() *Table[T] {
	return &Table[T]{}
}

func (t *Table[T]) GetAll(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (results []T, err error) {
	if err = db.Debug().Scopes(scopes...).Find(&results).Error; err != nil {
		return nil, err
	}
	return
}
