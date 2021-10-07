package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
	SlotBooking []*SlotBooking
}