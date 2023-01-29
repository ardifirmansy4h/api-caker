package musers

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NoHp      string    `json:"no_hp"`
	Gender    string    `json:"gender"`
	Photo     string    `json:"photo"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Register struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoHp     string `json:"no_hp"`
	Gender   string `json:"gender"`
	Photo    string `json:"photo"`
	Address  string `json:"address"`
}

func (input *Register) ValidInputRegister() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
