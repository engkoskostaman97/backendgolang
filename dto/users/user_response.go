package usersdto

type UserResponse struct {
	FullName  string `json:"fullname" form:"fullname" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Subscribe string `json:"subscribe" form:"subscribe" validate:"required"`
	Status    string `gorm:"type: varchar(255)" json:"status" validate:"required"`
}
