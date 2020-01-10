package apply

import "github.com/LibenHailu/fjobs/api/entity"

type ApplyService interface {
	StoreApply(apply *entity.Apply) (*entity.Apply, []error)
}
