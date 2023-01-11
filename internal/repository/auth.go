package repo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"

	"examples/kahootee/internal/entity"
	"examples/kahootee/internal/usecase"
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

func (repo *authRepo) Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Presentation, error) {
	user := &entity.User{}
	encryptedPass := getMD5Hash(request.Password)
	err := repo.db.Where("email=? and password=?", request.Email, encryptedPass).First(user).Error
	if err != nil {
		return nil, nil, nil, err
	}
	groups := []*entity.Group{}
	presentations := []*entity.Presentation{}
	if err := repo.db.Debug().Model(&entity.Group{}).Joins("left join group_users on group_users.group_id=groups.id").Where("group_users.user_id=?", user.ID).Scan(&groups).Error; err != nil {
		return nil, nil, nil, err
	}
	if err := repo.db.Debug().Model(&presentations).Where("owner=?", user.ID).Scan(&presentations).Error; err != nil {
		return nil, nil, nil, err
	}

	return user, groups, presentations, nil
}

func (repo *authRepo) Register(request *entity.User) error {
	user := &entity.User{}
	encryptedPass := getMD5Hash(request.Password)
	return repo.db.Debug().Create(&entity.User{Email: request.Email, Password: encryptedPass, Name: "kahoot_user", CoverImageURL: defaultAvatar}).Scan(user).Error
}

func (repo *authRepo) CreateRegisterOrder(request *entity.RegisterOrder) (uint32, error) {
	request.ExpiresAt = time.Now().Add(time.Minute * 5)
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
