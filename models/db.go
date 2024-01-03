package models

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB


func ConnectDatabase() {
    var err error

    
    DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    
    err = DB.AutoMigrate(&User{}, &Note{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}
