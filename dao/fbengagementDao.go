package dao

import (
	"fmt"
	model "fullmetalgo/fbengagement/models"
	"fullmetalgo/fmdatabase"

	"time"
)

var (
	db = fmdatabase.GetInstance()
)

type FbenagagementDao struct {
}

func (fbengagementDao *FbenagagementDao) GetPageIds() []string {
	rows, err := db.Query("SELECT fb_page_id FROM fb_page")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	// Fetch rows
	totalRows := []string{}

	for rows.Next() {
		var pageid string
		err = rows.Scan(&pageid)
		totalRows = append(totalRows, pageid)
		checkErr(err)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return totalRows
}

func (fbengagementDao *FbenagagementDao) saveFBPageHistory(postData model.Data, pageId string, postStoredTime time.Time) {
	stmt, err := db.Prepare("INSERT fb_post_history SET fk_fb_post_id =?, fb_post_history_likes=?, fb_post_history_comments=?,fb_post_history_shares=?, fb_post_history_time=?, fk_fb_page_id =?")

	_, err = stmt.Exec(postData.Id, postData.Likes.Summary.Total_count, postData.Comments.Summary.Total_count, postData.Shares.Count, postStoredTime, pageId)
	checkErr(err)
}

func (fbengagementDao *FbenagagementDao) WriteToTable() {

	stmt, err := db.Prepare("INSERT fb_page SET fb_page_name=?,fb_page_id=?,fb_page_likes=?,fb_page_talking_about_count=?")
	_, err = stmt.Exec("Mandir", "149027717730", "101", "1123")
	checkErr(err)

}

var (
	fb_post_id string
)

func (fbengagementDao *FbenagagementDao) IsPostExist(post_id string) bool {
	rows, err := db.Query("SELECT fb_post_id  FROM fb_post WHERE fb_post_id = ?", post_id)
	checkErr(err)
	defer rows.Close()
	var result bool
	if rows.Next() {
		err := rows.Scan(&fb_post_id)
		checkErr(err)
		if len(fb_post_id) > 0 {
			result = true
		}
	}
	err = rows.Err()
	checkErr(err)
	return result
}

func (fbengagementDao *FbenagagementDao) SaveFbPost(postData model.Data, id string, postStoredTime time.Time) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	if fbengagementDao.IsPostExist(postData.Id) {
		fbengagementDao.UpdateFbPostData(postData, id)
	} else {
		fbengagementDao.SaveFBPostData(postData, id)
	}
	fbengagementDao.saveFBPageHistory(postData, id, postStoredTime)

}

func (fbengagementDao *FbenagagementDao) SaveFBPostData(postData model.Data, pageId string) {
	stmt, err := db.Prepare("INSERT fb_post SET fk_fb_post_page_id =?, fb_post_id=?, fb_post_message=?,fb_post_description=?, fb_post_create_time=?, fb_post_update_time =?, fb_post_likes=?, fb_post_comments=?, fb_post_shares=?")

	checkErr(err)
	createdTime, updatedTime, err := getCreatedAndUpdateTimeStamp(postData.Created_time, postData.Updated_time)
	checkErr(err)
	_, err = stmt.Exec(pageId, postData.Id, postData.Message, postData.Description, createdTime, updatedTime, postData.Likes.Summary.Total_count, postData.Comments.Summary.Total_count, postData.Shares.Count)
	checkErr(err)
}

func (fbengagementDao *FbenagagementDao) UpdateFbPostData(postData model.Data, pageId string) {
	tx, err := db.Begin()
	stmt, err := tx.Prepare("UPDATE fb_post SET fk_fb_post_page_id=?, fb_post_message=?,fb_post_description=?, fb_post_create_time=?, fb_post_update_time =?, fb_post_likes=?, fb_post_comments=?, fb_post_shares=? WHERE fb_post_id=?")

	checkErr(err)
	createdTime, updatedTime, err := getCreatedAndUpdateTimeStamp(postData.Created_time, postData.Updated_time)
	checkErr(err)
	insert_stmt_r1 := tx.Stmt(stmt)
	_, err = insert_stmt_r1.Exec(pageId, postData.Message, postData.Description, createdTime, updatedTime, postData.Likes.Summary.Total_count, postData.Comments.Summary.Total_count, postData.Shares.Count, postData.Id)
	tx.Commit()

	checkErr(err)
}

func ConvertAsTimeStramp(timeValue string) (time.Time, error) {
	val, err := time.Parse("2006-01-02T15:04:05+0000", timeValue)
	checkErr(err)
	return val, err
}

func getCreatedAndUpdateTimeStamp(createdTime string, updatedTime string) (time.Time, time.Time, error) {
	created_time, err := ConvertAsTimeStramp(createdTime)
	updated_time, err := ConvertAsTimeStramp(updatedTime)
	return created_time, updated_time, err

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
