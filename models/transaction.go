package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string       `json:"startDate"`
	DueDate   string       `json:"dueDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attache   string       `json:"attache"`
	Price     int          `json:"price"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID        		int       		`json:"id"`
	StartDate      	string    		`json:"startDate"`
	DueDate      		string    		`json:"dueDate"`
	UserID      		int    			`json:"user_id"`
	User      		UserResponse    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attache     		string    		`json:"attache"`
	Status     		string    		`json:"status"`
	CreatedAt 		time.Time 		`json:"-"`
	UpdatedAt 		time.Time 		`json:"-"`
}

type TransactionUserResponse struct {
	ID        int    `json:"id"`
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id"`
	Attache   string `json:"attache"`
	Status    bool   `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}
