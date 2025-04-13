package domain

import "to-do-planner/internal/scheduler"

type Developer struct {
	Name     string
	Capacity int
}

type Developers []Developer

func (developer Developer) ToSchedularDevelopers() scheduler.Developer {
	return scheduler.Developer{
		Name:     developer.Name,
		Capacity: developer.Capacity,
	}
}

func (developers Developers) ToSchedularDevelopers() []scheduler.Developer {
	result := make([]scheduler.Developer, len(developers))

	for i, developer := range developers {
		result[i] = developer.ToSchedularDevelopers()
	}

	return result
}
