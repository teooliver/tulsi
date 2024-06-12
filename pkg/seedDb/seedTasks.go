package seedDb

import (
	"github.com/jaswdr/faker/v2"
	"github.com/teooliver/kanban/internal/repository/task"
)

func createRandomTask() task.TaskForCreate {
	fake := faker.New()

	task := task.TaskForCreate{
		Title:       fake.Lorem().Sentence(5),
		Description: fake.Lorem().Paragraph(1),
		StatusID:    fake.Lorem().Faker.UUID().V4(),
		Color:       fake.Lorem().Faker.Color().ColorName(),
		UserID:      fake.Lorem().Faker.UUID().V4(),
	}

	return task
}

func CreateMultipleTasks(nbTasks int) []task.TaskForCreate {
	tasks := make([]task.TaskForCreate, nbTasks)

	for i := 0; i < nbTasks; i++ {
		tasks = append(tasks, createRandomTask())
	}

	return tasks
}
