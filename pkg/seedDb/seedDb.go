package seedDb

import (
	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/internal/repository/user"
)

var tasksPerUser = 5

type DbData struct {
	StatusList []status.Status
	Users      []user.User
	Tasks      []task.Task
}

func seedData() DbData {
	statusList := createFakeStatusList()
	users := createMultipleFakeUsers(10)

	tasks := make([]task.Task, len(users)*tasksPerUser)
	for _, u := range users {
		userTasks := createMultipleTasks(tasksPerUser, statusList[0].ID, u.ID)
		tasks = append(tasks, userTasks...)
	}

	return DbData{
		StatusList: statusList,
		Users:      users,
		Tasks:      tasks,
	}

}

// TODO: Create CSV file with fake to seed database
func createCSV(data DbData) {

}
