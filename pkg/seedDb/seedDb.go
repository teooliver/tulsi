package seedDb

import (
	"log"
	"os"

	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/internal/repository/user"
)

// TODO: Create CLI tool to define those types of vars and create DbData
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

	// Create Tasks without users or status
	blankTasks := createMultipleTasks(20, "", "")
	tasks = append(tasks, blankTasks...)

	return DbData{
		StatusList: statusList,
		Users:      users,
		Tasks:      tasks,
	}
}

func CreateDbCSV() {
	// TODO: Should we receive dbData as an arg to the function instead?
	dbData := seedData()
	usersCSVTable := userIntoCSVString(dbData.Users)
	statusCSVTable := statusIntoCSVString(dbData.StatusList)
	tasksCSVTable := taskIntoCSVString(dbData.Tasks)

	writeCSVtoFile("users.csv", statusCSVTable)
	writeCSVtoFile("status.csv", usersCSVTable)
	writeCSVtoFile("tasks.csv", tasksCSVTable)

}

func writeCSVtoFile(fileName string, lines []string) {
	// TODO: Check if path already exist
	err := os.Mkdir("CSV_DB/", 0755)
	if err != nil {
		// TODO: If folder already exists do nothing
		// log.Fatal(err)
	}

	// TODO: Allow overriding the files if they already exist
	f, err := os.Create("CSV_DB/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
