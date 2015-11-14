package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Ticket struct
{
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Title string
	Description string `sql:"type:text"`
}