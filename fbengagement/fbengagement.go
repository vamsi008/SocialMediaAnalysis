package fbengagement

import (
	"bytes"
	"fmt"
	"fullmetalgo/communicator"
	daos "fullmetalgo/dao"
	model "fullmetalgo/fbengagement/models"
	"time"

	"github.com/BurntSushi/toml"
)

type Fbengagement struct {
	fbengagementDao daos.FbenagagementDao
}

type fbengagementConfig struct {
	Url, Access_token string
}

func (fbengagement *Fbengagement) ReadDataFromFB(pageNumber string, posthistoryTime time.Time) {

	encodedurl, _ := FormEncodedUrl(pageNumber)
	next := fbengagement.ProcessFbengagement(encodedurl, true, pageNumber, posthistoryTime)
	for next != "" {
		next = fbengagement.ProcessFbengagement(next, false, pageNumber, posthistoryTime)
	}
	fmt.Println("Process done")
}

func (fbengagement *Fbengagement) ProcessFbengagement(encodedurl string, isFirstTime bool, pageId string, posthistoryTime time.Time) string {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Into the differ before request")
			fmt.Println(err)

		}
	}()
	var httpCommunicator communicator.HttpCommunicator
	response, _ := httpCommunicator.Communicate(encodedurl)

	if isFirstTime {
		var fbPostDataEngagement model.FbengagementData
		fbPostDataEngagement.Unmarshall_FbPost(response)
		fbengagement.SaveFbPosts(fbPostDataEngagement.Posts, pageId, posthistoryTime)
		return fbPostDataEngagement.Posts.Paging.Next

	} else {
		fmt.Println("Processing next chunk of data of page :", pageId)
		var fbPostData model.FBPost
		fbPostData.Unmarshall_FbPost(response)
		fbengagement.SaveFbPosts(fbPostData, pageId, posthistoryTime)
		return fbPostData.Paging.Next
	}

}

func FormEncodedUrl(pageId string) (string, error) {
	var config fbengagementConfig
	var err error
	_, err = toml.DecodeFile("/usr/share/fullmetal/fbengagementconfig.toml", &config)
	checkErr(err)
	var buffer bytes.Buffer
	buffer.WriteString(config.Url)
	buffer.WriteString(pageId)
	buffer.WriteString("?")
	buffer.WriteString("access_token=")
	buffer.WriteString(config.Access_token)
	buffer.WriteString("&")
	buffer.WriteString("fields=posts.since(1470009600).limit(25){shares,likes.limit(1).summary(true),comments.limit(1).summary(true),description,created_time,updated_time,message}&format=json&method=get&pretty=0&suppress_http_code=1")

	return buffer.String(), err
}

func (fbengagement *Fbengagement) SaveFbPosts(fbPostData model.FBPost, id string, posthistoryTime time.Time) {
	fmt.Println("Total posts fetched ", len(fbPostData.Data))
	for i := len(fbPostData.Data) - 1; i >= 0; i-- {
		postData := fbPostData.Data[i]
		fbengagement.fbengagementDao.SaveFbPost(postData, id, posthistoryTime)

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
