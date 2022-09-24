package authdto

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" `
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Status   string `gorm:"type: varchar(255)" json:"status"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Status   string `gorm:"type: varchar(255)" json:"status" `
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
