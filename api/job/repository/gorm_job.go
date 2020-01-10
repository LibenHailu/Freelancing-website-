package repository

import (
	"fmt"

	"github.com/LibenHailu/fjobs/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type JobGormRepo struct {
	conn *gorm.DB
}

func NewJobGormRepo(dbconn *gorm.DB) *JobGormRepo {
	return &JobGormRepo{conn: dbconn}
}
func (jr *JobGormRepo) Jobs(category string) ([]entity.Job, []error) {
	jobs := []entity.Job{}
	errs := jr.conn.Where("Category = ?", category).Find(&jobs).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return jobs, nil
}
func (jr *JobGormRepo) Job(job *entity.Job) (*entity.Job, []error) {
	myJob := entity.Job{}
	errs := jr.conn.Where("category = ?", job.Category).Find(&myJob).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &myJob, nil
}
func (jr *JobGormRepo) JobByID(id uint) (*entity.Job, []error) {
	job := entity.Job{}
	errs := jr.conn.First(&job, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &job, nil
}

func (jr *JobGormRepo) UpdateJob(job *entity.Job) (*entity.Job, []error) {
	myJob := job
	errs := jr.conn.Save(job).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return myJob, nil
}

func (jr *JobGormRepo) DeleteJob(id uint) (*entity.Job, []error) {
	job, errs := jr.JobByID(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = jr.conn.Delete(job, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return job, nil
}

func (jr *JobGormRepo) StoreJob(job *entity.Job) (*entity.Job, []error) {
	myjob := job
	errs := jr.conn.Create(myjob).GetErrors()

	for _, err := range errs {
		pqerr := err.(*pq.Error)
		fmt.Println(pqerr)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return myjob, nil
}
