// novel/usecase/novel_usecase_impl.go
package usecase

import (
	"module-crud/novel/domain"
	"module-crud/novel/domain/model"
)

type novelUseCase struct {
	repo domain.NovelRepo
}

func NewNovelUseCase(repo domain.NovelRepo) domain.NovelUseCase {
	return &novelUseCase{repo: repo}
}

func (uc *novelUseCase) CreateNovel(novel model.Novel) error {
	return uc.repo.CreateNovel(novel)
}

func (uc *novelUseCase) GetNovelById(id int) (model.Novel, error) {
	return uc.repo.GetNovelById(id)
}
