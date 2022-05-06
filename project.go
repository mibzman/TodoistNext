package main

import "encoding/json"

func UnmarshalProjects(data []byte) ([]Project, error) {
	var r []Project
	err := json.Unmarshal(data, &r)
	return r, err
}

// func (r []Project) Marshal() ([]byte, error) {
// 	return json.Marshal(r)
// }

type Project struct {
	ID           int64  `json:"id"`
	Color        int64  `json:"color"`
	Name         string `json:"name"`
	CommentCount int64  `json:"comment_count"`
	Shared       bool   `json:"shared"`
	Favorite     bool   `json:"favorite"`
	SyncID       int64  `json:"sync_id"`
	InboxProject bool   `json:"inbox_project,omitempty"`
	Order        int64  `json:"order,omitempty"`
	Parent       int64  `json:"parent,omitempty"`
	ParentID     int64  `json:"parent_id,omitempty"`
}

func GetChildrenIDsByParentID(Projects []Project, ParentID int64) (Result []int64) {

	for i := range Projects {
		if Projects[i].ParentID == ParentID {
			CurrentProjectID := Projects[i].ID
			Result = append(Result, CurrentProjectID)
			Result = append(Result, GetChildrenIDsByParentID(Projects, CurrentProjectID)...)
		}
	}
	return
}

func GetProjects() ([]Project, error) {
	body := TodoistGetReq("https://api.todoist.com/rest/v1/projects")
	return UnmarshalProjects(body)
}
