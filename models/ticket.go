package models

import (
	"time"
)

type Ticket struct
{
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Title string
	Description string 
	User_id int
	Department string
}