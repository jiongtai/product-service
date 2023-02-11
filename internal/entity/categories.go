package entity

type Categories struct {
	Model
	Name string `json:"name" gorm:"column:name"`
}
