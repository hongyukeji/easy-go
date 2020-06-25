package models

import (
	"github.com/jinzhu/gorm"
	"time"

	orm "github.com/hongyukeji/easy-go/datasource"
)

func init() {
	orm.Eloquent.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Salt      string `gorm:"type:varchar(255)" json:"salt"`
	Username  string `gorm:"type:varchar(32)" json:"username"`
	Password  string `gorm:"type:varchar(200);column:password" json:"-"`
	Languages string `gorm:"type:varchar(200);column:languages" json:"languages"`
}

func (u User) TableName() string {
	return "users"
}

type UserSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Salt      string    `json:"salt"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Languages string    `json:"languages"`
}

func (self User) Serializer() UserSerializer {
	return UserSerializer{
		ID:        self.ID,
		CreatedAt: self.CreatedAt.Truncate(time.Second),
		UpdatedAt: self.UpdatedAt.Truncate(time.Second),
		Salt:      self.Salt,
		Password:  self.Password,
		Languages: self.Languages,
		Username:  self.Username,
	}
}
