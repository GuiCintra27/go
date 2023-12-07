package schemas

import "github.com/jinzhu/gorm"

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary	int64
}