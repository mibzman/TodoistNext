package main

import "log"

func MoveNotionNotes() {
	NotionProjectID := int64(<secret>)
	NotionTasks, _ := GetTasks(NotionProjectID, "")

	for _, Task := range NotionTasks {
		MoveTaskToNotion(Task)
	}
}

func MoveTaskToNotion(Task Task) error {
	log.Printf("Moving %v", Task.Content)
	err := AddToToTable("<secret>", Task.Content)
	if err != nil {
		return err
	}
	Task.Complete()
	return nil
}
