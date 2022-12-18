package repo

import (
	"crypto/md5"
	"encoding/hex"
	"examples/kahootee/internal/entity"
	"examples/kahootee/internal/usecase"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const defaultAvatar = "https://i.pinimg.com/564x/ec/18/a3/ec18a302c5672470c894939f2cc1a830.jpg"

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) usecase.AuthRepo {
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

func (repo *authRepo) Register(request *entity.User) error {
	user := &entity.User{}
	encryptedPass := getMD5Hash(request.Password)
	kh := entity.Kahoot{ID: 1}
	repo.db.Debug().Where("ID=?", kh.ID).First(&kh)
	return repo.db.Debug().Create(&entity.User{Email: request.Email, Password: encryptedPass, Name: "kahoot_user", CoverImageURL: defaultAvatar}).Scan(user).Error
}

func (repo *authRepo) CreateRegisterOrder(request *entity.RegisterOrder) (uint32, error) {
	request.ExpiresAt = time.Now().Add(time.Minute * 2)
	err := repo.db.Debug().Create(request).Error
	if err != nil {
		return 0, err
	}
	return request.ID, nil
}

func (repo *authRepo) VerifyEmail(email string, verifyCode int) bool {
	order := &entity.RegisterOrder{}
	err := repo.db.Where("email=? and verify_code=?", email, verifyCode).First(order).Error
	if err != nil || order.ID == 0 || order.ExpiresAt.Before(time.Now()) {
		return false
	}
	err = repo.db.Delete(order).Error
	if err != nil {
		fmt.Println("Delete register order failed")
	}
	return true
}

func (repo *authRepo) CheckEmailExisted(email string) bool {
	user := &entity.User{}
	err := repo.db.Debug().Where("email=?", email).First(user).Error
	if err != nil || user.ID == 0 {
		return false
	}
	return true
}
