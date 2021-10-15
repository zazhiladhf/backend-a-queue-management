package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id_user"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}