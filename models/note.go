package models

import "gorm.io/gorm"

type Note struct {
    gorm.Model
    Title     string
    Content   string
    UserID    uint
    SharedWithUsers []User `gorm:"many2many:note_shared_users;"`
}
