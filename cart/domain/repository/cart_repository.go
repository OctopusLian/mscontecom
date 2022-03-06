/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-10 13:32:01
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-10 20:03:25
 */
package repository

import (
	"cart/domain/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

//创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

//根据ID查找Cart信息
func (u *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDb.First(cart, cartID).Error
}

//创建Cart信息
func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

//根据ID删除Cart信息
func (u *CartRepository) DeleteCartByID(cartID int64) error {
	return u.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

//更新Cart信息
func (u *CartRepository) UpdateCart(cart *model.Cart) error {
	return u.mysqlDb.Model(cart).Update(cart).Error
}

//获取结果集
func (u *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDb.Where("user_id = ?", userID).Find(&cartAll).Error
}

//根据用户ID清空购物车
func (u *CartRepository) CleanCart(userID int64) error {
	return u.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

//添加商品数量
func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

//购物车减少商品
func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
