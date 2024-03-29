package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `grom:"unique"`
	Password string
}

// equals
// type User struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
//   Name string
// }
