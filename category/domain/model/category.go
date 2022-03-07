/*
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-02-10 12:57:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-03-06 11:50:41
 */
package model

type Category struct {
	ID                  int64  `gorm:"primary_key;not_null;auto_increment" json:"id"` //主键ID
	CategoryName        string `gorm:"unique_index,not_null" json:"category_name"`  //分类名称
	CategoryLevel       uint32 `json:"category_level"`  //分类等级
	CategoryParent      int64  `json:"category_parent"`
	CategoryImage       string `json:"category_image"`
	CategoryDescription string `json:"category_description"`  //分类描述
}
