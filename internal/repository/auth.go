package repo

import (
	"crypto/md5"
	"encoding/hex"
	"examples/identity/internal/entity"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (repo *authRepo) Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Kahoot, error) {
	user := &entity.User{}
	encryptedPass := getMD5Hash(request.Password)
	err := repo.db.Where("email=? and password=?", request.Email, encryptedPass).First(user).Error
	if err != nil {
		return nil, nil, nil, err
	}
	groups := []*entity.Group{}
	kahoots := []*entity.Kahoot{}

	repo.db.Model(user).Association("Groups").Find(&groups)
	repo.db.Model(user).Association("Kahoots").Find(&kahoots)

	return user, groups, kahoots, nil
}

func (repo *authRepo) Register(request *entity.User) (uint32, error) {
	user := &entity.User{}
	encryptedPass := getMD5Hash(request.Password)
	kh := entity.Kahoot{ID: 1}
	repo.db.Debug().Where("ID=?", kh.ID).First(&kh)
	err := repo.db.Debug().Create(&entity.User{Email: request.Email, Password: encryptedPass}).Scan(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
