package model

import "gorm.io/gorm"

type Novel struct {
	gorm.Model
	Name        string
	Description string
	Author      string
}
