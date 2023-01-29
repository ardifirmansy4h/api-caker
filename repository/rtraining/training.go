package rtraining

import (
	"caker/app/config"
	"caker/models/mtraining"
)

type TrainingRepo struct{}

func (tr *TrainingRepo) GetAllTraining() []mtraining.Training {
	var training []mtraining.Training
	config.DB.Preload("Mitra").Find(&training)
	return training
}

func (tr *TrainingRepo) GetByIDTraining(id string) mtraining.Training {
	var training mtraining.Training
	config.DB.First(&training, "id = ?", id).Preload("Mitra").Find(&training)
	return training
}

func (tr *TrainingRepo) CreateTraining(input mtraining.InputTraining) mtraining.Training {
	var training mtraining.Training = mtraining.Training{
		MitraID:  input.MitraID,
		Title:    input.Title,
		Caption:  input.Caption,
		Photo:    input.Photo,
		Duration: input.Duration,
		Cost:     input.Cost,
		Link:     input.Link,
	}
	res := config.DB.Create(&training).Preload("Mitra").Find(&training)
	res.Last(&training)
	return training
}

func (tr *TrainingRepo) UpdateTraining(id string, input mtraining.InputTraining) mtraining.Training {
	var training mtraining.Training = tr.GetByIDTraining(id)
	training.Title = input.Title
	training.Caption = input.Caption
	training.Photo = input.Photo
	training.Duration = input.Duration
	training.Cost = input.Cost

	res := config.DB.Save(&training).Preload("Mitra").Find(&training)
	res.Last(&training)
	return training
}

func (tr *TrainingRepo) DeleteTraining(id string) bool {
	var training mtraining.Training = tr.GetByIDTraining(id)
	res := config.DB.Delete(&training)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
