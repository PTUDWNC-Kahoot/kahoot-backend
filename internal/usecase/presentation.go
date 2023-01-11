package usecase

import (
	"examples/kahootee/internal/entity"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

type presentation struct {
	repo PresentationRepo
}

func NewPresentation(repo PresentationRepo) Presentation {
	return &presentation{
		repo: repo,
	}
}

func (r presentation) CreatePresentation(request *entity.Presentation) error {
	code := genCode()
	request.Code = code
	return r.repo.CreatePresentation(request)
}

func (r presentation) GetPresentation(id uint32) (*entity.Presentation, error) {
	return r.repo.GetPresentation(id)
}

func (r presentation) GetPresentationByCode(code string) (*entity.Presentation, error) {
	return r.repo.GetPresentationByCode(code)
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

func (r presentation) AddCollaborator(emails []string, presentationID uint32) error {
	return r.repo.AddCollaborator(emails, presentationID)
}

func (r presentation) GetCollaborators(presentationID uint32) ([]*entity.Collaborator, error) {
	return r.repo.GetCollaborators(presentationID)
}

func (r presentation) RemoveCollaborator(userID, presentationID uint32) error {
	return r.repo.RemoveCollaborator(userID, presentationID)
}

func genCode() string {
	joinCode := uuid.New()
	code := encode(joinCode)
	fmt.Println("Room code:", code)
	return code
}
func encode(u uuid.UUID) string {
	return new(big.Int).SetBytes(u[:]).Text(62)
}
