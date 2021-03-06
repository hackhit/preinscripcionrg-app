package student

import "time"

//Teacher Profesor de la aplicacion.
type Teacher struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID      int    `json:"userID" gorm:"type:integer"`
	FirstName   string `json:"firstName" gorm:"not null;type:varchar(30)"`
	LastName    string `json:"lastName" gorm:"not null;type:varchar(30)"`
	Email       string `json:"email" gorm:"not null;unique;type:varchar(40)"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null;type:varchar(15)"`
	CI          int    `json:"ci" gorm:"not null;unique;type:varchar(15)"`

	Message string `json:"message" gorm:"-"`
	Code    int    `json:"code" gorm:"-"`
}
