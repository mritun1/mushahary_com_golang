package models

type User struct {
	ID              uint
	AVATAR          string
	USER_NAME       string
	PASSWORD        string
	FIRST_NAME      string
	LAST_NAME       string
	JOIN_DATE       int64 `gorm:"autoCreateTime"`
	LAST_LOGIN_DATE int64 `gorm:"autoUpdateTime"`
}
