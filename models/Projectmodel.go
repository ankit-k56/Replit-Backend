package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name"`
	Tech string `json:"tech"`
	UserID uint `json:"user_id"`
}