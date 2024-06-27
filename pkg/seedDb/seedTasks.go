package seedDb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/pkg/error"
)

func createFakeTask(statusID string, userId string) task.Task {
	fake := faker.New()

	task := task.Task{
		ID:          uuid.New().String(),
		Title:       strings.Join(fake.Lorem().Words(3), " "),
		Description: strings.Join(fake.Lorem().Words(5), " "),
		Color:       fake.Lorem().Faker.Color().ColorName(),
		StatusID:    &statusID,
		UserID:      &userId,
	}

	return task
}

func createMultipleTasks(nbTasks int, statusID string, userId string) []task.Task {
	tasks := make([]task.Task, nbTasks)
	task := createFakeTask(statusID, userId)

	for i := 0; i < nbTasks; i++ {
		tasks = append(tasks, task)
	}

	return tasks
}

func taskIntoCSVString(tasks []task.Task) []string {
	s := make([]string, 0, len(tasks))

	for _, t := range tasks {
		result := fmt.Sprintf("%s, %s, %s, %s, %s, %s", t.ID, t.Title, t.Color, t.Description, error.ZeroOnNil(t.StatusID), error.ZeroOnNil(t.UserID))
		s = append(s, result)
	}

	return s
}
