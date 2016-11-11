package dao

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"esvodsApi/forms"
	"esvodsCore/models"
)

//WatcherDao ...
type WatcherDao struct{}

//Signin ...
func (d WatcherDao) Signin(form forms.SigninForm) (watcher models.Watcher, err error) {
	err = GetDB().Where(&models.Watcher{Email: form.Email}).First(&watcher).Error

	if watcher.ID != 0 && err == nil {

		bytePassword := []byte(form.Password)
		byteHashedPassword := []byte(watcher.Password)
		err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
		checkErr(err, "Invalid password")

		return watcher, nil
	}

	return watcher, errors.New("Create an account")
}

//Signup ...
func (d WatcherDao) Signup(form forms.SignupForm) (watcher models.Watcher, err error) {
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
	err = GetDB().First(&watcher, watcherID).Error
	return watcher, err
}
