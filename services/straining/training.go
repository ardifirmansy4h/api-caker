package straining

import (
	"caker/models/mtraining"
	"caker/repository"
	"caker/repository/rtraining"
)

type TrainingService struct {
	repo repository.TrainingRepo
}

func NewTrainingService() TrainingService {
	return TrainingService{
		repo: &rtraining.TrainingRepo{},
	}
}

func (ts *TrainingService) GetAllTraining() []mtraining.Training {
	return ts.repo.GetAllTraining()
}

func (ts *TrainingService) GetByIDTraining(id string) mtraining.Training {
	return ts.repo.GetByIDTraining(id)
}

func (ts *TrainingService) CreateTraining(input mtraining.InputTraining) mtraining.Training {
	return ts.repo.CreateTraining(input)
}

func (ts *TrainingService) UpdateTraining(id string, input mtraining.InputTraining) mtraining.Training {
	return ts.repo.UpdateTraining(id, input)
}

func (ts *TrainingService) DeleteTraining(id string) bool {
	return ts.repo.DeleteTraining(id)
}
