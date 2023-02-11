package entity

import "time"

type Model struct {
	Id         int       `json:"id" gorm:"column:id"`
	CreateTime time.Time `gorm:"column:createtime;default:null" json:"createtime"`
	UpdateTime time.Time `gorm:"column:updatetime;default:null" json:"updatetime"`
}
