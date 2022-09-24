package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Subscribe string    `json:"subscribe" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"type: varchar(255)"`
	Token     string    `json:"token" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Subscribe string    `json:"subscribe" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"type: varchar(255)"`
	Token     string    `json:"token" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAuthResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Subscribe string    `json:"subscribe" gorm:"type:varchar(50)"`
	Status    string    `gorm:"type: varchar(255)" json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserAuthResponse) TableName() string {
	return "users"
}
