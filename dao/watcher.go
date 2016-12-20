package dao

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"vodstv/core/models"
)

//SigninForm ...
type SigninForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

//SignupForm ...
type SignupForm struct {
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

//WatcherDao ...
type WatcherDao struct{}

//Signin ...
func (d WatcherDao) Signin(form SigninForm) (watcher models.Watcher, err error) {
	err = GetDB().Where(&models.Watcher{Email: form.Email}).First(&watcher).Error

	if watcher.ID != 0 && err == nil {

		bytePassword := []byte(form.Password)
		byteHashedPassword := []byte(watcher.Password)
		if fail := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword); fail != nil {
			err = errors.New("Invalid password")
		}

		return watcher, err
	}

	return watcher, errors.New("Create an account")
}

//Signup ...
func (d WatcherDao) Signup(form SignupForm) (watcher models.Watcher, err error) {
	err = GetDB().Where(&models.Watcher{Email: form.Email}).First(&watcher).Error

	if GetDB().NewRecord(watcher) && err == nil {
		watcher.Email = form.Email
		watcher.Name = form.Name
		watcher.Password = form.Password
		err = GetDB().Create(&watcher).Error

		return watcher, err
	}

	return watcher, errors.New("Watcher exists")
}

//Get ...
func (d WatcherDao) Get(watcherID uint) (watcher models.Watcher, err error) {
	err = GetDB().First(&watcher, watcherID).Related(&watcher.Feeds).Error
	return watcher, err
}

//Find ...
func (d WatcherDao) Find(q map[string]interface{}) (watchers []models.Watcher, err error) {
	err = GetDB().Where(q).Find(&watchers).Error
	return watchers, err
}

//Save ...
func (d WatcherDao) Save(watcher *models.Watcher) (err error) {
	if GetDB().NewRecord(watcher) {
		err = GetDB().Create(&watcher).Error
	} else {
		err = GetDB().Save(&watcher).Updates(getUpdates(watcher)).Error
	}

	return err
}

//Delete ...
func (d WatcherDao) Delete(id uint) error {
	watcher, err := d.Get(id)
	if err == nil {
		err = GetDB().Delete(&watcher).Error
	}
	return err
}
