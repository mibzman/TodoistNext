// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tasks, err := UnmarshalTasks(bytes)
//    bytes, err = tasks.Marshal()

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// type Tasks []Task

func UnmarshalTasks(data []byte) ([]Task, error) {
	var r []Task
	err := json.Unmarshal(data, &r)
	return r, err
}

// func (r *Tasks) Marshal() ([]byte, error) {
// 	return json.Marshal(r)
// }

type Task struct {
	ID           int64   `json:"id"`
	Assigner     int64   `json:"assigner"`
	ProjectID    int64   `json:"project_id"`
	SectionID    int64   `json:"section_id"`
	Order        int64   `json:"order"`
	Content      string  `json:"content"`
	Description  string  `json:"description"`
	Completed    bool    `json:"completed"`
	LabelIDS     []int64 `json:"label_ids"`
	Priority     int64   `json:"priority"`
	CommentCount int64   `json:"comment_count"`
	Creator      int64   `json:"creator"`
	Created      string  `json:"created"`
	URL          string  `json:"url"`
	Parent       *int64  `json:"parent,omitempty"`
	ParentID     *int64  `json:"parent_id,omitempty"`
}

func GetTasks(ProjectID int64, filter string) ([]Task, error) {
	body := TodoistGetReq(fmt.Sprintf("https://api.todoist.com/rest/v1/tasks?project_id=%v&filter=%v", ProjectID, filter))
	return UnmarshalTasks(body)
}

func GetNextTasks(ProjectID int64) ([]Task, error) {
	return GetTasks(ProjectID, "%40next")
}

func GetPotentialNextTasks(ProjectID int64) ([]Task, error) {
	return GetTasks(ProjectID, "!%40next%26no%20date")
}

func (Task Task) AddNextTag() {
	var jsonData = []byte(`{"label_ids": [<secret>]}`)

	TodoistPostReq(fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%v", Task.ID), bytes.NewBuffer(jsonData))
}

func (Task Task) Complete() {
	var jsonData = []byte(``)

	TodoistPostReq(fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%v/close", Task.ID), bytes.NewBuffer(jsonData))
}
