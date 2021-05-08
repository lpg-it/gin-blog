package model

import (
	"fmt"
	"gin-blog/utils"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init(){
	// 链接数据库
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"))
}
