package models

type User struct {
	ID       uint   `gorm:"column:id;primarykey" json:"id_user"`
	Username string `json:"username" gorm:"unique;not null;index"`
	Password []byte `json:"-" gorm:"not null"`
}
