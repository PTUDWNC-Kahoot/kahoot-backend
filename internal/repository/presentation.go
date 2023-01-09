package repo

import (
	"examples/kahootee/internal/entity"
	"examples/kahootee/internal/usecase"

	"gorm.io/gorm"
)

type presentationRepo struct {
	db *gorm.DB
}

func NewPresentationRepo(db *gorm.DB) usecase.PresentationRepo {
	return &presentationRepo{
		db: db,
	}
}

func (r presentationRepo) CreatePresentation(request *entity.Presentation) (uint32, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return 0, err
	}
	return request.ID, nil
}

func (r presentationRepo) GetPresentation(id uint32) (*entity.Presentation, error) {
	presentation := &entity.Presentation{ID: id}
	if err := r.db.Preload("Slides").First(&presentation).Error; err != nil {
		return nil, err
	}
	return presentation, nil
}

func (r presentationRepo) Collection(groupId uint32) ([]*entity.Presentation, error) {
	presentations := []*entity.Presentation{}
	if err := r.db.Where("group_id=?", groupId).Find(&presentations).Error; err != nil {
		return nil, err
	}
	return presentations, nil
}

func (r presentationRepo) UpdatePresentation(request *entity.Presentation) error {
	if err := r.db.Save(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r presentationRepo) DeletePresentation(id uint32, userId uint32) error {

	if err := r.db.Where("user_id=?", userId).Delete(&entity.Presentation{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
