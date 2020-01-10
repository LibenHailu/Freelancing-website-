package service

import (
	"fmt"

	"github.com/LibenHailu/fjobs/api/entity"
	"github.com/LibenHailu/fjobs/api/job"
)

type JobService struct {
	jobRepo job.JobRepository
}

func NewJobService(repo job.JobRepository) *JobService {
	return &JobService{jobRepo: repo}
}
func (js *JobService) Jobs(category string) ([]entity.Job, []error) {
	jobs, errs := js.jobRepo.Jobs(category)
	if len(errs) > 0 {
		return nil, errs
	}

	return jobs, nil
}

func (js *JobService) Job(job *entity.Job) (*entity.Job, []error) {
	job, errs := js.jobRepo.Job(job)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}

	return job, nil
}

func (js *JobService) JobByID(id uint) (*entity.Job, []error) {
	job, errs := js.jobRepo.JobByID(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return job, nil
}

func (js *JobService) UpdateJob(job *entity.Job) (*entity.Job, []error) {
	myjob, errs := js.jobRepo.UpdateJob(job)

	if len(errs) > 0 {
		return nil, errs
	}

	return myjob, nil
}

func (js *JobService) DeleteJob(id uint) (*entity.Job, []error) {
	usr, errs := js.jobRepo.DeleteJob(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

func (js *JobService) StoreJob(job *entity.Job) (*entity.Job, []error) {
	myjob, errs := js.jobRepo.StoreJob(job)
	if len(errs) > 0 {
		return nil, errs
	}

	return myjob, nil
}
