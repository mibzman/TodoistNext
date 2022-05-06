package main

import "log"

func UpdateOnosysNext() {
	UpdateNextTasks(int64(<taskid>))
}

func UpdateActiveNext() {
	UpdateNextTasks(int64(<taskid>))
}

func UpdateNextTasks(ActiveProject int64) {

	AllProjects, err := GetProjects()
	if err != nil {
		log.Fatal(err)
	}
	ProjectIDs := GetChildrenIDsByParentID(AllProjects, ActiveProject)
	ProjectIDs = append(ProjectIDs, ActiveProject)
	log.Printf("%v", ProjectIDs)

	for _, ProjectID := range ProjectIDs {
		ProcessProject(ProjectID)
	}
}

func ProcessProject(ProjectID int64) {
	NextTasks, err := GetNextTasks(ProjectID)
	if err != nil {
		log.Fatal(err)
	}

	if len(NextTasks) > 0 {
		return
	}

	NonNextTasks, err := GetPotentialNextTasks(ProjectID)
	if err != nil {
		log.Fatal(err)
	}

	if len(NonNextTasks) == 0 {
		return
	}

	First := GetFirstTask(NonNextTasks)

	log.Printf("Trying to add label to this task: %v #%v", First.Content, First.ID)
	First.AddNextTag()
}

func GetFirstTask(Tasks []Task) Task {
	Winner := Tasks[0]
	Winner.Order = 9999
	for i := range Tasks {
		if Tasks[i].ParentID == nil {
			if Tasks[i].Order < Winner.Order {
				Winner = Tasks[i]
			}
		}
	}
	return Winner
}
