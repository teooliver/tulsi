package seedDb

import (
	"log"
	"os"

	"github.com/teooliver/tulsi/internal/repository/column"
	"github.com/teooliver/tulsi/internal/repository/project"
	"github.com/teooliver/tulsi/internal/repository/status"
	"github.com/teooliver/tulsi/internal/repository/task"
	"github.com/teooliver/tulsi/internal/repository/user"
)

var usersAmount = 10
var tasksPerUserAmount = 5

type DbData struct {
	StatusList []status.Status
	Users      []user.User
	Tasks      []task.Task
	Projects   []project.Project
	Columns    []column.Column
}

func seedData() DbData {
	statusList := createFakeStatusList()
	users := createMultipleFakeUsers(usersAmount)
	projects := createMultipleProjects(5)
	columns := createMultipleColumns(projects[0].ID)

	tasks := make([]task.Task, 0, len(users)*tasksPerUserAmount)
	for _, u := range users {
		userTasks := createMultipleTasks(tasksPerUserAmount, statusList[0].ID, u.ID, columns[0].ID)
		tasks = append(tasks, userTasks...)
	}

	// Create Tasks without users or status
	// blankTasks := createMultipleTasks(20, "", "")
	// tasks = append(tasks, blankTasks...)

	return DbData{
		StatusList: statusList,
		Users:      users,
		Tasks:      tasks,
		Projects:   projects,
		Columns:    columns,
	}
}

func CreateDbCSV() {
	// TODO: Should we receive dbData as an arg to the function instead?
	dbData := seedData()
	usersCSVTable := userIntoCSVString(dbData.Users)
	statusCSVTable := statusIntoCSVString(dbData.StatusList)
	tasksCSVTable := taskIntoCSVString(dbData.Tasks)
	columnsCSVTable := columnsIntoCSVString(dbData.Columns)
	projectsCSVTable := projectsIntoCSVString(dbData.Projects)

	writeCSVtoFile("users.csv", usersCSVTable)
	writeCSVtoFile("status.csv", statusCSVTable)
	writeCSVtoFile("tasks.csv", tasksCSVTable)
	writeCSVtoFile("columns.csv", columnsCSVTable)
	writeCSVtoFile("projects.csv", projectsCSVTable)

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
