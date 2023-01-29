package mnews

import (
	"caker/models/mmitra"
	"time"

	"github.com/go-playground/validator/v10"
)

type News struct {
	ID        uint `json:"id"`
	MitraID   uint `json:"mitra_id"`
	Mitra     mmitra.Mitra
	Title     string    `json:"title"`
	Blog      string    `json:"blog"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputNews struct {
	MitraID uint   `json:"mitra_id" validate:"required"`
	Title   string `json:"title"`
	Blog    string `json:"blog"`
	Photo   string `json:"photo"`
}

func (input *InputNews) ValidNews() error {
	v := validator.New()
	ere := v.Struct(input)
	return ere
}

type UpdateNews struct {
	Title string `json:"title"`
	Blog  string `json:"blog"`
	Photo string `json:"photo"`
}
