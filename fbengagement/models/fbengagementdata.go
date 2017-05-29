package models

import "encoding/json"

type FbPostMarshaller interface {
	Unmarshall_FbPost(response string)
}
type FbengagementData struct {
	Posts FBPost
	Id    string
}

func (fbData *FbengagementData) Unmarshall_FbPost(response string) error {

	err := json.Unmarshal([]byte(response), fbData)
	return err
}
func (fbData *FBPost) Unmarshall_FbPost(response string) error {

	err := json.Unmarshal([]byte(response), fbData)
	return err
}

type FBPost struct {
	Data   []Data
	Paging Paging
}

type Paging struct {
	Next string
}

type Data struct {
	Shares                   Shares
	Created_time             string
	Updated_time             string
	Id, Message, Description string
	Likes                    Likes
	Comments                 Comments
}

type Shares struct {
	Count int
}

type Comments struct {
	Summary Summary
}
type Likes struct {
	Summary Summary
}

type Summary struct {
	Total_count int
}
