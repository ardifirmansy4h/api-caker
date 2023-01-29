package msocmeds

import (
	"caker/models/musers"
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Caption   string `json:"caption"`
	Photo     string `json:"photo"`
	User      musers.User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputSocmed struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Caption string `json:"caption"`
	Photo   string `json:"photo"`
}

func (input *InputSocmed) ValidSocmed() error {
	v := validator.New()
	err := v.Struct(input)
	return err
}
