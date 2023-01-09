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

func (r presentation) CreatePresentation(request *entity.Presentation) (uint32, error) {
	return r.repo.CreatePresentation(request)
}

func (r presentation) GetPresentation(id uint32) (*entity.Presentation, error) {
	return r.repo.GetPresentation(id)
}

func (r presentation) Collection(groupId uint32) ([]*entity.Presentation, error) {
	return r.repo.Collection(groupId)
}

func (r presentation) UpdatePresentation(request *entity.Presentation) error {
	return r.repo.UpdatePresentation(request)
}

func (r presentation) DeletePresentation(id uint32, userId uint32) error {
	return r.repo.DeletePresentation(id, userId)
}
