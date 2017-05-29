package main

import (
	"fmt"
	daos "fullmetalgo/dao"
	"fullmetalgo/fbengagement"
	"fullmetalgo/fbreach"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	Version = "1.0.0"
	Build   = "2016"
)

func main() {

	fbreach.FetchAndStoreReachInterest()

	s := gocron.NewScheduler()
	//s.Every(3).Seconds().Do(taskManagement)
	gocron.Every(1).Hour().Do(taskManagement)
	<-s.Start()

}

func taskManagement() {
	fmt.Println("calling task")
	var fbEngagementDao daos.FbenagagementDao
	var fbengagement_temp fbengagement.Fbengagement
	pageIds := fbEngagementDao.GetPageIds()
	posthistoryTime := time.Now().UTC()
	for i := len(pageIds) - 1; i >= 0; i-- {
		fmt.Println("page id", pageIds[i])
		fbengagement_temp.ReadDataFromFB(pageIds[i], posthistoryTime)
	}
}
