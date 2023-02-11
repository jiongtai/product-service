package entity

type Brands struct {
	Model
	Name string `json:"name" gorm:"column:name"`
}
