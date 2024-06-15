package seedDb

import (
	"strings"

	"github.com/jaswdr/faker/v2"
	"github.com/teooliver/kanban/internal/repository/task"
)

func createRandomTask() task.TaskForCreate {
	fake := faker.New()

	task := task.TaskForCreate{
		Title:       strings.Join(fake.Lorem().Words(3), " "),
		Description: strings.Join(fake.Lorem().Words(5), " "),
		Color:       fake.Lorem().Faker.Color().ColorName(),
		// Not needed for now:
		// StatusID:    fake.Lorem().Faker.UUID().V4(),
		// UserID:      fake.Lorem().Faker.UUID().V4(),
	}

	return task
}

func CreateMultipleTasks(nbTasks int) []task.TaskForCreate {
	tasks := make([]task.TaskForCreate, nbTasks)
	task := createRandomTask()

	for i := 0; i < nbTasks; i++ {
		tasks = append(tasks, task)
	}

	return tasks
}
