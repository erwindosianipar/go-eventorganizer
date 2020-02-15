package models

type User struct {
	OrmModel
	Email       string `gorm:"email" json:"email,omitempty"`
	Password    string `gorm:"password" json:"password,omitempty"`
	Name        string `gorm:"name" json:"name,omitempty"`
	Avatar      string `gorm:"avatar" json:"avatar,omitempty"`
	Role        string `gorm:"role" sql:"DEFAULT:'user'" json:"role,omitempty"`
	EventOrganizer `json:"event_organizer,omitempty"`
}
