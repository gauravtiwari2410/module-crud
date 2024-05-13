package repo

import (
	"module-crud/novel/domain/model"

	"gorm.io/gorm"
)

type NovelRepo struct {
	db *gorm.DB
}

func NewNovelRepo(db *gorm.DB) *NovelRepo {
	return &NovelRepo{db: db}
}

func (r *NovelRepo) CreateNovel(novel model.Novel) error {
	return r.db.Create(&novel).Error
}

func (r *NovelRepo) GetNovelById(id int) (model.Novel, error) {
	var novel model.Novel
	result := r.db.First(&novel, id)
	return novel, result.Error
}
