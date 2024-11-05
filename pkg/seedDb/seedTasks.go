package seedDb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/pkg/error"
)

func createFakeTask(statusID string, userId string, columnID string) task.Task {
	task := task.Task{
		ID:          uuid.New().String(),
		Title:       strings.Join(fake.Lorem().Words(3), " "),
		Description: strings.Join(fake.Lorem().Words(5), " "),
		Color:       fake.Lorem().Faker.Color().ColorName(),
		StatusID:    &statusID,
		ColumnID:    columnID,
		UserID:      &userId,
	}

	return task
}

func createMultipleTasks(nbTasks int, statusID string, userId string, columnID string) []task.Task {
	tasks := make([]task.Task, 0, nbTasks)

	for i := 0; i < nbTasks; i++ {
		task := createFakeTask(statusID, userId, columnID)
		tasks = append(tasks, task)
	}

	return tasks
}

func taskIntoCSVString(tasks []task.Task) []string {
	s := make([]string, 0, len(tasks))

	tasksCSVHeader := "id,title,description,color,user_id,status_id,column_id"
	s = append(s, tasksCSVHeader)

	for _, t := range tasks {
		result := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", t.ID, t.Title, t.Description, t.Color, error.ZeroOrNil(t.UserID), error.ZeroOrNil(t.StatusID), t.ColumnID)
		s = append(s, result)
	}

	return s
}
