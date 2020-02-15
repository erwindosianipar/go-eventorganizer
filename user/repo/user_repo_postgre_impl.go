package repo

import (
	"errors"

	"eventorganizer/golang/models"
	"eventorganizer/golang/user"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var _ user.UserRepo = &UserRepoPostgreImpl{}

type UserRepoPostgreImpl struct {
	db *gorm.DB
}

func CreateUserRepoPostgreImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoPostgreImpl{db}
}

func (h *UserRepoPostgreImpl) Register(user *models.User) (*models.User, error)  {
	if err := h.db.Table("users").Save(&user).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: insert data user")
	}
	
	return user, nil
}

func (h *UserRepoPostgreImpl) GetAllUser() ([]*models.User, error) {
	userList := make([]*models.User, 0)

	if err := h.db.Table("users").Find(&userList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: get all data users")
	}

	return userList, nil
}

func (h *UserRepoPostgreImpl) GetUserByID(id int) (*models.User, error) {
	dataUser := new(models.User)

	if err := h.db.Table("users").Where("id = ?", id).First(&dataUser).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: get data user by id")
	}

	return dataUser, nil
}

func (h *UserRepoPostgreImpl) DeleteUser(id int) (*models.User, error)  {
	if err := h.db.Table("users").Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: delete data user")
	}

	return nil, nil
}

func (h *UserRepoPostgreImpl) IsAnyEmailUser(email string) bool {
	var total int
	h.db.Table("users").Where("email = ?", email).Count(&total)

	if total > 0 {
		return true
	}

	return false
}

func (h *UserRepoPostgreImpl) GetUserByEmail(email string) (*models.User, error) {
	dataUser := new(models.User)

	if err := h.db.Table("users").Where("email = ?", email).First(&dataUser).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ERROR: get data user by email")
	}

	return dataUser, nil
}

func (h *UserRepoPostgreImpl) UpgradeUser(user *models.User) (*models.User, error) {
	h.db.Table("users").Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name_eo" : user.EventOrganizer.NameEo,
		"ktp_number" : user.EventOrganizer.KTPNumber,
		"siup_number" : user.EventOrganizer.SIUPNumber,
	})

	//h.db.Model(&user).Updates(map[string]interface{}{"name_eo" : nameEo, "ktp_number" : ktpNumber, "siup_number" : siupNumber})

	return user, nil
}
