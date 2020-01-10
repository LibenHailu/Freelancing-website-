package apply
import  "github.com/LibenHailu/fjobs/api/entity"

type ApplyRepository interface{
	StoreApply(apply *entity.Apply) (*entity.Apply, []error)
}