package rjob

import (
	"caker/app/config"
	"caker/models/mjob"
)

type JobRepoIMPL struct{}

func (jr *JobRepoIMPL) GetAllJob() []mjob.Job {
	var job []mjob.Job
	config.DB.Preload("Mitra").Find(&job)
	return job
}

func (jr *JobRepoIMPL) GetByIDJob(id string) mjob.Job {
	var job mjob.Job
	config.DB.First(&job, "id = ?", id).Preload("Mitra").Find(&job)
	return job
}

func (jr *JobRepoIMPL) CreateJob(input mjob.InputJob) mjob.Job {
	var job mjob.Job = mjob.Job{
		MitraID:     input.MitraID,
		Title:       input.Title,
		Logo:        input.Logo,
		Description: input.Description,
		Requirement: input.Requirement,
		Skill:       input.Skill,
		Address:     input.Address,
		Email:       input.Email,
		NoHp:        input.NoHp,
		Link:        input.Link,
	}
	res := config.DB.Create(&job).Preload("Mitra").Find(&job)
	res.Last(&job)
	return job
}

func (jr *JobRepoIMPL) UpdateJob(id string, input mjob.InputJob) mjob.Job {
	var job mjob.Job = jr.GetByIDJob(id)
	job.Title = input.Title
	job.Logo = input.Logo
	job.Description = input.Description
	job.Requirement = input.Requirement
	job.Skill = input.Skill
	job.Address = input.Address
	job.Email = input.Email
	job.NoHp = input.NoHp
	job.Link = input.Link

	res := config.DB.Save(&job).Preload("Mitra").Find(&job)
	res.Last(&job)
	return job
}

func (jr *JobRepoIMPL) DeleteJob(id string) bool {
	var job mjob.Job = jr.GetByIDJob(id)
	res := config.DB.Delete(&job)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
