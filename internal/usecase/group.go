package usecase

import (
	"examples/kahootee/internal/entity"
	"math/big"

	"github.com/google/uuid"
)

const FE_HOST = "http://localhost:3000/join-group/"

type groupUsecase struct {
	repo GroupRepo
}

func NewGroupUsecase(repo GroupRepo) GroupUsecase {
	return &groupUsecase{
		repo: repo,
	}
}

func (g *groupUsecase) GetGroups() ([]*entity.Group, error) {
	return g.repo.Collection()
}

func (g *groupUsecase) Get(id uint32) (*entity.Group, error) {
	return g.repo.GetOne(id)
}

func (g *groupUsecase) Create(request *entity.Group) (uint32, error) {
	inviteCode := uuid.New()
	request.InvitationLink = FE_HOST + encode(inviteCode)
	return g.repo.CreateOne(request)
}

func (g *groupUsecase) Update(request *entity.Group) error {
	return g.repo.UpdateOne(request)
}

func (g *groupUsecase) Delete(id uint32) error {
	return g.repo.DeleteOne(id)
}

func encode(u uuid.UUID) string {
	return new(big.Int).SetBytes(u[:]).Text(62)
}

func (g *groupUsecase) JoinGroupByLink(userEmail string, groupCode string) (*entity.Group, error) {
	return g.repo.JoinGroupByLink(userEmail, groupCode)
}

func (g *groupUsecase) Invite(email_list []string, groupID uint32) error {
	return g.repo.Invite(email_list, groupID)
}

func (g *groupUsecase) AssignRole(groupUser *entity.GroupUser, ownerEmail string) error {
	return g.repo.AssignRole(groupUser, ownerEmail)
}
