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

func (db *authRepo) Login(request *entity.User) bool {
	var result entity.User
	encryptedPass := getMD5Hash(request.Password)
	err := db.db.Where("username=? and password=?", request.Email, encryptedPass).First(&result).Error
	if err != nil {
		return false
	}
	return true
}
func (db *authRepo) Register(request *entity.User) bool {
	var result entity.User
	encryptedPass := getMD5Hash(request.Password)
	err := db.db.Create(&entity.User{Id: request.Id, Email: request.Email, Password: encryptedPass}).Scan(&result).Error
	if err != nil {
		return false
	}
	return true
}
