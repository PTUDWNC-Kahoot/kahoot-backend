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

func (r presentationRepo) CreatePresentation(request *entity.Presentation) error {
	return r.db.Create(&request).Error
}

func (r presentationRepo) GetPresentation(id uint32) (*entity.Presentation, error) {
	presentation := &entity.Presentation{ID: id}
	if err := r.db.Preload("Slides").Preload("Slides.Options").Preload("Collaborators").First(&presentation).Error; err != nil {
		return nil, err
	}
	return presentation, nil
}

func (r presentationRepo) GetPresentationByCode(code string) (*entity.Presentation, error) {
	presentation := &entity.Presentation{}
	if err := r.db.Preload("Slides").Preload("Slides.Options").Preload("Collaborators").Where("code=?",code).First(&presentation).Error; err != nil {
		return nil, err
	}
	return presentation, nil
}

func (r presentationRepo) GroupCollection(groupId uint32) ([]*entity.Presentation, error) {
	presentations := []*entity.Presentation{}
	if err := r.db.Where("group_id=?", groupId).Find(&presentations).Error; err != nil {
		return nil, err
	}
	return presentations, nil
}

func (r presentationRepo) MyCollection(userId uint32) ([]*entity.Presentation, error) {
	presentations := []*entity.Presentation{}
	if err := r.db.Where("user_id=?", userId).Find(&presentations).Error; err != nil {
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

func (r presentationRepo) CreateSlide(slide *entity.Slide) error {
	return r.db.Create(&slide).Error
}

func (r presentationRepo) UpdateSlide(slide *entity.Slide) error {
	return r.db.Debug().Updates(&slide).Error
}

func (r presentationRepo) DeleteSlide(id uint32) error {
	return r.db.Delete(&entity.Slide{ID: id}).Error
}

func (r presentationRepo) AddCollaborator(emails []string, presentationID uint32) error {
	users := []*entity.User{}

	for _, email := range emails {
		user := &entity.User{}
		err := r.db.Where("email=?", email).First(&user).Error
		if err != nil {
			continue
		}

		existed := &entity.Collaborator{}
		r.db.Where("user_id=?", user.ID).Where("presentation_id=?", presentationID).First(existed)
		if existed.UserID != 0 {
			continue
		}

		users = append(users, user)
	}

	collaborators := []*entity.Collaborator{}

	for _, user := range users {
		c := &entity.Collaborator{
			UserID:         user.ID,
			Name:           user.Name,
			PresentationID: presentationID,
		}

		collaborators = append(collaborators, c)
	}

	return r.db.Create(&collaborators).Error
}

func (r presentationRepo) GetCollaborators(presentationID uint32) ([]*entity.Collaborator, error) {
	collaborators := []*entity.Collaborator{}
	if err := r.db.Where("presentation_id=?", presentationID).Find(&collaborators).Error; err != nil {
		return nil, err
	}
	return collaborators, nil
}

func (r presentationRepo) RemoveCollaborator(userID, presentationID uint32) error {
	return r.db.Where("user_id=? AND presentation_id=?", userID, presentationID).Delete(&entity.Collaborator{}).Error
}
