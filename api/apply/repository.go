package apply
import  "github.com/LibenHailu/fjobs/api/entity"
// Repository interface for apply
type ApplyRepository interface{
	StoreApply(apply *entity.Apply) (*entity.Apply, []error)
}