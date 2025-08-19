package models

type Users struct {
	ID   int     `gorm:"primaryKey;autoIncrement" json:"id"`
	SID  string  `gorm:"column:sid" json:"sid"`
	Name string  `json:"name"`
	CGPA float32 `json:"cgpa"`
}
