package domain

import "module-crud/novel/domain/model"

type NovelUseCase interface {
	CreateNovel(novel model.Novel) error
	GetNovelById(id int) (model.Novel, error)
}
