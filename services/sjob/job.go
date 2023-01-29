package sjob

import (
	"caker/models/mjob"
	"caker/repository"
	"caker/repository/rjob"
)

type JobService struct {
	repo repository.JobRepo
}

func NewJobRepo() JobService {
	return JobService{
		repo: &rjob.JobRepoIMPL{},
	}
}

func (js *JobService) GetAllJob() []mjob.Job {
	return js.repo.GetAllJob()
}
func (js *JobService) GetByIDJob(id string) mjob.Job {
	return js.repo.GetByIDJob(id)
}
func (js *JobService) CreateJob(input mjob.InputJob) mjob.Job {
	return js.repo.CreateJob(input)
}
func (js *JobService) UpdateJob(id string, input mjob.InputJob) mjob.Job {
	return js.repo.UpdateJob(id, input)
}
func (js *JobService) DeleteJob(id string) bool {
	return js.repo.DeleteJob(id)
}
