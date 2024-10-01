// models/user.go
package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
	Age  string `gorm:"size:100"`
}
