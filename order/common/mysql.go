/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-03-03 22:38:41
 * @LastEditors: neozhang
 * @LastEditTime: 2022-03-03 22:41:37
 */
package common

import "github.com/micro/go-micro/v2/config"

//创建结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	//获取配置
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
