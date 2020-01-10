package myjob

import "github.com/LibenHailu/fjobs/api/entity"

type MyJobService interface {
	StoreMyJob(myjob *entity.MyJob) (*entity.MyJob, []error)
	GetMyJob(id int) ([]entity.MyJob, []error)
	// UpdateJob(id int) (*entity., []error)
}
