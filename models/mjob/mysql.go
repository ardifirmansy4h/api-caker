package mjob

import (
	"caker/models/mmitra"
	"time"

	"github.com/go-playground/validator/v10"
)

type Job struct {
	ID          uint `json:"id"`
	MitraID     uint `json:"mitra_id"`
	Mitra       mmitra.Mitra
	Title       string    `json:"title"`
	Logo        string    `json:"logo"`
	Description string    `json:"description"`
	Requirement string    `json:"requirement"`
	Skill       string    `json:"skill"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	NoHp        string    `json:"no_hp"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InputJob struct {
	MitraID     uint   `json:"mitra_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	Skill       string `json:"skill"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	NoHp        string `json:"no_hp"`
	Link        string `json:"link"`
}

func (input *InputJob) ValidJob() error {
	v := validator.New()
	err := v.Struct(input)
	return err
}
