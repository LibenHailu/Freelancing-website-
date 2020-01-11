package apply

import "github.com/LibenHailu/fjobs/api/entity"
// Service interface for apply
type ApplyService interface {
	StoreApply(apply *entity.Apply) (*entity.Apply, []error)
}
