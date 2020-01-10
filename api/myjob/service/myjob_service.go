package service

import (
	"github.com/LibenHailu/fjobs/api/entity"
	"github.com/LibenHailu/fjobs/api/myjob"
)

type MyJobService struct {
	myjobRepo myjob.MyJobRepository
}

func NewMyJobService(repo myjob.MyJobRepository) *MyJobService {
	return &MyJobService{myjobRepo: repo}
}
func (mjs *MyJobService) StoreMyJob(myjob *entity.MyJob) (*entity.MyJob, []error) {
	jobs, errs := mjs.myjobRepo.StoreMyJob(myjob)
	if len(errs) > 0 {
		return nil, errs
	}

	return jobs, nil
}
func (mjs *MyJobService) GetMyJob(id int) ([]entity.MyJob, []error) {
	job, errs := mjs.myjobRepo.GetMyJob(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return job, nil
}
