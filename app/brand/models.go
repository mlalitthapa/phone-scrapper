package brand

type Brand struct {
	//gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Slug string `json:"slug" gorm:"type:varchar(100);unique_index"`

	Devices string `json:"devices" gorm:"-"`
}

type deviceImage struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type device struct {
	Name  string      `json:"name"`
	Slug  string      `json:"slug"`
	Image deviceImage `json:"image"`
}
