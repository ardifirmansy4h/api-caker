package mmitra

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Mitra struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Photo      string    `json:"photo"`
	Nohp       string    `json:"no_hp"`
	Address    string    `json:"address"`
	Founder    string    `json:"founder"`
	CreatedAt  time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Photo    string `json:"photo"`
	Nohp     string `json:"no_hp"`
	Address  string `json:"address"`
	Founder  string `json:"founder"`
}

func (input *Register) ValidRegister() error {
	v := validator.New()
	err := v.Struct(input)
	return err
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
