package mtraining

import (
	"caker/models/mmitra"
	"time"

	"github.com/go-playground/validator/v10"
)

type Training struct {
	ID        uint `json:"id"`
	MitraID   uint `json:"mitra_id"`
	Mitra     mmitra.Mitra
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo     string    `json:"photo"`
	Duration  string    `json:"duration"`
	Cost      string    `json:"cost"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputTraining struct {
	MitraID  uint   `json:"mitra_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	Photo    string `json:"photo"`
	Duration string `json:"duration"`
	Cost     string `json:"cost"`
	Link     string `json:"link"`
}

func (input *InputTraining) ValidTraining() error {
	v := validator.New()
	err := v.Struct(input)
	return err
}

type UpdateTraining struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Photo    string `json:"photo"`
	Duration string `json:"duration"`
	Cost     string `json:"cost"`
	Link     string `json:"link"`
}
