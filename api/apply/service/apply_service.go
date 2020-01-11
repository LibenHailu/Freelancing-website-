package service

import (
	"github.com/LibenHailu/fjobs/api/apply"
	"github.com/LibenHailu/fjobs/api/entity"
)
// struc which used for implementing ApplyRepository interfaces
type ApplyService struct {
	applyRepo apply.ApplyRepository
}
// init ApplyService
func NewApplyService(repo apply.ApplyRepository) *ApplyService {
	return &ApplyService{applyRepo: repo}
}
// calls the  StoreApply which apperes in apply's repository layer
func (as *ApplyService) StoreApply(apply *entity.Apply) (*entity.Apply, []error) {
	aly, errs := as.applyRepo.StoreApply(apply)
	if len(errs) > 0 {
		return nil, errs
	}

	return aly, nil
}
