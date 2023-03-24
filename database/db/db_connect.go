package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://github.com/go-sql-driver/mysql
var dsn = "root:password@tcp(mysql_name)/mushahary_com?charset=utf8&parseTime=True&loc=Local"
var Con, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
