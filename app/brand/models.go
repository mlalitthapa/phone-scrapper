package brand

type Brand struct {
	//gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Slug string `json:"slug" gorm:"type:varchar(100);unique_index"`

	Devices string `json:"devices" gorm:"-"`
}
