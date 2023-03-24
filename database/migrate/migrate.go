package main

func main() {
	// db.Con.AutoMigrate(&models.User{})
	// db.Con.AutoMigrate(&models.Articles{})
	// db.Con.AutoMigrate(&models.Category{})
	// db.Con.AutoMigrate(&models.Photos{})
	// db.Con.AutoMigrate(&models.Photo_category{})

	//ARTICLE CATEGORIES
	// db.Con.Model(&models.Category{}).Create([]map[string]interface{}{
	// 	{"CATEGORY_ICON": "https://cdn-icons-png.flaticon.com/128/3137/3137927.png", "CATEGORY_NAME": "Golang DSA"},
	// 	{"CATEGORY_ICON": "https://miro.medium.com/max/559/1*oZ1j-s22SCUMZamIyVeQtQ.jpeg", "CATEGORY_NAME": "Golang Development"},
	// 	{"CATEGORY_ICON": "https://miro.medium.com/max/512/1*t2ceb9fkddzb9UTrruGgWw.png", "CATEGORY_NAME": "Problems Solution"},
	// })

	//ARTICLE CATEGORIES
	// db.Con.Model(&models.Photo_category{}).Create([]map[string]interface{}{
	// 	{"CATEGORY_ICON": "https://wallpapers.com/images/featured/2ygv7ssy2k0lxlzu.jpg", "CATEGORY_NAME": "Nature"},
	// 	{"CATEGORY_ICON": "https://cdn.mos.cms.futurecdn.net/2gHPhDWjds5q8nqLM2FG9Y.jpg", "CATEGORY_NAME": "People"},
	// 	{"CATEGORY_ICON": "https://cdn.telanganatoday.com/wp-content/uploads/2021/12/Numaish-1.jpg", "CATEGORY_NAME": "Exhibition"},
	// })
}
