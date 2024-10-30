package seedDb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/teooliver/kanban/internal/repository/project"
)

func createFakeProject() project.Project {
	project := project.Project{
		ID:          uuid.New().String(),
		Name:        strings.Join(fake.Lorem().Words(3), " "),
		Description: strings.Join(fake.Lorem().Words(5), " "),
		IsArchived:  false,
	}

	return project
}

func createMultipleProjects(nbProjects int) []project.Project {
	projects := make([]project.Project, 0, nbProjects)

	for i := 0; i < nbProjects; i++ {
		project := createFakeProject()
		projects = append(projects, project)
	}

	return projects
}

func projectsIntoCSVString(project []project.Project) []string {
	s := make([]string, 0, len(project))

	projectsCSVHeader := "id,name,description,is_archived"
	s = append(s, projectsCSVHeader)

	for _, p := range project {
		result := fmt.Sprintf("%s,%s,%s,%v", p.ID, p.Name, p.Description, p.IsArchived)
		s = append(s, result)
	}

	return s
}
