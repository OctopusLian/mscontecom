/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-09 22:12:30
 * @LastEditors: neozhang
 * @LastEditTime: 2022-03-06 11:18:47
 */
package repository

import (
	"user/domain/model"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	//根据用户ID查找用户信息
	FindUserByID(int64) (*model.User, error)
	//创建用户
	CreateUser(*model.User) (int64, error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有用户
	FindAll() ([]model.User, error)
}

//创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

//根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

//根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

//创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

//根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

//更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

//查找所有用户
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
