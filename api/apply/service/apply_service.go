package service

import (
	"github.com/LibenHailu/fjobs/api/apply"
	"github.com/LibenHailu/fjobs/api/entity"
)

type ApplyService struct {
	applyRepo apply.ApplyRepository
}

func NewApplyService(repo apply.ApplyRepository) *ApplyService {
	return &ApplyService{applyRepo: repo}
}

func (as *ApplyService) StoreApply(apply *entity.Apply) (*entity.Apply, []error) {
	aly, errs := as.applyRepo.StoreApply(apply)
	if len(errs) > 0 {
		return nil, errs
	}

	return aly, nil
}
