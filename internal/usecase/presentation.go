package usecase

import "examples/kahootee/internal/entity"

type presentation struct {
	repo PresentationRepo
}

func NewPresentation(repo PresentationRepo) Presentation {
	return &presentation{
		repo: repo,
	}
}

func (r presentation) CreatePresentation(request *entity.Presentation) error {
	return r.repo.CreatePresentation(request)
}

func (r presentation) GetPresentation(id uint32) (*entity.Presentation, error) {
	return r.repo.GetPresentation(id)
}

func (r presentation) GroupCollection(groupId uint32) ([]*entity.Presentation, error) {
	return r.repo.GroupCollection(groupId)
}

func (r presentation) MyCollection(userId uint32) ([]*entity.Presentation, error) {
	return r.repo.MyCollection(userId)
}

func (r presentation) UpdatePresentation(request *entity.Presentation) error {
	return r.repo.UpdatePresentation(request)
}

func (r presentation) DeletePresentation(id uint32, userId uint32) error {
	return r.repo.DeletePresentation(id, userId)
}

func (r presentation) CreateSlide(slide *entity.Slide) error {
	return r.repo.CreateSlide(slide)
}

func (r presentation) UpdateSlide(slide *entity.Slide) error {
	return r.repo.UpdateSlide(slide)
}

func (r presentation) DeleteSlide(id uint32) error {
	return r.repo.DeleteSlide(id)
}
