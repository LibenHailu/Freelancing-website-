package repository

import (
	"fmt"

	"github.com/LibenHailu/fjobs/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type ApplyGormRepo struct {
	conn *gorm.DB
}

func NewApplyGormRepo(dbconn *gorm.DB) *ApplyGormRepo {
	return &ApplyGormRepo{conn: dbconn}
}

func (ar *ApplyGormRepo) StoreApply(apply *entity.Apply) (*entity.Apply, []error) {
	aly := apply
	errs := ar.conn.Create(aly).GetErrors()

	for _, err := range errs {
		pqerr := err.(*pq.Error)
		fmt.Println(pqerr)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return aly, nil
}
