package fbengagement

import (
	"fmt"
	daos "fullmetalgo/dao"
	"testing"
	"time"
)

//"fmt"

func TestFbengagement(t *testing.T) {
	fmt.Println("Processing engagement data...")
	var fbEngagementDao daos.FbenagagementDao
	var fbengagement_temp Fbengagement
	pageIds := fbEngagementDao.GetPageIds()
	posthistoryTime := time.Now().UTC()
	for i := len(pageIds) - 1; i >= 0; i-- {
		fmt.Println("*******Page id *****", pageIds[i])
		fbengagement_temp.ReadDataFromFB(pageIds[i], posthistoryTime)
	}
}
