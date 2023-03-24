package models

import "os"

type Photos struct {
	ID                uint
	USER_ID           string
	PHOTO_CATEGORY_ID int64
	PHOTO_FILE        string
	PHOTO_TITLE       string
	PHOTO_DES         string
	UPLOADED_DATE     int64 `gorm:"autoCreateTime"`
	UPDATED_DATE      int64 `gorm:"autoUpdateTime"`
}

type Photo_category struct {
	ID            uint
	CATEGORY_ICON string
	CATEGORY_NAME string
}

type JoinPhotosList struct {
	SL                int64
	ID                uint
	USER_ID           string
	PHOTO_CATEGORY_ID int64
	PHOTO_FILE        string
	PHOTO_TITLE       string
	PHOTO_DES         string
	UPLOADED_DATE     int64
	UPDATED_DATE      int64
	CATEGORY_ICON     string
	CATEGORY_NAME     string
}

type ForCreatePhotos struct {
	ID                uint
	USER_ID           string
	PHOTO_CATEGORY_ID int64
	PHOTO_FILE        *os.File
	PHOTO_TITLE       string
	PHOTO_DES         string
	UPLOADED_DATE     int64
	UPDATED_DATE      int64
}
