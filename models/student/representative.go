package student

import "time"

//Representative Representante de la aplicacion.
type Representative struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID       int    `json:"userID" gorm:"type:integer"`
	FirstName    string `json:"firstName" gorm:"not null;type:varchar(30)"`
	LastName     string `json:"lastName" gorm:"not null;type:varchar(30)"`
	PhoneNumber  string `json:"phoneNumber" gorm:"not null;type:varchar(15)"`
	CI           string `json:"ci" gorm:"not null;unique;type:varchar(15)"`
	Relationship string `json:"relationship" gorm:"not null;type:varchar(30)"`
	Address      string `json:"address" gorm:"not null;type:varchar(30)"`

	Message string `json:"message" gorm:"-"`
	Code    int    `json:"code" gorm:"-"`
}
