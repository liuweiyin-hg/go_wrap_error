package user_app

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Salt      string `gorm:"type:varchar(255)" json:"salt"`
	Username  string `gorm:"type:varchar(32);column:username" json:"username"`
	Password  string `gorm:"type:varchar(200);column:password" json:"password"`
	Languages string `gorm:"type:varchar(200);column:languages" json:"languages"`
	Number    int    `gorm:"type:int" json:"number"`
}

func (u User) TableName() string {
	return "gorm_user"
}

type UserSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Salt      string    `json:"salt"`
	UserName  string    `json:"username"`
	Password  string    `json:"-"`
	Languages string    `json:"languages"`
	Number    int       `json:"number"`
}

func (self User) Serializer() UserSerializer {
	return UserSerializer{
		ID:        self.ID,
		CreatedAt: self.CreatedAt.Truncate(time.Second),
		UpdatedAt: self.UpdatedAt.Truncate(time.Second),
		Salt:      self.Salt,
		Password:  self.Password,
		Languages: self.Languages,
		UserName:  self.Username,
		Number:    self.Number}
}
