package main5

import "fmt"

// Repository pattern example - in-memory implementation (cont.)

func (r *ProjectRepository) Store(p project.Project) (project.Project, error) {
	if r.projects == nil {
		r.projects = map[int64]string{}
	}
	if p.ID <= 0 {
		r.highestID++
		p.ID = r.highestID
	} else {
		if _, exists := r.projects[p.ID]; !exists {
			return p, fmt.Errorf("Cannot update project: %d", p.ID)
		}
	}
	r.projects[p.ID] = p.Name
	return p, nil
}
