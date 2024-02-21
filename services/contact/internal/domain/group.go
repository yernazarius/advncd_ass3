package domain

import "fmt"

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) (*Group, error) {
	if len(name) > 250 {
		return nil, fmt.Errorf("too long: %d max characters = 250", len(name))
	}

	return &Group{
		ID:   id,
		Name: name,
	}, nil
}
