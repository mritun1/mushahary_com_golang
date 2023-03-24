package models

type Articles struct {
	ID          uint
	USER_ID     string
	THUMBNAIL   string
	CATEGORY_ID int64
	TITLE       string
	DES         string
	CONTENT     string
	CREATE_DATE int64 `gorm:"autoCreateTime"`
	UPDATE_DATE int64 `gorm:"autoUpdateTime"`
}

type Category struct {
	ID            uint
	CATEGORY_ICON string
	CATEGORY_NAME string
}

type JoinArticleList struct {
	SL            int64
	ID            uint
	USER_ID       string
	THUMBNAIL     string
	CATEGORY_ID   int64
	TITLE         string
	DES           string
	CONTENT       string
	CREATE_DATE   int64
	UPDATE_DATE   int64
	CATEGORY_ICON string
	CATEGORY_NAME string
}
