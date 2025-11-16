package storage


import {
	
    "github.com/CatKap/sixPsyh/models"
	  "gorm.io/driver/sqlite"
    "gorm.io/gorm"	
}

type Storage struct{
  db   *sql.DB
  log  *loger.Loger
	meetings *sql.Statement,
}


func New(db *DB.sql, log *loger.Loger) *Storage{
	
}
