package model

import "time"

type User struct {
	ID				int64		`json:"id"`
	Username		string		`json:"username"`
	HashedPassword	string		`json:"-"`
	CreatedAt		time.Time	`json:"created_at"`
}

type Todo struct {
	ID			int64		`json:"id"`
	UserID		int64		`json:"user_id"`
	Task		string		`json:"task"`
	Completed	bool		`json:"completed"`
	CreatedAt	time.Time	`json:"created_at"`
	
}



